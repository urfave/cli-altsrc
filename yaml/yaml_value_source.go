package yaml

import (
	altsrc "github.com/urfave/cli-altsrc/v3"
	"go.yaml.in/yaml/v3"
)

// YAML is a helper function to encapsulate a number of
// yamlValueSource together as a cli.ValueSourceChain
func YAML(key string, source altsrc.Sourcer) *altsrc.ValueSource {
	return altsrc.NewValueSource(unmarshal, "yaml", key, source)
}

// legacyBoolScalars maps the YAML 1.1 boolean lexemes to their resolved
// values. go.yaml.in/yaml/v3 follows the YAML 1.2 core schema and resolves
// only "true"/"false", which regressed parsing of the other lexemes that
// earlier versions (and the YAML 1.1 spec cited in #51) treated as booleans.
var legacyBoolScalars = map[string]bool{
	"y": true, "Y": true, "yes": true, "Yes": true, "YES": true,
	"on": true, "On": true, "ON": true,
	"n": false, "N": false, "no": false, "No": false, "NO": false,
	"off": false, "Off": false, "OFF": false,
}

// unmarshal decodes YAML the same way yaml.Unmarshal does, but additionally
// resolves unquoted YAML 1.1 boolean lexemes to booleans so values like
// "no" and "off" reach cli.BoolFlag as "false" instead of failing to parse.
// Quoted scalars and mapping keys are left untouched: a YAML key spelled
// "on" must remain findable by the string "on", and a quoted "no" is a
// string by spec.
func unmarshal(in []byte, out any) error {
	var node yaml.Node
	if err := yaml.Unmarshal(in, &node); err != nil {
		return err
	}
	normalizeLegacyBoolValues(&node)
	return node.Decode(out)
}

// normalizeLegacyBoolValues rewrites plain scalar nodes in value position
// whose text is a YAML 1.1 boolean lexeme into !!bool nodes. Mapping keys
// are intentionally skipped.
func normalizeLegacyBoolValues(n *yaml.Node) {
	switch n.Kind {
	case yaml.ScalarNode:
		if n.Tag == "!!str" && n.Style == 0 {
			if b, ok := legacyBoolScalars[n.Value]; ok {
				n.Tag = "!!bool"
				n.Value = "false"
				if b {
					n.Value = "true"
				}
			}
		}
	case yaml.MappingNode:
		// Content alternates key,value pairs; only the values normalize.
		for i := 1; i < len(n.Content); i += 2 {
			normalizeLegacyBoolValues(n.Content[i])
		}
	default:
		for _, c := range n.Content {
			normalizeLegacyBoolValues(c)
		}
	}
}
