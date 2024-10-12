package genericdynamic

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// NewStructField is used to create new struct fields to be used with NewStructDefinition as arguments. This function
// takes a fieldName string, fieldType any, and fieldTag string as arguments. The field must be exported, therefore the
// first letter of the fieldName will automatically be capitalized if it is not already capitalized. fieldType should
// be a var with the type that the field should hold. For string type, it would be "string", for int, it would be an
// int like 1 or 2. fieldTag takes a string which contains the tags that would normally be in a struct tag. For example `json:"fieldName"`.
func NewStructField(fieldName string, fieldType any, fieldTag string) reflect.StructField {
	fieldName = strings.ReplaceAll(fieldName, " ", "")
	if !unicode.IsUpper(rune(fieldName[0])) {
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(rune(fieldName[0]))
		fieldName = string(r)
	}

	return reflect.StructField{Name: fieldName, Type: reflect.TypeOf(fieldType), Tag: reflect.StructTag(fieldTag)}
}

// NewStructFieldWithReflectTag is used to create new struct fields to be used with NewStructDefinition as arguments. This function
// takes a fieldName string, fieldType reflect.Type, and fieldTag reflect.StructTag as arguments. The field must be exported, therefore the
// first letter of the fieldName will automatically be capitalized if it is not already capitalized. fieldType should
// be a var with the type that the field should hold. For string type, it would be "string", for int, it would be an
// int like 1 or 2. fieldTag takes a reflect.StructTag.
func NewStructFieldWithReflectTag(fieldName string, fieldType any, fieldTag reflect.StructTag) reflect.StructField {
	fieldName = strings.ReplaceAll(fieldName, " ", "")
	if !unicode.IsUpper(rune(fieldName[0])) {
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(rune(fieldName[0]))
		fieldName = string(r)
	}

	return reflect.StructField{Name: fieldName, Type: reflect.TypeOf(fieldType), Tag: fieldTag}
}

// NewStructFieldWithReflectTagAndType is used to create new struct fields to be used with NewStructDefinition as arguments. This function
// takes a fieldName string, fieldType reflect.Type, and fieldTag reflect.StructTag as arguments. The field must be exported, therefore the
// first letter of the fieldName will automatically be capitalized if it is not already capitalized. fieldType should
// be a reflect.Type, it can also be a struct definition created by refract. fieldTag takes a reflect.StructTag.
func NewStructFieldWithReflectTagAndType(fieldName string, fieldType reflect.Type, fieldTag reflect.StructTag) reflect.StructField {
	fieldName = strings.ReplaceAll(fieldName, " ", "")
	if !unicode.IsUpper(rune(fieldName[0])) {
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(rune(fieldName[0]))
		fieldName = string(r)
	}

	return reflect.StructField{Name: fieldName, Type: fieldType, Tag: fieldTag}
}

// NewStructFieldWithReflectType is used to create new struct fields to be used with NewStructDefinition as arguments. This function
// takes a fieldName string, fieldType reflect.Type, and fieldTag string as arguments. The field must be exported, therefore the
// first letter of the fieldName will automatically be capitalized if it is not already capitalized. fieldType should
// be a reflect.Type, it can also be a struct definition created by refract. fieldTag takes a string which contains the tags that
// would normally be in a struct tag. For example `json:"fieldName"`.
func NewStructFieldWithReflectType(fieldName string, fieldType reflect.Type, fieldTag string) reflect.StructField {
	fieldName = strings.ReplaceAll(fieldName, " ", "")
	if !unicode.IsUpper(rune(fieldName[0])) {
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(rune(fieldName[0]))
		fieldName = string(r)
	}

	return reflect.StructField{Name: fieldName, Type: fieldType, Tag: reflect.StructTag(fieldTag)}
}

// NewStructDefinition takes a variadic of fields which are reflect.StructField. reflect.StructField can be created by
// calling the NewStructField function. NewStructDefinition creates a reflect.Type which can be used in the NewTypeInstance,
// NewSliceOfType, and NewMapOfType functions.
func NewStructDefinition(fields ...reflect.StructField) reflect.Type {
	return reflect.StructOf(fields)
}

// NewTypeInstance takes a typeDefinition as an argument. typeDefinition is a reflect.Type which can be created by using the
// NewStructDefinition function. NewTypeInstance creates a pointer to an instance of the type specified. By using this function,
// you can create dynamic and generic instances of a type. You can also use this function with any reflect.Type to create a new
// instance of a native type.
func NewTypeInstance(typeDefinition reflect.Type) any {
	if typeDefinition.Kind() == reflect.Ptr {
		typeDefinition = typeDefinition.Elem()
	}
	val := reflect.New(typeDefinition)
	return val.Interface()
}

// NewSliceOfType is used with a typeDefinition. typeDefinition is a reflect.Type which can be created by using the NewStructDefinition
// function. NewSliceOfType creates a pointer to a slice of the type specified. This function can be used to create slices of dynamic types.
// This function can be used with any reflect.Type.
func NewSliceOfType(typeDefinition reflect.Type) any {
	si := NewTypeInstance(typeDefinition)
	sd := reflect.ValueOf(&si).Type().Elem()
	sTyp := reflect.SliceOf(sd)
	newSlice := reflect.MakeSlice(sTyp, 0, 0)
	return newSlice.Interface()
}

