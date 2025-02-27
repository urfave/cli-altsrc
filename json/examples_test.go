package json

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"
)

func ExampleJSON() {
	configFiles := []altsrc.Sourcer{
		altsrc.StringSourcer(filepath.Join(testdataDir, "config.json")),
		altsrc.StringSourcer(filepath.Join(testdataDir, "alt-config.json")),
	}

	app := &cli.Command{
		Name: "greet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Sources: cli.NewValueSourceChain(JSON("greet.name", configFiles[0]), JSON("greet.name", configFiles[1])),
			},
			&cli.IntFlag{
				Name:    "enthusiasm",
				Aliases: []string{"!"},
				Sources: cli.NewValueSourceChain(JSON("greet.enthusiasm", configFiles[0]), JSON("greet.enthusiasm", configFiles[1])),
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			punct := ""
			if cmd.Int("enthusiasm") > 9000 {
				punct = "!"
			}

			fmt.Fprintf(os.Stdout, "Hello, %[1]v%[2]v\n", cmd.String("name"), punct)

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
