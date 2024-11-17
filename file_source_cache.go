package altsrc

type FileSourceCache[T any] struct {
	file string
	m    *T
	f    func(string, any) error
}

func NewFileSourceCache[T any](file string, f func(string, any) error) *FileSourceCache[T] {
	return &FileSourceCache[T]{
		file: file,
		f:    f,
	}
}

func (fsc *FileSourceCache[T]) Get() T {
	if fsc.m == nil {
		res := new(T)
		if err := fsc.f(fsc.file, res); err == nil {
			fsc.m = res
		} else {
			tracef("failed to unmarshal from file %[1]q: %[2]v", fsc.file, err)
		}
	}

	if fsc.m == nil {
		tracef("returning empty")

		return *(new(T))
	}

	return *fsc.m
}

type MapAnyAnyFileSourceCache = FileSourceCache[map[any]any]

func NewMapAnyAnyFileSourceCache(file string, f func(string, any) error) *MapAnyAnyFileSourceCache {
	return NewFileSourceCache[map[any]any](file, f)
}
