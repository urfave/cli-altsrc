module github.com/urfave/cli-altsrc/v3

go 1.23.2

require github.com/stretchr/testify v1.9.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/urfave/cli-altsrc/yaml v0.0.1 => ./yaml/

replace github.com/urfave/cli-altsrc/json v0.0.1 => ./json/

replace github.com/urfave/cli-altsrc/toml v0.0.1 => ./toml/
