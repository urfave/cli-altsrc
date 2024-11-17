module github.com/urfave/cli-altsrc/v3

go 1.23.2

require (
	github.com/urfave/cli-altsrc/json v0.0.1
	github.com/urfave/cli-altsrc/toml v0.0.1
	github.com/urfave/cli-altsrc/yaml v0.0.1
	github.com/urfave/cli/v3 v3.0.0-alpha9.3
)

require (
	github.com/BurntSushi/toml v1.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/urfave/cli-altsrc/yaml v0.0.1 => ./yaml/

replace github.com/urfave/cli-altsrc/json v0.0.1 => ./json/

replace github.com/urfave/cli-altsrc/toml v0.0.1 => ./toml/
