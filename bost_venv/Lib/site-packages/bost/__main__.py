"""
This module is the entry point for the bost command line interface.
"""
import re
import shutil
from pathlib import Path

import click

from bost.cache import is_cache_dir_valid
from bost.config import AdditionalEnumItem, AdditionalField, load_config
from bost.logger import logger
from bost.operations import add_additional_enum_items, add_additional_property, optional_to_required, update_references
from bost.pull import SchemaMetadata, additional_schema_iterator, resolve_latest_version, schema_iterator
from bost.schema import AnyOf, Object, StrEnum


@click.command()
@click.option(
    "--output",
    "-o",
    help="Output directory to pull the JSON schemas into",
    type=click.Path(file_okay=False, path_type=Path),
    required=True,
)
@click.option(
    "--target-version",
    "-t",
    help="Target BO4E version. Defaults to latest.",
    type=str,
    default="latest",
)
@click.option(
    "--config-file",
    "-c",
    help="Path to the config file",
    type=click.Path(exists=True, dir_okay=False, path_type=Path),
    required=False,
    default=None,
)
@click.option(
    "--update-refs/--no-update-refs",
    "-r/-R",
    help="Automatically update the references in the schema files - github URLs are replaced with relative paths",
    is_flag=True,
    default=True,
)
@click.option(
    "--set-default-version/--no-set-default-version",
    "-d/-D",
    help="Automatically set or overrides the default version for '_version' fields with --target-version. "
    "This is especially useful if you want to define additional models which should "
    "always have the correct version.",
    is_flag=True,
    default=True,
)
@click.option(
    "--clear-output",
    help="Clear the output directory before saving the schemas",
    is_flag=True,
    default=False,
)
@click.option(
    "--cache-dir",
    help="Path to the optional cache dir. If not set the cache is disabled. "
    "It will cache the raw schema files downloaded from github.",
    type=click.Path(file_okay=False, path_type=Path),
    required=False,
    default=None,
)
@click.version_option(package_name="BO4E-Schema-Tool")
def main_command_line(*args, **kwargs) -> None:
    """
    Entry point for the bost command line interface.
    """
    main(*args, **kwargs)


def transform_all_required_fields(required_field_patters: list[str], schemas: dict[str, SchemaMetadata]):
    """
    Apply the required field patterns to all schemas.
    """
    field_paths = [
        (field_path, field_name, schema)
        for schema in schemas.values()
        for field_path, field_name in schema.field_paths()
    ]
    for pattern in required_field_patters:
        compiled_pattern = re.compile(pattern)
        matches = 0
        for field_path, field_name, schema in field_paths:
            if (
                compiled_pattern.fullmatch(field_path)
                and isinstance(schema.schema_parsed, Object)
                and isinstance(schema.schema_parsed.properties[field_name], AnyOf)
                and "default" in schema.schema_parsed.properties[field_name].__pydantic_fields_set__
            ):
                matches += 1
                schema.schema_parsed.properties[field_name] = optional_to_required(
                    schema.schema_parsed.properties[field_name]  # type: ignore[arg-type]
                )
                if field_name not in schema.schema_parsed.required:
                    if "required" not in schema.schema_parsed.__pydantic_fields_set__:
                        schema.schema_parsed.required = []
                    schema.schema_parsed.required.append(field_name)
                logger.info("Applied pattern '%s' to field %s", pattern, field_path)
        if matches == 0:
            logger.warning("Pattern '%s' did not match any fields", pattern)
        else:
            logger.info("Pattern '%s' matched %d fields", pattern, matches)


