package altsrc

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "test_config.json")
	altConfigPath := filepath.Join(testdataDir, "test_alt_config.json")

	vsc := YAML(
		"water_fountain.water",
		"/dev/null/nonexistent.json",
		configPath,
		altConfigPath,
	)
	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)

	yvs := vsc.Chain[0].(*yamlValueSource)
	r.Equal("yaml file \"/dev/null/nonexistent.json\" at key \"water_fountain.water\"", yvs.String())
	r.Equal("&yamlValueSource{file:\"/dev/null/nonexistent.json\",keyPath:\"water_fountain.water\"}", yvs.GoString())
}
