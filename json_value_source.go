package altsrc

import "github.com/urfave/cli/v3"

// JSON is a helper function that wraps the YAML helper function
// and loads via yaml.Unmarshal
func JSON(key string, paths ...string) cli.ValueSourceChain {
	return YAML(key, paths...)
}
