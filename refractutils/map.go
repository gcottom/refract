package refractutils

import (
	"errors"
	"reflect"
)

func GetMapIndex(m any, key any) (any, error) {
	val := reflect.ValueOf(m)
	if val.Kind() == reflect.Map {
		return val.MapIndex(reflect.ValueOf(key)), nil
	}
	return nil, errors.New("map argument was not a map")
}

func GetMapIndexValue(m any, key any) (any, error) {
	val := reflect.ValueOf(m)
	if val.Kind() == reflect.Map {
		keyValPtr := val.MapIndex(reflect.ValueOf(key)).Interface()
		return reflect.ValueOf(keyValPtr).Elem().Interface(), nil
	}
	return nil, errors.New("map argument was not a map")
}

func PutMapIndex(m any, key any, value any) error {
	val := reflect.ValueOf(m)
	if val.Kind() == reflect.Map {
		nval := reflect.ValueOf(&value)
		if nval.Kind() == reflect.Ptr {
			nval = nval.Elem()
		}
		val.SetMapIndex(reflect.ValueOf(key), nval)
		return nil
	}
	return errors.New("map argument was not a map")
}
