module github.com/hochfrequenz/go-bo4e

go 1.16

require (
	github.com/corbym/gocrest v1.0.5
	github.com/go-playground/validator/v10 v10.11.0
	github.com/shopspring/decimal v1.3.1
	github.com/stretchr/testify v1.7.1
)

replace (
	github.com/hochfrequenz/go-bo4e/bo => ./bo
	github.com/hochfrequenz/go-bo4e/com => ./com
	github.com/hochfrequenz/go-bo4e/enum => ./enum
)
