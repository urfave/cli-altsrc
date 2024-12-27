package json

import (
	yaml "github.com/urfave/cli-altsrc/yaml"
	"github.com/urfave/cli/v3"
)

// JSON is a helper function that wraps the YAML helper function
// and loads via yaml.Unmarshal
func JSON(key string, paths ...string) cli.ValueSourceChain {
	return yaml.YAML(key, paths...)
}
