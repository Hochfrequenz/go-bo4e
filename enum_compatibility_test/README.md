# Enum Compatibility Test

This test verifies that enum values serialized by the C# [BO4E-dotnet](https://github.com/Hochfrequenz/BO4E-dotnet) library can be correctly deserialized by the Go [go-bo4e](https://github.com/Hochfrequenz/go-bo4e) library.

## How It Works

1. **C# Script** (`serialize_enums.csx`): A C# script using [dotnet-script](https://github.com/dotnet-script/dotnet-script) that serializes all enum values from BO4E-dotnet to JSON using `System.Text.Json` with `JsonStringEnumConverter`.

2. **Go Test** (`go_deserializer_test.go`): A Go test that reads the JSON output and deserializes it into a matching `AllEnumsTest` struct using go-bo4e's enum types.

Both files define an `AllEnumsTest` struct/class with fields for each enum value. The field names follow the pattern `{EnumTypeName}_{MemberName}` and use JSON tags to match the serialized keys.

## Running Locally

```bash
# Install dotnet-script (one-time)
dotnet tool install -g dotnet-script

# From the enum_compatibility_test directory:

# 1. Run the C# serializer
dotnet script serialize_enums.csx > csharp_enum_output.json

# 2. Run the Go test
go test -v ./...
```

Note: If `csharp_enum_output.json` doesn't exist, the Go test will skip instead of failing.

## CI Integration

The GitHub Actions workflow [`.github/workflows/enum_compatibility.yml`](../.github/workflows/enum_compatibility.yml) runs this test automatically on pull requests.

[![Enum Compatibility Test](https://github.com/Hochfrequenz/go-bo4e/actions/workflows/enum_compatibility.yml/badge.svg)](https://github.com/Hochfrequenz/go-bo4e/actions/workflows/enum_compatibility.yml)

## Updating BO4E-dotnet Version

The NuGet package version is specified inline in `serialize_enums.csx`:

```csharp
#r "nuget: Hochfrequenz.BO4Enet, 0.45.1"
```

To update, change the version number in this line.

---

## Instructions for AI Agents

### Extending Tests with New Enums

When adding new enums to go-bo4e that also exist in BO4E-dotnet, follow these steps to extend the compatibility test:

#### Step 1: Add to C# Script

Edit `serialize_enums.csx` and add properties for each enum member in the `AllEnumsTest` class:

```csharp
public NewEnumType NewEnumType_MEMBER_NAME { get; set; } = NewEnumType.MEMBER_NAME;
// Repeat for each enum member
```

The property name pattern is: `{EnumTypeName}_{MemberName}`

#### Step 2: Add to Go Test

Edit `go_deserializer_test.go`:

1. Add the import:
```go
"github.com/hochfrequenz/go-bo4e/enum/newenumtype"
```

2. Add fields to the `AllEnumsTest` struct for each enum member:
```go
NewEnumType_MEMBER_NAME newenumtype.NewEnumType `json:"NewEnumType_MEMBER_NAME"`
// Repeat for each enum member
```

#### Example: Adding "Beispielenum"

**C# (serialize_enums.csx):**
```csharp
public Beispielenum Beispielenum_WERT_EINS { get; set; } = Beispielenum.WERT_EINS;
public Beispielenum Beispielenum_WERT_ZWEI { get; set; } = Beispielenum.WERT_ZWEI;
```

**Go (go_deserializer_test.go):**
```go
// Import
"github.com/hochfrequenz/go-bo4e/enum/beispielenum"

// Struct fields
Beispielenum_WERT_EINS  beispielenum.Beispielenum `json:"Beispielenum_WERT_EINS"`
Beispielenum_WERT_ZWEI  beispielenum.Beispielenum `json:"Beispielenum_WERT_ZWEI"`
```

### Updating to a Newer BO4E-dotnet Version

When BO4E-dotnet releases a new version with new or modified enums:

1. **Update the NuGet version** in `serialize_enums.csx`:
   ```csharp
   #r "nuget: Hochfrequenz.BO4Enet, X.Y.Z"  // Update version here
   ```

2. **Identify new/changed enums** by checking the [BO4E-dotnet releases](https://github.com/Hochfrequenz/BO4E-dotnet/releases) or comparing the `BO4E.ENUM` namespace.

3. **Add new enum members to C# script**: For each new enum type or member, add corresponding properties to the `AllEnumsTest` class in `serialize_enums.csx`.

4. **Add matching fields to Go test**: For each new enum added to C#, add the corresponding import and struct fields in `go_deserializer_test.go`.

5. **Ensure Go enum exists**: Before adding to the test, verify the enum type exists in go-bo4e under `enum/{enumname}/`. If not, create the Go enum first.

### Syncing Enums Between C# and Go

To ensure all enums are covered:

1. **List C# enums**: Check `BO4E.ENUM` namespace in BO4E-dotnet for all available enum types.

2. **List Go enums**: Check `enum/` directory in go-bo4e for all implemented enum packages.

3. **Compare and add missing**: For each enum that exists in both libraries but is not yet in the test:
   - Add all members to `serialize_enums.csx`
   - Add matching fields to `go_deserializer_test.go`

4. **Run the test locally** to verify compatibility before committing.

### Important Notes

- The JSON tag in Go must exactly match the C# property name
- The enum member values (e.g., `WERT_EINS`) are serialized as strings and must match between C# and Go
- Both C# and Go structs must have matching fields for the test to pass
- Keep fields in alphabetical order by enum type name for easier maintenance
- Some enum members may have special characters (e.g., German umlauts) - ensure JSON tags handle these correctly
