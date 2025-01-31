package altsrc

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadURI(t *testing.T) {
	tests := []struct {
		name string
		path string
		err  string
	}{
		{
			name: "Empty path",
			err:  fmt.Sprintf("%s: unable to determine how to load from \"\"", Err),
		},
		{
			name: "Invalid file path",
			path: filepath.Join(testdataDir, "action.zzzz"),
			err:  fmt.Sprintf("%s: cannot read from", Err),
		},
		{
			name: "valid file path",
			path: filepath.Join(testdataDir, "alt-config.json"),
		},
		{
			name: "Unknown URI",
			path: "ftp://github.com",
			err:  "is unsupported",
		},
		{
			name: "Invalid http URL",
			path: "http://foo",

			// locally we get: "dial tcp: lookup foo: no such host"
			// but on CI local networks are disabled,
			// so the error is: "dial tcp: lookup foo on 127.0.0.11:53: server misbehaving"
			// therefore let's check for the "lookup foo", which is in both errors
			err: "lookup foo",
		},
		{
			name: "valid http URL",
			path: "http://github.com",
		},
		{
			name: "valid https URL",
			path: "https://github.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := readURI(test.path)
			if test.err != "" {
				assert.ErrorContains(t, err, test.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestURISourceCache(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		contents     string
		unmarshaller func([]byte, any) error
		m            map[any]any
	}{
		{
			name: "no file",
		},
		{
			name: "invalid file",
			path: "/path/to/sssss/junk/file",
		},
		{
			name:     "invalid format",
			contents: "junkddddd",
			unmarshaller: func(b []byte, a any) error {
				return fmt.Errorf("invalid format")
			},
		},
		{
			name:     "invalid format return nil",
			contents: "junkddddd",
			unmarshaller: func(b []byte, a any) error {
				if m, ok := a.(*map[any]any); !ok {
					return fmt.Errorf("not correct map")
				} else {
					*m = nil
				}
				return nil
			},
		},
		{
			name:     "valid format",
			contents: "foo.bar.z=1",
			unmarshaller: func(b []byte, a any) error {
				if m, ok := a.(*map[any]any); !ok {
					return fmt.Errorf("not correct map")
				} else {
					*m = map[any]any{}
					(*m)["foo"] = map[any]any{
						"bar": map[any]any{
							"z": 1,
						},
					}
				}
				return nil
			},
			m: map[any]any{
				"foo": map[any]any{
					"bar": map[any]any{
						"z": 1,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := test.path
			if test.contents != "" {
				fp, err := os.CreateTemp(t.TempDir(), test.name)
				assert.NoError(t, err)
				path = fp.Name()
				defer fp.Close()
			}
			sc := NewMapAnyAnyURISourceCache(path, test.unmarshaller)
			assert.Equal(t, test.m, sc.Get())
		})
	}
}
