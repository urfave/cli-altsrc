package altsrc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

// JSON is a helper function that wraps the YAML helper function
// and loads via yaml.Unmarshal
func JSON(key string, paths ...string) cli.ValueSourceChain {
	return YAML(key, paths...)
}

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
				ymc:  yamlMapInputSourceCache{file: path},
			},
		)
	}

	return vsc
}

type yamlMapInputSourceCache struct {
	file string
	m    *map[any]any
}

func (ymc *yamlMapInputSourceCache) Get() map[any]any {
	if ymc.m == nil {
		res := map[any]any{}
		if err := yamlUnmarshalFile(ymc.file, &res); err == nil {
			ymc.m = &res
		} else {
			tracef("failed to unmarshal yaml from file %[1]q: %[2]v", ymc.file, err)
		}
	}

	if ymc.m == nil {
		tracef("returning empty map")
		return map[any]any{}
	}

	return *ymc.m
}

type yamlValueSource struct {
	file string
	key  string

	ymc yamlMapInputSourceCache
}

func (yvs *yamlValueSource) Lookup() (string, bool) {
	if v, ok := nestedVal(yvs.key, yvs.ymc.Get()); ok {
		return fmt.Sprintf("%[1]v", v), ok
	}

	return "", false
}

func (yvs *yamlValueSource) String() string {
	return fmt.Sprintf("yaml file %[1]q at key %[2]q", yvs.file, yvs.key)
}

func (yvs *yamlValueSource) GoString() string {
	return fmt.Sprintf("&yamlValueSource{file:%[1]q,keyPath:%[2]q", yvs.file, yvs.key)
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

func readURI(filePath string) ([]byte, error) {
	u, err := url.Parse(filePath)
	if err != nil {
		return nil, err
	}

	if u.Host != "" { // i have a host, now do i support the scheme?
		switch u.Scheme {
		case "http", "https":
			res, err := http.Get(filePath)
			if err != nil {
				return nil, err
			}
			return io.ReadAll(res.Body)
		default:
			return nil, fmt.Errorf("scheme of %s is unsupported", filePath)
		}
	} else if u.Path != "" { // i dont have a host, but I have a path. I am a local file.
		if _, notFoundFileErr := os.Stat(filePath); notFoundFileErr != nil {
			return nil, fmt.Errorf("Cannot read from file: '%s' because it does not exist.", filePath)
		}
		return os.ReadFile(filePath)
	} else if runtime.GOOS == "windows" && strings.Contains(u.String(), "\\") {
		// on Windows systems u.Path is always empty, so we need to check the string directly.
		if _, notFoundFileErr := os.Stat(filePath); notFoundFileErr != nil {
			return nil, fmt.Errorf("Cannot read from file: '%s' because it does not exist.", filePath)
		}
		return os.ReadFile(filePath)
	}

	return nil, fmt.Errorf("unable to determine how to load from path %s", filePath)
}
