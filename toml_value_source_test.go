package altsrc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTOML(t *testing.T) {
	r := require.New(t)

	tmpDir := t.TempDir()

	configPath := filepath.Join(tmpDir, "config.toml")
	altConfigPath := filepath.Join(tmpDir, "alt-config.toml")

	r.NoError(os.WriteFile(configPath, []byte(`
[water_fountain]
water = false

[woodstock]
wood = false
`), 0644))

	r.NoError(os.WriteFile(altConfigPath, []byte(`
[water_fountain]
water = true

[phone_booth]
phone = false
`), 0644))

	vsc := TOML(
		"water_fountain.water",
		"/dev/null/nonexistent.toml",
		configPath,
		altConfigPath,
	)
	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)
}
