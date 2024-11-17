package yaml

import (
	"fmt"

	altsrc "github.com/urfave/cli-altsrc/v3"
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
				file:   path,
				key:    key,
				maafsc: altsrc.NewMapAnyAnyFileSourceCache(path, yamlUnmarshalFile),
			},
		)
	}

	return vsc
}

type yamlValueSource struct {
	file string
	key  string

	maafsc *altsrc.MapAnyAnyFileSourceCache
}

func (yvs *yamlValueSource) Lookup() (string, bool) {
	if v, ok := altsrc.NestedVal(yvs.key, yvs.maafsc.Get()); ok {
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
	b, err := altsrc.ReadURI(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, container); err != nil {
		return err
	}

	return nil
}
