package altsrc

type FileSourceCache[T any] struct {
	file         string
	m            *T
	unmarshaller func([]byte, any) error
}

func NewFileSourceCache[T any](file string, f func([]byte, any) error) *FileSourceCache[T] {
	return &FileSourceCache[T]{
		file:         file,
		unmarshaller: f,
	}
}

func (fsc *FileSourceCache[T]) Get() T {
	if fsc.m == nil {
		res := new(T)
		if b, err := ReadURI(fsc.file); err != nil {
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

type MapAnyAnyFileSourceCache = FileSourceCache[map[any]any]

func NewMapAnyAnyFileSourceCache(file string, f func([]byte, any) error) *MapAnyAnyFileSourceCache {
	return NewFileSourceCache[map[any]any](file, f)
}
