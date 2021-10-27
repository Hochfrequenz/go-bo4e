# go-bo4e

![Unittest status badge](https://github.com/Hochfrequenz/go-bo4e/workflows/Unittests/badge.svg)
![Coverage status badge](https://github.com/Hochfrequenz/go-bo4e/workflows/coverage/badge.svg)
![Linter status badge](https://github.com/Hochfrequenz/go-bo4e/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hochfrequenz/go-bo4e)](https://goreportcard.com/report/github.com/Hochfrequenz/go-bo4e)
[![Go Reference](https://pkg.go.dev/badge/github.com/hochfrequenz/go-bo4e.svg)](https://pkg.go.dev/github.com/hochfrequenz/go-bo4e)

**B**usiness **O**bjects for **E**nergy ([BO4E](https://www.bo4e.de/)) Implementation in Go. Highlights are

- includes `json` tags for easy (un)marshalling
- comes with builtin [validator](https://github.com/go-playground/validator) logic
- is linted and has decent test coverage

**This package (as of September 2021) implements**

- BO Energiemenge
- BO Geschaeftspartner
- BO Lastgang
- BO Marktteilnehmer
- BO Messlokation
- BO Rechnung
- BO Vertrag
- BO Zaehler

so far.

## Installation

```
go get github.com/hochfrequenz/go-bo4e
```

### Breaking Changes introduced in v0.0.13 and v0.0.14:

- The struct that is embedded in all BusinessObjects is now called `Geschaeftsobjekt` (was `BusinessObject` <=v0.0.12) to be consistent with the official documentation
- `BusinessObject` is now the name of the interface that all structs with embedded `Geschaeftsobjekt` implement
- Enums are consistently written in upper case ([PR 32](https://github.com/Hochfrequenz/go-bo4e/pull/32))

## Minimal Working Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

func main() {
	melo := bo.Messlokation{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.MESSLOKATION,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
		},
		MesslokationsId: "DE0000011111222223333344444555556",
		Sparte:          sparte.STROM,
		Messadresse: &com.Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Strasse:      "Nördliche Münchner Straße",
			Hausnummer:   "27A",
			Landescode:   landescode.DE,
		},
	}
	vali := validator.New()
	err := vali.Struct(melo)
	if err == nil {
		fmt.Println("The MeLo is valid.")
	}
	meloBytes, err := json.Marshal(melo)
	meloJson := string(meloBytes)
	fmt.Print(meloJson)
}
```

[Use this in the Go Playground](https://play.golang.org/p/wq8B_31Odni).

## Other Noteworthy BO4E Implementations

- [C# / .NET](https://github.com/Hochfrequenz/BO4E-dotnet/)
- [Python](https://github.com/Hochfrequenz/BO4E-python/)
- [Kotlin](https://github.com/openEnWi/ktBO4E-lib)

## Contributing

Contributions are welcome. Feel free to open a Pull Request against the main branch of this repository. Please provide
unit tests if you contribute logic beyond bare bare business object definitions. We do track our modification proposals
to the official BO4E standard in a separate
repository: [BO4E-modification-proposals](https://github.com/Hochfrequenz/bo4e-modification-proposals).

### Adding Enums
Wenn adding Enums there are two packages ([stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer), [jsonenums](https://github.com/campoy/jsonenums)) needed to generate additional files.
Since they are just needed for the generation, but not for running the program we don't want them in the go.mod file.
One way to install them is outside of your directory with:
```
go install github.com/campoy/jsonenums@latest
go install golang.org/x/tools/cmd/stringer@latest
```


## Hochfrequenz

[Hochfrequenz Unternehmensberatung GmbH](https://www.hochfrequenz.de) is a Grünwald (near Munich) based consulting
company with offices in Berlin and Bremen. We're not only the main contributor to open source BO4E software but,
according to [Kununu ratings](https://www.kununu.com/de/hochfrequenz-unternehmensberatung1), also among the most
attractive employers within the German energy market. Applications of talented developers are welcome at any time!
Please consider visiting
our [career page](https://www.hochfrequenz.de/index.php/karriere/aktuelle-stellenausschreibungen/full-stack-entwickler) (
German only).
