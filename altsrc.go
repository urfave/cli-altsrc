package altsrc

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

var (
	Err = errors.New("urfave/cli-altsrc error")

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

func readURI(uriString string) ([]byte, error) {
	u, err := url.Parse(uriString)
	if err != nil {
		return nil, err
	}

	if u.Host != "" { // i have a host, now do i support the scheme?
		switch u.Scheme {
		case "http", "https":
			res, err := http.Get(uriString)
			if err != nil {
				return nil, err
			}
			return io.ReadAll(res.Body)
		default:
			return nil, fmt.Errorf("%[1]w: scheme of %[2]q is unsupported", Err, uriString)
		}
	} else if u.Path != "" ||
		(runtime.GOOS == "windows" && strings.Contains(u.String(), "\\")) {
		if _, notFoundFileErr := os.Stat(uriString); notFoundFileErr != nil {
			return nil, fmt.Errorf("%[1]w: cannot read from %[2]q because it does not exist", Err, uriString)
		}
		return os.ReadFile(uriString)
	}

	return nil, fmt.Errorf("%[1]w: unable to determine how to load from %[2]q", Err, uriString)
}

// nestedVal checks if the name has '.' delimiters.
// If so, it tries to traverse the tree by the '.' delimited sections to find
// a nested value for the key.
func nestedVal(name string, tree map[any]any) (any, bool) {
	if sections := strings.Split(name, "."); len(sections) > 1 {
		node := tree
		for _, section := range sections[:len(sections)-1] {
			child, ok := node[section]
			if !ok {
				return nil, false
			}

			switch child := child.(type) {
			case map[string]any:
				node = make(map[any]any, len(child))
				for k, v := range child {
					node[k] = v
				}
			case map[any]any:
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
