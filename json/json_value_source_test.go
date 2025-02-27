package json

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli-altsrc/v3/yaml"
	"github.com/urfave/cli/v3"
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

	vs1 := yaml.YAML(
		"water_fountain.water",
		altsrc.StringSourcer("/dev/null/nonexistent.json"),
	)
	vs2 := yaml.YAML(
		"water_fountain.water",
		altsrc.StringSourcer(configPath),
	)

	vs3 := yaml.YAML(
		"water_fountain.water",
		altsrc.StringSourcer(altConfigPath),
	)

	vsc := cli.NewValueSourceChain(vs1, vs2, vs3)
	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)

	yvs := vsc.Chain[0]
	r.Equal("yaml file \"/dev/null/nonexistent.json\" at key \"water_fountain.water\"", yvs.String())
	r.Equal("yamlValueSource{file:\"/dev/null/nonexistent.json\",keyPath:\"water_fountain.water\"}", yvs.GoString())
}
