package altsrc_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli-altsrc/v3/internal"
	"github.com/urfave/cli/v3"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return internal.MustTestdataDir(ctx)
	}()
)

func ExampleYAML() {
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

	// Output:
	// Hello, Berry!
}

func ExampleJSON() {
	configFiles := []string{
		filepath.Join(testdataDir, "config.json"),
		filepath.Join(testdataDir, "alt-config.json"),
	}

	app := &cli.Command{
		Name: "greet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Sources: altsrc.JSON("greet.name", configFiles...),
			},
			&cli.IntFlag{
				Name:    "enthusiasm",
				Aliases: []string{"!"},
				Sources: altsrc.JSON("greet.enthusiasm", configFiles...),
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

func ExampleTOML() {
	configFiles := []string{
		filepath.Join(testdataDir, "config.toml"),
		filepath.Join(testdataDir, "alt-config.toml"),
	}

	app := &cli.Command{
		Name: "greet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Sources: altsrc.TOML("greet.name", configFiles...),
			},
			&cli.IntFlag{
				Name:    "enthusiasm",
				Aliases: []string{"!"},
				Sources: altsrc.TOML("greet.enthusiasm", configFiles...),
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
