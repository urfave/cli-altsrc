package yaml

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return altsrc.MustTestdataDir(ctx)
	}()
)

func TestYAML(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "test_config.yaml")
	altConfigPath := filepath.Join(testdataDir, "test_alt_config.yaml")

	vs1 := YAML(
		"water_fountain.water",
		altsrc.StringSourcer("/dev/null/nonexistent.yaml"),
	)
	vs2 := YAML(
		"water_fountain.water",
		altsrc.StringSourcer(configPath),
	)
	vs3 := YAML(
		"water_fountain.water",
		altsrc.StringSourcer(altConfigPath),
	)

	vsc := cli.NewValueSourceChain(vs1, vs2, vs3)

	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)

	yvs := vsc.Chain[0]
	r.Equal("yaml file \"/dev/null/nonexistent.yaml\" at key \"water_fountain.water\"", yvs.String())
	r.Equal("yamlValueSource{file:\"/dev/null/nonexistent.yaml\",keyPath:\"water_fountain.water\"}", yvs.GoString())
}

func TestYAMLSlice(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "slice.yaml")

	t.Run("numbers", func(t *testing.T) {
		vs := YAML(
			"slice_types.numbers",
			altsrc.StringSourcer(configPath),
		)

		vsc := cli.NewValueSourceChain(vs)

		v, ok := vsc.Lookup()
		r.Equal("1,2,3", v)
		r.True(ok)
	})

	t.Run("strings", func(t *testing.T) {
		vs := YAML(
			"slice_types.strings",
			altsrc.StringSourcer(configPath),
		)

		vsc := cli.NewValueSourceChain(vs)

		v, ok := vsc.Lookup()
		r.Equal("word,this is a sentence", v)
		r.True(ok)
	})
}