def transform_all_additional_fields(additional_fields: list[AdditionalField], schemas: dict[str, SchemaMetadata]):
    """
    Apply the additional field patterns to all schemas and adds the respective field definition.
    """
    schema_paths = [(schema.module_name, schema) for schema in schemas.values()]
    for additional_field in additional_fields:
        compiled_pattern = re.compile(additional_field.pattern)
        matches = 0
        for schema_path, schema in schema_paths:
            if compiled_pattern.fullmatch(schema_path) and isinstance(schema.schema_parsed, Object):
                matches += 1
                add_additional_property(schema.schema_parsed, additional_field.field_def, additional_field.field_name)

                if (
                    "default" not in additional_field.field_def.__pydantic_fields_set__
                    and additional_field.field_name not in schema.schema_parsed.required
                ):
                    if "required" not in schema.schema_parsed.__pydantic_fields_set__:
                        schema.schema_parsed.required = []
                    schema.schema_parsed.required.append(additional_field.field_name)
                logger.info(
                    "Applied pattern '%s' to schema %s. Added field %s",
                    additional_field.pattern,
                    schema_path,
                    additional_field.field_name,
                )
        if matches == 0:
            logger.warning("Pattern '%s' did not match any fields", additional_field.pattern)
        else:
            logger.info("Pattern '%s' matched %d fields", additional_field.pattern, matches)


def transform_all_additional_enum_items(
    additional_enum_items: list[AdditionalEnumItem], schemas: dict[str, SchemaMetadata]
):
    """
    Apply the additional enum item patterns to all schemas and adds the respective enum items.
    """
    schema_paths = [(schema.module_name, schema) for schema in schemas.values()]
    for additional_item in additional_enum_items:
        compiled_pattern = re.compile(additional_item.pattern)
        matches = 0
        for schema_path, schema in schema_paths:
            if compiled_pattern.fullmatch(schema_path) and isinstance(schema.schema_parsed, StrEnum):
                matches += 1
                add_additional_enum_items(schema.schema_parsed, additional_item.items)
                logger.info(
                    "Applied pattern '%s' to schema %s. Added enum items %s",
                    additional_item.pattern,
                    schema_path,
                    str(additional_item.items),
                )
        if matches == 0:
            logger.warning("Pattern '%s' did not match any fields", additional_item.pattern)
        else:
            logger.info("Pattern '%s' matched %d fields", additional_item.pattern, matches)


# pylint: disable=too-many-branches, too-many-arguments
def main(
    output: Path,
    target_version: str,
    update_refs: bool,
    set_default_version: bool,
    clear_output: bool,
    config_file: Path | None = None,
    cache_dir: Path | None = None,
) -> None:
    """
    Pull the schemas from the BO4E repository and apply the operations defined in the config file.
    """
    if config_file is not None:
        config = load_config(config_file)
    else:
        config = None

    if target_version == "latest":
        target_version = resolve_latest_version()

    if not is_cache_dir_valid(cache_dir, target_version):
        cache_dir = None
    schemas = dict(schema_iterator(target_version, output, cache_dir))

    if config is not None:
        schemas.update(additional_schema_iterator(config, config_file, output))
        logger.info("Added all additional models")
        transform_all_required_fields(config.required_fields, schemas)
        logger.info("Transformed all required fields")
        transform_all_additional_fields(config.additional_fields, schemas)  # type: ignore[arg-type]
        # the load_config function ensures that the references are resolved.
        logger.info("Added all additional fields")
        transform_all_additional_enum_items(config.additional_enum_items, schemas)
        logger.info("Added all additional enum items")

    if update_refs:
        for schema in schemas.values():
            update_references(schema.schema_parsed, schema.module_path)
        logger.info("Updated github references")

    if set_default_version:
        for schema in schemas.values():
            if isinstance(schema.schema_parsed, Object) and "_version" in schema.schema_parsed.properties:
                schema.schema_parsed.properties["_version"].default = target_version
        logger.info("Set default versions to %s", target_version)

    if clear_output and output.exists():
        shutil.rmtree(output)
        logger.info("Cleared output directory")
    for schema in schemas.values():
        schema.save()
        logger.info("Saved %s", schema.file_path)

    if set_default_version:
        for schema in schemas.values():
            if isinstance(schema.schema_parsed, Object) and "_version" in schema.schema_parsed.properties:
                schema.schema_parsed.properties["_version"].default = target_version
        logger.info("Set default versions to %s", target_version)

    for schema in schemas.values():
        schema.save()
        logger.info("Saved %s", schema.file_path)


if __name__ == "__main__":
    main_command_line()
