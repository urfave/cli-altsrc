package altsrc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return MustTestdataDir(ctx)
	}()
)

func TestNestedVal(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		m     map[any]any
		val   any
		found bool
	}{
		{
			name: "No map no key",
		},
		{
			name: "No map with key",
			key:  "foo",
		},
		{
			name: "Empty map no key",
			m:    map[any]any{},
		},
		{
			name: "Empty map with key",
			key:  "foo",
			m:    map[any]any{},
		},
		{
			name: "Level 1 no key",
			key:  ".foob",
			m: map[any]any{
				"foo": 10,
			},
		},
		{
			name: "Level 1",
			key:  "foobar",
			m: map[any]any{
				"foobar": 10,
			},
			val:   10,
			found: true,
		},
		{
			name: "Level 2",
			key:  "foo.bar",
			m: map[any]any{
				"foo": map[any]any{
					"bar": 10,
				},
			},
			val:   10,
			found: true,
		},
		{
			name: "Level 2 invalid key",
			key:  "foo.bar1",
			m: map[any]any{
				"foo": map[any]any{
					"bar": 10,
				},
			},
		},
		{
			name: "Level 2 string map type",
			key:  "foo.bar1",
			m: map[any]any{
				"foo": map[string]any{
					"bar": "10",
				},
			},
		},
		{
			name: "Level 3 no entry",
			key:  "foo.bar.t",
			m: map[any]any{
				"foo": map[any]any{
					"bar": "sss",
				},
			},
		},
		{
			name: "Level 3",
			key:  "foo.bar.t",
			m: map[any]any{
				"foo": map[any]any{
					"bar": map[any]any{
						"t": "sss",
					},
				},
			},
			val:   "sss",
			found: true,
		},
		{
			name: "Level 3 invalid key",
			key:  "foo.bar.t",
			m: map[any]any{
				"foo": map[any]any{
					"bar": map[any]any{
						"t1": 10,
					},
				},
			},
		},
		{
			name: "Level 4 no entry",
			key:  "foo.bar.t.gh",
			m: map[any]any{
				"foo": map[any]any{
					"bar": map[any]any{
						"t1": 10,
					},
				},
			},
		},
		{
			name: "Level 4 slice entry",
			key:  "foo.bar.t.gh",
			m: map[any]any{
				"foo": map[any]any{
					"bar": map[any]any{
						"t": map[any]any{
							"gh": []int{10},
						},
					},
				},
			},
			val:   []int{10},
			found: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, b := NestedVal(test.key, test.m)
			if !test.found {
				assert.False(t, b)
			} else {
				assert.True(t, b)
				assert.Equal(t, val, test.val)
			}
		})
	}
}
