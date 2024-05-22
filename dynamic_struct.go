package refract

import (
	"fmt"
	"reflect"
)

func NewStructField(fieldName string, fieldType any, fieldTag string) reflect.StructField {
	return reflect.StructField{Name: fieldName, Type: reflect.TypeOf(fieldType), Tag: reflect.StructTag(fieldTag)}
}

func NewStructDefinition(fields ...reflect.StructField) reflect.Type {
	return reflect.StructOf(fields)
}

func NewStructInstance(structDefinition reflect.Type) any {
	val := reflect.New(structDefinition)
	return val.Interface()
}

func NewSliceOfStruct(structDefinition reflect.Type) any {
	sTyp := reflect.SliceOf(structDefinition)
	newSlice := reflect.MakeSlice(sTyp, 0, 0)
	return newSlice.Interface()
}

func NewMapOfStruct[T comparable](keyType T, structDefinition reflect.Type) any {
	mTyp := reflect.MapOf(reflect.TypeOf(keyType), structDefinition)
	val := reflect.MakeMap(mTyp)
	return val.Interface()
}

func SetStructFieldValue[T any](structInstance any, fieldName string, fieldValue T) error {
	if reflect.ValueOf(structInstance).Kind() != reflect.Ptr ||
		reflect.ValueOf(structInstance).Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct instance, got %s", reflect.ValueOf(structInstance).Elem().Kind())
	}
	if !reflect.ValueOf(structInstance).Elem().FieldByName(fieldName).IsValid() {
		return fmt.Errorf("field with name: \"%s\" does not exist on struct instance", fieldName)
	}
	if reflect.TypeFor[T]().Kind() != reflect.ValueOf(structInstance).Elem().FieldByName(fieldName).Kind() {
		if reflect.TypeFor[T]().Kind() == reflect.Interface {
			reflect.ValueOf(structInstance).Elem().FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
			return nil
		}
		return fmt.Errorf("field with name: \"%s\" has underlying type: %s, but fieldValue argument has type: %s", fieldName, reflect.ValueOf(structInstance).Elem().FieldByName(fieldName).Kind().String(), reflect.TypeFor[T]().Kind().String())
	}
	reflect.ValueOf(structInstance).Elem().FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
	return nil
}

func GetStructFieldValue[T any](structInstance any, fieldName string) (T, error) {
	val := reflect.ValueOf(structInstance)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), fmt.Errorf("expected a struct instance")
	}
	if !val.FieldByName(fieldName).IsValid() {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), fmt.Errorf("field not valid for struct instance")
	}
	if val.FieldByName(fieldName).Kind() != reflect.TypeFor[T]().Kind() {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), fmt.Errorf("field with name: \"%s\" has underlying type: %s, but generic type assertion was for type: %s", fieldName, val.FieldByName(fieldName).Kind().String(), reflect.TypeFor[T]().Kind().String())
	}
	return val.FieldByName(fieldName).Interface().(T), nil
}

func GetStructFieldValueAny(structInstance any, fieldName string) (any, error) {
	val := reflect.ValueOf(structInstance)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct instance")
	}
	if !val.FieldByName(fieldName).IsValid() {
		return nil, fmt.Errorf("field not valid for struct instance")
	}
	return val.FieldByName(fieldName).Interface(), nil
}
