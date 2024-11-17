package altsrc

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	altsrc "github.com/urfave/cli-altsrc/v3"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return altsrc.MustTestdataDir(ctx)
	}()
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
