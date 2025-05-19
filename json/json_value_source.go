package json

import (
	altsrc "github.com/urfave/cli-altsrc/v3"
	"go.yaml.in/yaml/v3"
)

// JSON is a helper function that wraps the YAML helper function
// and loads via yaml.Unmarshal
func JSON(key string, source altsrc.Sourcer) *altsrc.ValueSource {
	return altsrc.NewValueSource(yaml.Unmarshal, "json", key, source)
}
