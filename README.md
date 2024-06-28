# Welcome to `urfave/cli-altsrc/v3`

[![Run Tests](https://github.com/urfave/cli-altsrc/actions/workflows/main.yml/badge.svg)](https://github.com/urfave/cli-altsrc/actions/workflows/main.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/urfave/cli-altsrc/v3.svg)](https://pkg.go.dev/github.com/urfave/cli-altsrc/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/urfave/cli-altsrc/v3)](https://goreportcard.com/report/github.com/urfave/cli-altsrc/v3)

[`urfave/cli-altsrc/v3`](https://pkg.go.dev/github.com/urfave/cli-altsrc/v3) is an extended value source integration library for [`urfave/cli/v3`] with support for JSON,
YAML, and TOML. The primary reason for this to be a separate library is that third-party libraries are used for these
features which are otherwise not used throughout [`urfave/cli/v3`].

[`urfave/cli/v3`]: https://github.com/urfave/cli

### Example

```go
configFiles := []string{
	filepath.Join(testdataDir, "config.yaml"),
	filepath.Join(testdataDir, "alt-config.yaml"),
}

app := &cli.Command{
	Name: "greet",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Sources: altsrc.YAML("greet.name", configFiles...),
		},
		&cli.IntFlag{
			Name:    "enthusiasm",
			Aliases: []string{"!"},
			Sources: altsrc.YAML("greet.enthusiasm", configFiles...),
		},
	},
	Action: func(cCtx *cli.Context) error {
		punct := ""
		if cCtx.Int("enthusiasm") > 9000 {
			punct = "!"
		}

		fmt.Fprintf(os.Stdout, "Hello, %[1]v%[2]v\n", cCtx.String("name"), punct)

		return nil
	},
}

// Simulating os.Args
os.Args = []string{"greet"}

if err := app.Run(context.Background(), os.Args); err != nil {
	fmt.Fprintf(os.Stdout, "OH NO: %[1]v\n", err)
}
```
