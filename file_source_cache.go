package altsrc

type fileSourceCache[T any] struct {
	file string
	m    *T
	nf   func() T
	f    func(string, any) error
}

func (fsc *fileSourceCache[T]) Get() T {
	if fsc.m == nil {
		res := fsc.nf()
		if err := fsc.f(fsc.file, &res); err == nil {
			fsc.m = &res
		} else {
			tracef("failed to unmarshal from file %[1]q: %[2]v", fsc.file, err)
		}
	}

	if fsc.m == nil {
		tracef("returning empty")
		return fsc.nf()
	}

	return *fsc.m
}

func newMapAnyAny() map[any]any {
	return map[any]any{}
}

type mapAnyAnyFileSourceCache = fileSourceCache[map[any]any]

func newTomlMap() tomlMap {
	return tomlMap{Map: map[any]any{}}
}

type tomlMapFileSourceCache = fileSourceCache[tomlMap]
