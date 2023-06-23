package altsrc_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"
)

var (
	tmpDir string
)

func init() {
	if err := setupYAMLValueSourceExamples(); err != nil {
		panic(err)
	}
}

func setupYAMLValueSourceExamples() error {
	tmpDir = filepath.Join(os.TempDir(), "urfave-cli-altsrc-examples")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(tmpDir, "config.yaml"), []byte(`
greet:
  enthusiasm: 9001
`), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(tmpDir, "alt-config.yaml"), []byte(`
greet:
  name: Berry
  enthusiasm: eleven
`), 0644); err != nil {
		return err
	}

	return nil
}

func ExampleYAMLValueSource() {
	configFiles := []string{
		filepath.Join(tmpDir, "config.yaml"),
		filepath.Join(tmpDir, "alt-config.yaml"),
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

	// Output:
	// Hello, Berry!
}
