package altsrc

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
