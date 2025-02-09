module github.com/urfave/cli-altsrc/toml

go 1.23.2

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/stretchr/testify v1.9.0
	github.com/urfave/cli-altsrc/v3 v3.0.0-alpha9
	github.com/urfave/cli/v3 v3.0.0-beta1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/urfave/cli-altsrc/v3 v3.0.0-alpha9 => ../
