package altsrc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {
	r := require.New(t)

	tmpDir := t.TempDir()

	configPath := filepath.Join(tmpDir, "config.json")
	altConfigPath := filepath.Join(tmpDir, "alt-config.json")

	r.NoError(os.WriteFile(configPath, []byte(`
{
  "water_fountain": {
    "water": false
  },
  "woodstock": {
    "wood": false
  }
}
`), 0644))

	r.NoError(os.WriteFile(altConfigPath, []byte(`
{
  "water_fountain": {
    "water": true
  },
  "phone_booth": {
    "phone": false
  }
}
`), 0644))

	vsc := YAML(
		"water_fountain.water",
		"/dev/null/nonexistent.json",
		configPath,
		altConfigPath,
	)
	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)
}
