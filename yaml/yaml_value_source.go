package yaml

import (
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

// YAML is a helper function to encapsulate a number of
// yamlValueSource together as a cli.ValueSourceChain
func YAML(key string, source altsrc.Sourcer) cli.ValueSource {
	return altsrc.NewValueSource(yaml.Unmarshal, "yaml", key, source)
}
