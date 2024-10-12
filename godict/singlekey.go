package godict

import (
	"encoding/json"
	"errors"
	"reflect"
)

func UnmarshalSingleJSONKey[T any](key string, data []byte) (T, error) {
	m := make(map[string]any)
	if err := json.Unmarshal(data, &m); err != nil {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), err
	}
	val, ok := m[key]
	if !ok {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), errors.New("key not found")
	}
	if out, ok := val.(T); ok {
		return out, nil
	} else {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), errors.New("value not of type specified")
	}

}
