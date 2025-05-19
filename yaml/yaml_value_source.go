package yaml

import (
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/goccy/go-yaml"
)

// YAML is a helper function to encapsulate a number of
// yamlValueSource together as a cli.ValueSourceChain
func YAML(key string, source altsrc.Sourcer) *altsrc.ValueSource {
	return altsrc.NewValueSource(yaml.Unmarshal, "yaml", key, source)
}
