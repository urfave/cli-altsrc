package altsrc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

var (
	isTracingOn = os.Getenv("URFAVE_CLI_TRACING") == "on"
)

func tracef(format string, a ...any) {
	if !isTracingOn {
		return
	}

	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}

	pc, file, line, _ := runtime.Caller(1)
	cf := runtime.FuncForPC(pc)

	fmt.Fprintf(
		os.Stderr,
		strings.Join([]string{
			"## URFAVE CLI TRACE ",
			file,
			":",
			fmt.Sprintf("%v", line),
			" ",
			fmt.Sprintf("(%s)", cf.Name()),
			" ",
			format,
		}, ""),
		a...,
	)
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

// nestedVal checks if the name has '.' delimiters.
// If so, it tries to traverse the tree by the '.' delimited sections to find
// a nested value for the key.
func nestedVal(name string, tree map[interface{}]interface{}) (interface{}, bool) {
	if sections := strings.Split(name, "."); len(sections) > 1 {
		node := tree
		for _, section := range sections[:len(sections)-1] {
			child, ok := node[section]
			if !ok {
				return nil, false
			}

			switch child := child.(type) {
			case map[string]interface{}:
				node = make(map[interface{}]interface{}, len(child))
				for k, v := range child {
					node[k] = v
				}
			case map[interface{}]interface{}:
				node = child
			default:
				return nil, false
			}
		}
		if val, ok := node[sections[len(sections)-1]]; ok {
			return val, true
		}
	}
	return nil, false
}
