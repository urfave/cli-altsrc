package json

import (
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

// JSON is a helper function that wraps the YAML helper function
// and loads via yaml.Unmarshal
func JSON(key string, source altsrc.Sourcer) cli.ValueSource {
	return altsrc.NewValueSource(yaml.Unmarshal, "json", key, source)
}
