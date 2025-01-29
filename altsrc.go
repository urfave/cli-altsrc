package altsrc

import (
	"errors"
	"fmt"
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

// NestedVal returns a value from the given map. The lookup name may be a dot-separated path into the map.
// If that is the case, it will recursively traverse the map based on the '.' delimited sections to find
// a nested value for the key.
func NestedVal(name string, tree map[any]any) (any, bool) {
	sections := strings.Split(name, ".")
	if name == "" || len(sections) == 0 {
		return nil, false
	}

	node := tree

	// traverse into the map based on the dot-separated sections
	if len(sections) >= 2 { // the last section is the value we want, we will return it directly at the end
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
	}

	if val, ok := node[sections[len(sections)-1]]; ok {
		return val, true
	}
	return nil, false
}
