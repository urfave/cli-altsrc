package altsrc

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v3"
)

func TOML(key string, paths ...string) cli.ValueSourceChain {
	vsc := cli.ValueSourceChain{Chain: []cli.ValueSource{}}

	for _, path := range paths {
		vsc.Chain = append(
			vsc.Chain,
			&tomlValueSource{
				file: path,
				key:  key,
				tmc:  tomlMapFileSourceCache{file: path, f: tomlUnmarshalFile, nf: newTomlMap},
			},
		)
	}

	return vsc
}

type tomlValueSource struct {
	file string
	key  string

	tmc tomlMapFileSourceCache
}

func (tvs *tomlValueSource) Lookup() (string, bool) {
	if v, ok := nestedVal(tvs.key, tvs.tmc.Get().Map); ok {
		return fmt.Sprintf("%[1]v", v), ok
	}

	return "", false
}

func (tvs *tomlValueSource) String() string {
	return fmt.Sprintf("yaml file %[1]q at key %[2]q", tvs.file, tvs.key)
}

func (tvs *tomlValueSource) GoString() string {
	return fmt.Sprintf("&tomlValueSource{file:%[1]q,keyPath:%[2]q", tvs.file, tvs.key)
}

func tomlUnmarshalFile(filePath string, container any) error {
	b, err := readURI(filePath)
	if err != nil {
		return err
	}

	if err := toml.Unmarshal(b, container); err != nil {
		return err
	}

	return nil
}
