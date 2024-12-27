package json_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	json "github.com/urfave/cli-altsrc/json"
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return altsrc.MustTestdataDir(ctx)
	}()
)

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
				Sources: json.JSON("greet.name", configFiles...),
			},
			&cli.IntFlag{
				Name:    "enthusiasm",
				Aliases: []string{"!"},
				Sources: json.JSON("greet.enthusiasm", configFiles...),
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
