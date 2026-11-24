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

func TestYAMLLegacyBooleans(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "bool.yaml")

	cases := map[string]string{
		"yes_val": "true", "Yes_val": "true", "YES_val": "true",
		"y_val": "true", "Y_val": "true",
		"on_val": "true", "On_val": "true", "ON_val": "true",
		"no_val": "false", "No_val": "false", "NO_val": "false",
		"n_val": "false", "N_val": "false",
		"off_val": "false", "Off_val": "false", "OFF_val": "false",
		"true_val": "true", "True_val": "true",
		"false_val": "false", "False_val": "false",
		// Quoted scalars remain strings
		"quoted_no": "no", "quoted_yes": "yes",
		// Non-boolean scalars are untouched
		"plain_word": "maybe", "number": "42",
	}

	for key, want := range cases {
		vs := YAML("bools."+key, altsrc.StringSourcer(configPath))
		vsc := cli.NewValueSourceChain(vs)
		v, ok := vsc.Lookup()
		r.True(ok, "key %s should resolve", key)
		r.Equal(want, v, "key %s", key)
	}
}
