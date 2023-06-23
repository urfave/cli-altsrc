package altsrc

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTOML(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "test_config.toml")
	altConfigPath := filepath.Join(testdataDir, "test_alt_config.toml")

	vsc := TOML(
		"water_fountain.water",
		"/dev/null/nonexistent.toml",
		configPath,
		altConfigPath,
	)
	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)

	tvs := vsc.Chain[0].(*tomlValueSource)
	r.Equal("toml file \"/dev/null/nonexistent.toml\" at key \"water_fountain.water\"", tvs.String())
	r.Equal("&tomlValueSource{file:\"/dev/null/nonexistent.toml\",keyPath:\"water_fountain.water\"}", tvs.GoString())
}