// NewMapOfType is used with a keyType (comparable) and a typeDefinition. keyType must implement the comparable interface. typeDefinition is a reflect.Type
// which can be created by using the NewStructDefinition function. NewMapOfType creates a pointer to a map of the type specified.
// This function can create instances of maps of dynamic types. This function can be used with any reflect.Type.
func NewMapOfType[T comparable](keyType T, typeDefinition reflect.Type) any {
	si := NewTypeInstance(typeDefinition)
	sd := reflect.ValueOf(&si).Type().Elem()
	mTyp := reflect.MapOf(reflect.TypeOf(keyType), sd)
	newMap := reflect.MakeMap(mTyp)
	return newMap.Interface()
}

// NewMapOfTypeWithReflectTypeKey is used with a keyType reflect.Type and a typeDefinition reflect.Type. keyType must implement the comparable interface.
// NewMapOfTypeWithReflectTypeKey creates a pointer to a map of the type specified. This function can create instances of maps of
// dynamic types. This function can be used with any reflect.Type.
func NewMapOfTypeWithReflectTypeKey(keyType reflect.Type, typeDefinition reflect.Type) (any, error) {
	if !keyType.Comparable() {
		return nil, errors.New("keyType is not comparable")
	}
	si := NewTypeInstance(typeDefinition)
	sd := reflect.ValueOf(&si).Type().Elem()
	mTyp := reflect.MapOf(reflect.TypeOf(keyType), sd)
	newMap := reflect.MakeMap(mTyp)
	return newMap.Interface(), nil
}

// SetStructFieldValue takes a typeInstance (generic/dynamic struct or reflect created instance), a fieldName string to specify the field that to set
// the value of, and the value to set the field to. The typeInstance passed to this function must be a pointer to a type instance. Instances created
// by the NewTypeInstance function are pointers. If the type instance is not a pointer to a struct instance, the fieldName does not exist on this
// typeInstance, or the underlying type of the field does not match the type of the fieldValue this function returns an error.
func SetStructFieldValue[T any](typeInstance any, fieldName string, fieldValue T) error {
	if reflect.ValueOf(typeInstance).Kind() != reflect.Ptr ||
		reflect.ValueOf(typeInstance).Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct instance, got %s", reflect.ValueOf(typeInstance).Elem().Kind())
	}
	if !reflect.ValueOf(typeInstance).Elem().FieldByName(fieldName).IsValid() {
		return fmt.Errorf("field with name: \"%s\" does not exist on struct instance", fieldName)
	}
	if reflect.TypeFor[T]().Kind() != reflect.ValueOf(typeInstance).Elem().FieldByName(fieldName).Kind() {
		if reflect.TypeFor[T]().Kind() == reflect.Interface {
			reflect.ValueOf(typeInstance).Elem().FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
			return nil
		}
		return fmt.Errorf("field with name: \"%s\" has underlying type: %s, but fieldValue argument has type: %s", fieldName, reflect.ValueOf(typeInstance).Elem().FieldByName(fieldName).Kind().String(), reflect.TypeFor[T]().Kind().String())
	}
	reflect.ValueOf(typeInstance).Elem().FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
	return nil
}

// GetStructFieldValue is a generic function. It accepts a typeInstance, a fieldName, and the expected return type T. Returns the value of
// the field specified in fieldName. If typeInstance is not a struct, the fieldName doesn't exist on the struct, or the underlying type of
// the field can not be type asserted to the type T, this function returns an error.
func GetStructFieldValue[T any](typeInstance any, fieldName string) (T, error) {
	val := reflect.ValueOf(typeInstance)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), errors.New("expected a struct instance")
	}
	if !val.FieldByName(fieldName).IsValid() {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), errors.New("field not valid for struct instance")
	}
	if val.FieldByName(fieldName).Kind() != reflect.TypeFor[T]().Kind() {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), fmt.Errorf("field with name: \"%s\" has underlying type: %s, but generic type assertion was for type: %s", fieldName, val.FieldByName(fieldName).Kind().String(), reflect.TypeFor[T]().Kind().String())
	}
	return val.FieldByName(fieldName).Interface().(T), nil
}

// GetStructFieldValueAny is like GetStructFieldValue, however, it doesn't take a generic T. This function returns The value on the typeInstance of the fieldName specified.
// If fieldName is not present on the typeInstance or the typeInstance is not a struct this function returns an error. To use the value returned from this function,
// it should be type asserted.
func GetStructFieldValueAny(typeInstance any, fieldName string) (any, error) {
	val := reflect.ValueOf(typeInstance)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, errors.New("expected a struct instance")
	}
	if !val.FieldByName(fieldName).IsValid() {
		return nil, errors.New("field not valid for struct instance")
	}
	return val.FieldByName(fieldName).Interface(), nil
}
