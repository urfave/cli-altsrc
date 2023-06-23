package altsrc

type fileSourceCache[T any] struct {
	file string
	m    *T
	f    func(string, any) error
}

func (fsc *fileSourceCache[T]) Get() T {
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

type mapAnyAnyFileSourceCache = fileSourceCache[map[any]any]

type tomlMapFileSourceCache = fileSourceCache[tomlMap]
