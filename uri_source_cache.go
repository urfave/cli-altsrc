package altsrc

type URISourceCache[T any] struct {
	file         string
	m            *T
	unmarshaller func([]byte, any) error
}

func NewURISourceCache[T any](file string, f func([]byte, any) error) *URISourceCache[T] {
	return &URISourceCache[T]{
		file:         file,
		unmarshaller: f,
	}
}

func (fsc *URISourceCache[T]) Get() T {
	if fsc.m == nil {
		res := new(T)
		if b, err := readURI(fsc.file); err != nil {
			tracef("failed to read uri %[1]q: %[2]v", fsc.file, err)
		} else if err := fsc.unmarshaller(b, res); err != nil {
			tracef("failed to unmarshal from file %[1]q: %[2]v", fsc.file, err)
		} else {
			fsc.m = res
		}
	}

	if fsc.m == nil {
		tracef("returning empty")

		return *(new(T))
	}

	return *fsc.m
}

type MapAnyAnyURISourceCache = URISourceCache[map[any]any]

func NewMapAnyAnyURISourceCache(file string, f func([]byte, any) error) *MapAnyAnyURISourceCache {
	return NewURISourceCache[map[any]any](file, f)
}
