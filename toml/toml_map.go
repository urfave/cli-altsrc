package toml

import (
	"fmt"
	"reflect"

	altsrc "github.com/urfave/cli-altsrc/v3"
)

type tomlMap struct {
	Map map[any]any
}

func (tm *tomlMap) UnmarshalTOML(i any) error {
	if v, err := unmarshalMap(i); err == nil {
		tm.Map = v
	} else {
		return err
	}

	return nil
}

func unmarshalMap(i any) (ret map[any]any, err error) {
	ret = make(map[any]any)
	m := i.(map[string]any)
	for key, val := range m {
		v := reflect.ValueOf(val)
		switch v.Kind() {
		case reflect.Bool:
			ret[key] = val.(bool)
		case reflect.String:
			ret[key] = val.(string)
		case reflect.Int:
			ret[key] = val.(int)
		case reflect.Int8:
			ret[key] = int(val.(int8))
		case reflect.Int16:
			ret[key] = int(val.(int16))
		case reflect.Int32:
			ret[key] = int(val.(int32))
		case reflect.Int64:
			ret[key] = int(val.(int64))
		case reflect.Uint:
			ret[key] = int(val.(uint))
		case reflect.Uint8:
			ret[key] = int(val.(uint8))
		case reflect.Uint16:
			ret[key] = int(val.(uint16))
		case reflect.Uint32:
			ret[key] = int(val.(uint32))
		case reflect.Uint64:
			ret[key] = int(val.(uint64))
		case reflect.Float32:
			ret[key] = float64(val.(float32))
		case reflect.Float64:
			ret[key] = val.(float64)
		case reflect.Map:
			if tmp, err := unmarshalMap(val); err == nil {
				ret[key] = tmp
			} else {
				return nil, err
			}
		case reflect.Array, reflect.Slice:
			ret[key] = val.([]any)
		default:
			return nil, fmt.Errorf("%[1]w: unsupported type %#[2]v", altsrc.Err, v.Kind())
		}
	}
	return ret, nil
}
