module github.com/hochfrequenz/go-bo4e

go 1.24.0

toolchain go1.24.1

require (
	github.com/corbym/gocrest v1.1.2
	github.com/go-playground/validator/v10 v10.28.0
	github.com/shopspring/decimal v1.4.0
)

require (
	github.com/gabriel-vasile/mimetype v1.4.10 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
)

replace (
	github.com/hochfrequenz/go-bo4e/bo => ./bo
	github.com/hochfrequenz/go-bo4e/com => ./com
	github.com/hochfrequenz/go-bo4e/enum => ./enum
)
