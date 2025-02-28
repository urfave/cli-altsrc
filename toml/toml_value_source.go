package toml

import (
	"github.com/BurntSushi/toml"
	altsrc "github.com/urfave/cli-altsrc/v3"
)

// TOML is a helper function to encapsulate a number of
// tomlValueSource together as a cli.ValueSourceChain
func TOML(key string, source altsrc.Sourcer) *altsrc.ValueSource {
	return altsrc.NewValueSource(toml.Unmarshal, "toml", key, source)
}
