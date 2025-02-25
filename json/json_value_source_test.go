package json

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	altsrc "github.com/urfave/cli-altsrc/v3"
	yaml "github.com/urfave/cli-altsrc/yaml"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return altsrc.MustTestdataDir(ctx)
	}()
)

func TestJSON(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "test_config.json")
	altConfigPath := filepath.Join(testdataDir, "test_alt_config.json")

	vsc := yaml.YAML(
		"water_fountain.water",
		altsrc.StringSourcer("/dev/null/nonexistent.json"),
		altsrc.StringSourcer(configPath),
		altsrc.StringSourcer(altConfigPath),
	)
	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)

	yvs := vsc.Chain[0]
	r.Equal("yaml file \"/dev/null/nonexistent.json\" at key \"water_fountain.water\"", yvs.String())
	r.Equal("yamlValueSource{file:\"/dev/null/nonexistent.json\",keyPath:\"water_fountain.water\"}", yvs.GoString())
}
