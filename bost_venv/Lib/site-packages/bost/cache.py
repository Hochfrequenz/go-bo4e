"""
Implement functionality to cache the queried data from GitHub.
"""
import shutil
from pathlib import Path
from typing import TYPE_CHECKING, Optional, overload

from pydantic import BaseModel

from bost.logger import logger

if TYPE_CHECKING:
    from bost.pull import SchemaLists

CACHE_FILE_NAME = ".cache"


class CacheData(BaseModel):
    """
    Data to be cached
    """

    version: str
    file_tree: "SchemaLists"


CACHED_DATA: CacheData


def load_cache(cache_file: Path) -> CacheData:
    """
    Load the cache file.
    Returns named tuple of:
    - the cached version as GitHub tag e.g. "v0.6.1-rc13"
    - the cached file tree in the same format as the GitHub API
    """
    if not cache_file.exists():
        raise FileNotFoundError(f"Cache file {cache_file} does not exist")

    # pylint: disable=global-statement
    global CACHED_DATA
    if "CACHED_DATA" in globals():
        return CACHED_DATA

    CACHED_DATA = CacheData.model_validate_json(cache_file.read_text(encoding="utf-8"))
    return CACHED_DATA


@overload
def save_cache(cache_file: Path, *, version: str) -> None:
    ...


@overload
def save_cache(cache_file: Path, *, file_tree: "SchemaLists") -> None:
    ...


@overload
def save_cache(cache_file: Path, *, version: str, file_tree: "SchemaLists") -> None:
    ...


@overload
def save_cache(cache_file: Path, *, cache_data: CacheData) -> None:
    ...


def save_cache(
    cache_file: Path,
    version: str | None = None,
    file_tree: Optional["SchemaLists"] = None,
    cache_data: CacheData | None = None,
) -> None:
    """
    Save the cache file.
    """
    if not cache_file.parent.exists():
        cache_file.parent.mkdir(parents=True, exist_ok=True)

    # pylint: disable=global-statement
    global CACHED_DATA
    if cache_data is not None:
        CACHED_DATA = cache_data
        cache_file.write_text(cache_data.model_dump_json(), encoding="utf-8")
        return
    if "CACHED_DATA" not in globals() and (version is None or file_tree is None):
        raise ValueError("No complete set of cache data provided. This shouldn't happen.")

    update_dict: dict[str, str | SchemaLists] = {}
    if version is not None:
        update_dict["version"] = version
    if file_tree is not None:
        update_dict["file_tree"] = file_tree
    if "CACHED_DATA" in globals():
        CACHED_DATA = CACHED_DATA.model_copy(update=update_dict)
    else:
        CACHED_DATA = CacheData(**update_dict)  # type: ignore[arg-type]
    cache_file.write_text(CACHED_DATA.model_dump_json(), encoding="utf-8")


def is_cache_dir_valid(cache_dir: Path | None, target_version: str) -> bool:
    """
    Check if the cache directory is valid.
    It is valid if it is empty or doesn't exist yet.
    If it is not empty but the version file is missing, raise an FileNotFoundError.
    If it doesn't contain the correct version the cache will be cleared.
    """
    if cache_dir is None:
        return False
    cache_dir.mkdir(parents=True, exist_ok=True)
    cache_file = cache_dir / CACHE_FILE_NAME
    if not any(cache_dir.iterdir()):
        return True
    if not cache_file.exists():
        raise FileNotFoundError("Cache directory is not empty but does not contain a version file")
    cached_version = load_cache(cache_file).version
    if cached_version != target_version:
        logger.warning(
            "Version mismatch: The cache directory contains version %s but the target version is %s. "
            "The files will be downloaded again and the cache will be overwritten.",
            cached_version,
            target_version,
        )
        shutil.rmtree(cache_dir)
        cache_dir.mkdir()

        # pylint: disable=global-statement
        globals().pop("CACHED_DATA")
    return True


def get_cached_file_tree(cache_dir: Path) -> Optional["SchemaLists"]:
    """
    Get the cached file tree if the cache is not empty
    """
    cache_file = cache_dir / CACHE_FILE_NAME
    if not cache_file.exists():
        return None
    return load_cache(cache_file).file_tree


def get_cached_file(relative_path: Path, cache_dir: Path | None) -> Path | None:
    """
    Get the path to the cached file specific to a schema if the cache directory is set.
    """
    if cache_dir is None:
        return None
    return cache_dir / relative_path
