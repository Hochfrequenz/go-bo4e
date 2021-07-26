module github.com/hochfrequenz/go-bo4e

go 1.16

require (
	github.com/campoy/jsonenums v0.0.0-20201009151607-0f0230183423 // indirect
	github.com/corbym/gocrest v1.0.5
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/stretchr/testify v1.7.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace (
	github.com/hochfrequenz/go-bo4e/bo => ./bo
	github.com/hochfrequenz/go-bo4e/com => ./com
	github.com/hochfrequenz/go-bo4e/enum => ./enum
)
