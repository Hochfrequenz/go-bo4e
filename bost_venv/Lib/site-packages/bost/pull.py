"""
Contains functions to pull the BO4E-Schemas from GitHub.
"""
from functools import lru_cache
from itertools import chain
from pathlib import Path
from typing import Iterable, Union

import requests
from pydantic import BaseModel, TypeAdapter, ValidationError
from requests import Response

from bost.cache import CACHE_FILE_NAME, CacheData, get_cached_file, get_cached_file_tree, save_cache
from bost.config import Config
from bost.logger import logger
from bost.schema import Object, Reference, SchemaRootType, StrEnum

OWNER = "Hochfrequenz"
REPO = "BO4E-Schemas"
TIMEOUT = 10  # in seconds


class SchemaMetadata(BaseModel):
    """
    Metadata about a schema file
    """

    _schema: SchemaRootType | None = None
    class_name: str
    download_url: str
    module_path: tuple[str, ...]
    "e.g. ('bo', 'Angebot')"
    file_path: Path
    cached_path: Path | None

    @property
    def module_name(self) -> str:
        """
        Joined module path. E.g. "bo.Angebot"
        """
        return ".".join(self.module_path)

    @property
    def schema_parsed(self) -> SchemaRootType:
        """
        The parsed schema. Downloads the schema from GitHub if needed.
        """
        if self._schema is None:
            if self.cached_path is not None and self.cached_path.exists():
                self._schema = TypeAdapter(SchemaRootType).validate_json(  # type: ignore[assignment]
                    self.cached_path.read_text()
                )
                logger.info("Loaded %s from cache", self.cached_path)
            else:
                schema_response = self._download_schema()
                self._schema = TypeAdapter(SchemaRootType).validate_json(  # type: ignore[assignment]
                    schema_response.text
                )
        assert self._schema is not None
        return self._schema

    @schema_parsed.setter
    def schema_parsed(self, value: SchemaRootType):
        self._schema = value

    def _download_schema(self) -> Response:
        """
        Download the schema from GitHub. Returns the response object.
        """
        response = requests.get(self.download_url, timeout=TIMEOUT)
        if response.status_code != 200:
            raise ValueError(f"Could not download schema from {self.download_url}: {response.text}")
        logger.info("Downloaded %s", self.download_url)
        if self.cached_path is not None:
            self.cached_path.parent.mkdir(parents=True, exist_ok=True)
            self.cached_path.write_text(response.text)
            logger.debug("Cached %s", self.cached_path)
        return response

    def save(self):
        """
        Save the parsed schema to the file defined by `file_path`. Creates parent directories if needed.
        """
        self.file_path.parent.mkdir(parents=True, exist_ok=True)
        self.file_path.write_text(
            self.schema_parsed.model_dump_json(indent=2, exclude_unset=True, by_alias=True), encoding="utf-8"
        )

    def field_paths(self) -> Iterable[tuple[str, str]]:
        """
        Get all field paths of the schema.
        """
        if not isinstance(self.schema_parsed, Object):
            return
        for field_name in self.schema_parsed.properties:
            yield ".".join((self.module_name, field_name)), field_name

    def __str__(self):
        return self.module_name


class SchemaInFileTree(BaseModel):
    """
    A schema in the file tree returned by the GitHub API. Only contains the relevant information.
    """

    name: str
    path: str
    download_url: str


class SchemaLists(BaseModel):
    """
    A list of schemas
    """

    bo: list[SchemaInFileTree]
    com: list[SchemaInFileTree]
    enum: list[SchemaInFileTree]


CacheData.model_rebuild()


@lru_cache(maxsize=None)
def _github_tree_query(pkg: str, version: str) -> list[SchemaInFileTree]:
    """
    Query the github tree api for a specific package and version.
    """
    response = requests.get(
        f"https://api.github.com/repos/{OWNER}/{REPO}/contents/src/bo4e_schemas/{pkg}?ref={version}", timeout=TIMEOUT
    )
    if response.status_code != 200:
        raise ValueError(f"Could not query repository tree from {response.request.url}: {response.text}")
    return TypeAdapter(list[SchemaInFileTree]).validate_json(response.text)


@lru_cache(maxsize=1)
def resolve_latest_version() -> str:
    """
    Resolve the latest BO4E version from the github api.
    """
    response = requests.get(f"https://api.github.com/repos/{OWNER}/{REPO}/releases/latest", timeout=TIMEOUT)
    response.raise_for_status()
    return response.json()["tag_name"]


def get_schema_list(version: str, cache_dir: Path | None) -> SchemaLists:
    """
    Get all files metadata from the BO4E-Schemas repository or from cache.
    """
    if cache_dir is not None:
        possible_schemas = get_cached_file_tree(cache_dir)
        if possible_schemas is not None:
            return possible_schemas

    schemas = SchemaLists(**{pkg: _github_tree_query(pkg, version) for pkg in ("bo", "com", "enum")})
    if cache_dir is not None:
        save_cache(cache_dir / CACHE_FILE_NAME, version=version, file_tree=schemas)

    return schemas


SCHEMA_CACHE: dict[tuple[str, ...], SchemaMetadata] = {}


def schema_iterator(version: str, output: Path, cache_dir: Path | None) -> Iterable[tuple[str, SchemaMetadata]]:
    """
    Get all files from the BO4E-Schemas repository.
    This generator function yields tuples of class name and SchemaMetadata objects containing various information about
    the schema.
    """
    schemas = get_schema_list(version, cache_dir)
    for file in chain(schemas.bo, schemas.com, schemas.enum):
        if not file.name.endswith(".json"):
            continue
        relative_path = Path(file.path).relative_to("src/bo4e_schemas")
        module_path = (*relative_path.parent.parts, relative_path.stem)
        if module_path not in SCHEMA_CACHE:
            SCHEMA_CACHE[module_path] = SchemaMetadata(
                class_name=relative_path.stem,
                download_url=file.download_url,
                module_path=module_path,
                file_path=output / relative_path,
                cached_path=get_cached_file(relative_path, cache_dir),
            )
        yield SCHEMA_CACHE[module_path].class_name, SCHEMA_CACHE[module_path]


def load_schema(path: Path) -> Object | StrEnum:
    """
    Load a schema from a file.
    """
    try:
        return TypeAdapter(Union[Object, StrEnum]).validate_json(path.read_text())  # type: ignore[return-value]
    except ValidationError as error:
        logger.error("Could not load schema from %s:", path, exc_info=error)
        raise


def additional_schema_iterator(
    config: Config | None, config_path: Path | None, output: Path
) -> Iterable[tuple[str, SchemaMetadata]]:
    """
    Get all additional models from the config file.
    """
    if config is None:
        return
    assert config_path is not None, "Config path must be set if config is set"

    for additional_model in config.additional_models:
        if isinstance(additional_model.schema_parsed, Reference):
            reference_path = Path(additional_model.schema_parsed.ref)
            if not reference_path.is_absolute():
                reference_path = config_path.parent / reference_path
            schema_parsed = load_schema(reference_path)
        else:
            schema_parsed = additional_model.schema_parsed

        if schema_parsed.title == "":
            raise ValueError("Config Error: Title is required for additional models to determine the class name")

        schema_metadata = SchemaMetadata(
            class_name=schema_parsed.title,
            download_url="",
            module_path=(additional_model.module, schema_parsed.title),
            file_path=output / f"{additional_model.module}/{schema_parsed.title}.json",
            cached_path=None,
        )
        schema_metadata.schema_parsed = schema_parsed
        yield schema_metadata.class_name, schema_metadata
