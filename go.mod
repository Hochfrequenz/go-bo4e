module github.com/hochfrequenz/go-bo4e

go 1.16

require (
	github.com/corbym/gocrest v1.0.5 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
)

replace (
    github.com/hochfrequenz/go-bo4e/enum => ./enum
    github.com/hochfrequenz/go-bo4e/com => ./com
)
