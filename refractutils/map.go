package refractutils

import (
	"errors"

	"github.com/gcottom/refract/safereflect"
)

func GetMapIndex(m any, key any) (any, error) {
	val := safereflect.ValueOf(m)
	if val.Kind() == safereflect.Map {
		return val.MapIndex(safereflect.ValueOf(key))
	}
	return nil, errors.New("map argument was not a map")
}

func GetMapIndexValue(m any, key any) (any, error) {
	val := safereflect.ValueOf(m)
	if val.Kind() == safereflect.Map {
		keyValPtr, err := val.MapIndex(safereflect.ValueOf(key))
		if err != nil {
			return nil, err
		}
		keyValPtrI, err := keyValPtr.Interface()
		if err != nil {
			return nil, err
		}
		return safereflect.ValueOf(keyValPtrI).Interface()
	}
	return nil, errors.New("map argument was not a map")
}

func PutMapIndex(m any, key any, value any) error {
	var err error
	val := safereflect.ValueOf(m)
	if val.Kind() == safereflect.Map {
		nval := safereflect.ValueOf(value)
		if nval.Kind() == safereflect.Pointer {
			nval, err = nval.Elem()
			if err != nil {
				return err
			}
		}
		return val.SetMapIndex(safereflect.ValueOf(key), nval)
	}
	return errors.New("map argument was not a map")
}
