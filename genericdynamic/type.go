package genericdynamic

import (
	"fmt"
	"reflect"
)

func Assert[T any](v any) (T, error) {
	out, ok := v.(T)
	if ok {
		return out, nil
	}
	return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), fmt.Errorf("couldn't assert as %s", reflect.TypeFor[T]().String())
}

func GetReflectType(v any) reflect.Type {
	return reflect.ValueOf(v).Type()
}

func IsPtr(v any) bool {
	return GetReflectType(v).Kind() == reflect.Ptr
}

func IsMap(v any) bool {
	if IsPtr(v) {
		return GetReflectType(v).Elem().Kind() == reflect.Map
	}
	return GetReflectType(v).Kind() == reflect.Map
}

func IsSlice(v any) bool {
	if IsPtr(v) {
		return GetReflectType(v).Elem().Kind() == reflect.Slice
	}
	return GetReflectType(v).Kind() == reflect.Slice
}

func IsStruct(v any) bool {
	if IsPtr(v) {
		return GetReflectType(v).Elem().Kind() == reflect.Struct
	}
	return GetReflectType(v).Kind() == reflect.Struct
}
