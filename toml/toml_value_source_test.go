package toml

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

func TestTOML(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "test_config.toml")
	altConfigPath := filepath.Join(testdataDir, "test_alt_config.toml")

	vs1 := TOML(
		"water_fountain.water",
		altsrc.StringSourcer("/dev/null/nonexistent.toml"),
	)
	vs2 := TOML(
		"water_fountain.water",
		altsrc.StringSourcer(configPath),
	)

	vs3 := TOML(
		"water_fountain.water",
		altsrc.StringSourcer(altConfigPath),
	)

	vsc := cli.NewValueSourceChain(vs1, vs2, vs3)

	v, ok := vsc.Lookup()
	r.Equal("false", v)
	r.True(ok)

	tvs := vsc.Chain[0]
	r.Equal("toml file \"/dev/null/nonexistent.toml\" at key \"water_fountain.water\"", tvs.String())
	r.Equal("tomlValueSource{file:\"/dev/null/nonexistent.toml\",keyPath:\"water_fountain.water\"}", tvs.GoString())
}

func TestTOMLSlice(t *testing.T) {
	r := require.New(t)

	configPath := filepath.Join(testdataDir, "slice.toml")

	t.Run("numbers", func(t *testing.T) {
		vs := TOML(
			"slice_types.numbers",
			altsrc.StringSourcer(configPath),
		)

		vsc := cli.NewValueSourceChain(vs)

		v, ok := vsc.Lookup()
		r.Equal("1,2,3", v)
		r.True(ok)
	})

	t.Run("strings", func(t *testing.T) {
		vs := TOML(
			"slice_types.strings",
			altsrc.StringSourcer(configPath),
		)

		vsc := cli.NewValueSourceChain(vs)

		v, ok := vsc.Lookup()
		r.Equal("word,this is a sentence", v)
		r.True(ok)
	})
}
