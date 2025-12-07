# Enum Compatibility Test

This test verifies that enum values serialized by the C# [BO4E-dotnet](https://github.com/Hochfrequenz/BO4E-dotnet) library can be correctly deserialized by the Go [go-bo4e](https://github.com/Hochfrequenz/go-bo4e) library.

## How It Works

1. **C# Script** (`serialize_enums.csx`): A C# script using [dotnet-script](https://github.com/dotnet-script/dotnet-script) that serializes all enum values from BO4E-dotnet to JSON using `System.Text.Json` with `JsonStringEnumConverter`.

2. **Go Test** (`go_deserializer_test.go`): A Go test that reads the JSON output and attempts to deserialize each enum value using go-bo4e's enum types.

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

## CI Integration

The GitHub Actions workflow `.github/workflows/enum_compatibility.yml` runs this test automatically on pull requests.

## Updating BO4E-dotnet Version

The NuGet package version is specified inline in `serialize_enums.csx`:

```csharp
#r "nuget: Hochfrequenz.BO4Enet, 0.45.1"
```

To update, change the version number in this line.

## Instructions for AI Agents: Extending Tests with New Enums

When adding new enums to go-bo4e that also exist in BO4E-dotnet, follow these steps to extend the compatibility test:

### Step 1: Add to C# Script

Edit `serialize_enums.csx` and add properties for each enum member in the `AllEnumsTest` class:

```csharp
public NewEnumType NewEnumType_MEMBER_NAME { get; set; } = NewEnumType.MEMBER_NAME;
// Repeat for each enum member
```

The property name pattern is: `{EnumTypeName}_{MemberName}`

### Step 2: Add to Go Test

Edit `go_deserializer_test.go`:

1. Add the import:
```go
"github.com/hochfrequenz/go-bo4e/enum/newenumtype"
```

2. Add a case in the switch statement:
```go
case "NewEnumType":
    var v newenumtype.NewEnumType
    unmarshalErr = json.Unmarshal(jsonBytes, &v)
```

### Example: Adding "Beispielenum"

**C# (serialize_enums.csx):**
```csharp
public Beispielenum Beispielenum_WERT_EINS { get; set; } = Beispielenum.WERT_EINS;
public Beispielenum Beispielenum_WERT_ZWEI { get; set; } = Beispielenum.WERT_ZWEI;
```

**Go (go_deserializer_test.go):**
```go
// Import
"github.com/hochfrequenz/go-bo4e/enum/beispielenum"

// Switch case
case "Beispielenum":
    var v beispielenum.Beispielenum
    unmarshalErr = json.Unmarshal(jsonBytes, &v)
```

### Notes

- The C# enum type name in the property must match exactly what the Go test expects in the switch statement
- The enum member values (e.g., `WERT_EINS`) are serialized as strings and must match between C# and Go
- If an enum exists in C# but not in Go (or vice versa), it will be skipped with a log message
