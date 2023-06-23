package altsrc

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

// YAML is a helper function to encapsulate a number of
// yamlValueSource together as a cli.ValueSourceChain
func YAML(key string, paths ...string) cli.ValueSourceChain {
	vsc := cli.ValueSourceChain{Chain: []cli.ValueSource{}}

	for _, path := range paths {
		vsc.Chain = append(
			vsc.Chain,
			&yamlValueSource{
				file: path,
				key:  key,
				maafsc: mapAnyAnyFileSourceCache{
					file: path,
					f:    yamlUnmarshalFile,
				},
			},
		)
	}

	return vsc
}

type yamlValueSource struct {
	file string
	key  string

	maafsc mapAnyAnyFileSourceCache
}

func (yvs *yamlValueSource) Lookup() (string, bool) {
	if v, ok := nestedVal(yvs.key, yvs.maafsc.Get()); ok {
		return fmt.Sprintf("%[1]v", v), ok
	}

	return "", false
}

func (yvs *yamlValueSource) String() string {
	return fmt.Sprintf("yaml file %[1]q at key %[2]q", yvs.file, yvs.key)
}

func (yvs *yamlValueSource) GoString() string {
	return fmt.Sprintf("&yamlValueSource{file:%[1]q,keyPath:%[2]q}", yvs.file, yvs.key)
}

func yamlUnmarshalFile(filePath string, container any) error {
	b, err := readURI(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, container); err != nil {
		return err
	}

	return nil
}
