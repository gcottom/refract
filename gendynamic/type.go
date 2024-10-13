package gendynamic

import (
	"fmt"
	"reflect"

	"github.com/gcottom/refract/safereflect"
)

func Assert[T any](v any) (T, error) {
	out, ok := v.(T)
	if ok {
		return out, nil
	}
	return safereflect.ZeroGeneric[T](), fmt.Errorf("couldn't assert as %s", reflect.TypeFor[T]().String())
}

func GetReflectType(v any) safereflect.Type {
	return safereflect.ValueOf(v).Type()
}

func IsPtr(v any) bool {
	return GetReflectType(v).Kind() == safereflect.Pointer
}

func IsMap(v any) bool {
	if IsPtr(v) {
		return GetReflectType(v).Elem().Kind() == safereflect.Map
	}
	return GetReflectType(v).Kind() == safereflect.Map
}

func IsSlice(v any) bool {
	if IsPtr(v) {
		return GetReflectType(v).Elem().Kind() == safereflect.Slice
	}
	return GetReflectType(v).Kind() == safereflect.Slice
}

func IsStruct(v any) bool {
	if IsPtr(v) {
		return GetReflectType(v).Elem().Kind() == safereflect.Struct
	}
	return GetReflectType(v).Kind() == safereflect.Struct
}
