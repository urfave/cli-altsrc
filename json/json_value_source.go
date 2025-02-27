package json

import (
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

// JSON is a helper function that wraps the YAML helper function
// and loads via yaml.Unmarshal
func JSON(key string, sources ...altsrc.Sourcer) cli.ValueSourceChain {
	return altsrc.NewValueSourceChain(yaml.Unmarshal, "json", key, sources...)
}
