package gendynamic

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/gcottom/refract/safereflect"
)

// NewStructField is used to create new struct fields to be used with NewStructDefinition as arguments. This function
// takes a fieldName string, fieldType any, and fieldTag string as arguments. The field must be exported, therefore the
// first letter of the fieldName will automatically be capitalized if it is not already capitalized. fieldType should
// be a var with the type that the field should hold. For string type, it would be "string", for int, it would be an
// int like 1 or 2. fieldTag takes a string which contains the tags that would normally be in a struct tag. For example `json:"fieldName"`.
func NewStructField(fieldName string, fieldType any, fieldTag string) safereflect.StructField {
	fieldName = strings.ReplaceAll(fieldName, " ", "")
	if !unicode.IsUpper(rune(fieldName[0])) {
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(rune(fieldName[0]))
		fieldName = string(r)
	}

	return safereflect.StructField{Name: fieldName, Type: safereflect.TypeOf(fieldType), Tag: safereflect.StructTag(fieldTag)}
}

// NewStructFieldWithReflectTag is used to create new struct fields to be used with NewStructDefinition as arguments. This function
// takes a fieldName string, fieldType reflect.Type, and fieldTag reflect.StructTag as arguments. The field must be exported, therefore the
// first letter of the fieldName will automatically be capitalized if it is not already capitalized. fieldType should
// be a var with the type that the field should hold. For string type, it would be "string", for int, it would be an
// int like 1 or 2. fieldTag takes a reflect.StructTag.
func NewStructFieldWithReflectTag(fieldName string, fieldType any, fieldTag safereflect.StructTag) safereflect.StructField {
	fieldName = strings.ReplaceAll(fieldName, " ", "")
	if !unicode.IsUpper(rune(fieldName[0])) {
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(rune(fieldName[0]))
		fieldName = string(r)
	}

	return safereflect.StructField{Name: fieldName, Type: safereflect.TypeOf(fieldType), Tag: fieldTag}
}

// NewStructFieldWithReflectTagAndType is used to create new struct fields to be used with NewStructDefinition as arguments. This function
// takes a fieldName string, fieldType reflect.Type, and fieldTag reflect.StructTag as arguments. The field must be exported, therefore the
// first letter of the fieldName will automatically be capitalized if it is not already capitalized. fieldType should
// be a reflect.Type, it can also be a struct definition created by refract. fieldTag takes a reflect.StructTag.
func NewStructFieldWithReflectTagAndType(fieldName string, fieldType safereflect.Type, fieldTag safereflect.StructTag) safereflect.StructField {
	fieldName = strings.ReplaceAll(fieldName, " ", "")
	if !unicode.IsUpper(rune(fieldName[0])) {
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(rune(fieldName[0]))
		fieldName = string(r)
	}

	return safereflect.StructField{Name: fieldName, Type: fieldType, Tag: fieldTag}
}

// NewStructFieldWithReflectType is used to create new struct fields to be used with NewStructDefinition as arguments. This function
// takes a fieldName string, fieldType reflect.Type, and fieldTag string as arguments. The field must be exported, therefore the
// first letter of the fieldName will automatically be capitalized if it is not already capitalized. fieldType should
// be a reflect.Type, it can also be a struct definition created by refract. fieldTag takes a string which contains the tags that
// would normally be in a struct tag. For example `json:"fieldName"`.
func NewStructFieldWithReflectType(fieldName string, fieldType safereflect.Type, fieldTag string) safereflect.StructField {
	fieldName = strings.ReplaceAll(fieldName, " ", "")
	if !unicode.IsUpper(rune(fieldName[0])) {
		r := []rune(fieldName)
		r[0] = unicode.ToUpper(rune(fieldName[0]))
		fieldName = string(r)
	}

	return safereflect.StructField{Name: fieldName, Type: fieldType, Tag: safereflect.StructTag(fieldTag)}
}

// NewStructDefinition takes a variadic of fields which are reflect.StructField. reflect.StructField can be created by
// calling the NewStructField function. NewStructDefinition creates a reflect.Type which can be used in the NewTypeInstance,
// NewSliceOfType, and NewMapOfType functions.
func NewStructDefinition(fields ...safereflect.StructField) (safereflect.Type, error) {
	return safereflect.StructOf(fields)
}

// NewTypeInstance takes a typeDefinition as an argument. typeDefinition is a reflect.Type which can be created by using the
// NewStructDefinition function. NewTypeInstance creates a pointer to an instance of the type specified. By using this function,
// you can create dynamic and generic instances of a type. You can also use this function with any reflect.Type to create a new
// instance of a native type.
func NewTypeInstance(typeDefinition safereflect.Type) (any, error) {
	if typeDefinition.Kind() == safereflect.Pointer {
		typeDefinition = typeDefinition.Elem()
	}
	val, err := safereflect.New(typeDefinition)
	if err != nil {
		return nil, err
	}
	return val.Interface()
}

// NewSliceOfType is used with a typeDefinition. typeDefinition is a reflect.Type which can be created by using the NewStructDefinition
// function. NewSliceOfType creates a pointer to a slice of the type specified. This function can be used to create slices of dynamic types.
// This function can be used with any reflect.Type.
func NewSliceOfType(typeDefinition safereflect.Type) (any, error) {
	si, err := NewTypeInstance(typeDefinition)
	if err != nil {
		return nil, err
	}
	sd := safereflect.ValueOf(&si).Type().Elem()
	sTyp := safereflect.SliceOf(sd)
	newSlice, err := safereflect.MakeSlice(sTyp, 0, 0)
	if err != nil {
		return nil, err
	}
	return newSlice.Interface()
}

// NewMapOfType is used with a keyType (comparable) and a typeDefinition. keyType must implement the comparable interface. typeDefinition is a reflect.Type
// which can be created by using the NewStructDefinition function. NewMapOfType creates a pointer to a map of the type specified.
// This function can create instances of maps of dynamic types. This function can be used with any reflect.Type.
func NewMapOfType[T comparable](keyType T, typeDefinition safereflect.Type) (any, error) {
	si, err := NewTypeInstance(typeDefinition)
	if err != nil {
		return nil, err
	}
	sd := safereflect.ValueOf(si).Type().Elem()
	fmt.Println(sd)
	mTyp, err := safereflect.MapOf(safereflect.TypeOf(keyType), sd)
	if err != nil {
		return nil, err
	}
	fmt.Println(mTyp)
	newMap, err := safereflect.MakeMap(mTyp)
	if err != nil {
		return nil, err
	}
	fmt.Println(newMap)
	return newMap.Interface()
}

// NewMapOfTypeWithReflectTypeKey is used with a keyType reflect.Type and a typeDefinition reflect.Type. keyType must implement the comparable interface.
// NewMapOfTypeWithReflectTypeKey creates a pointer to a map of the type specified. This function can create instances of maps of
// dynamic types. This function can be used with any reflect.Type.
func NewMapOfTypeWithReflectTypeKey(keyType safereflect.Type, typeDefinition safereflect.Type) (any, error) {
	if !keyType.Comparable() {
		return nil, errors.New("keyType is not comparable")
	}
	si, err := NewTypeInstance(typeDefinition)
	if err != nil {
		return nil, err
	}
	sd := safereflect.ValueOf(&si).Type().Elem()
	mTyp, err := safereflect.MapOf(safereflect.TypeOf(keyType), sd)
	if err != nil {
		return nil, err
	}
	newMap, err := safereflect.MakeMap(mTyp)
	if err != nil {
		return nil, err
	}
	return newMap.Interface()
}

// SetStructFieldValue takes a typeInstance (generic/dynamic struct or reflect created instance), a fieldName string to specify the field that to set
// the value of, and the value to set the field to. The typeInstance passed to this function must be a pointer to a type instance. Instances created
// by the NewTypeInstance function are pointers. If the type instance is not a pointer to a struct instance, the fieldName does not exist on this
// typeInstance, or the underlying type of the field does not match the type of the fieldValue this function returns an error.
func SetStructFieldValue[T any](typeInstance any, fieldName string, fieldValue T) error {
	e, err := safereflect.ValueOf(typeInstance).Elem()
	if err != nil {
		return err
	}
	if safereflect.ValueOf(typeInstance).Kind() != safereflect.Pointer || e.Kind() != safereflect.Struct {
		return fmt.Errorf("expected a pointer to a struct instance, got %s", e.Kind())
	}
	efn, err := e.FieldByName(fieldName)
	if err != nil {
		return fmt.Errorf("field with name: \"%s\" does not exist on struct instance", fieldName)
	}
	if !efn.IsValid() {
		return fmt.Errorf("field with name: \"%s\" does not exist on struct instance", fieldName)
	}
	if safereflect.ValueOf(fieldValue).Type().Kind() != efn.Kind() {
		if safereflect.TypeFor[T]().Kind() == safereflect.Interface {
			return efn.Set(safereflect.ValueOf(fieldValue))
		}
		return fmt.Errorf("field with name: \"%s\" has underlying type: %s, but fieldValue argument has type: %s", fieldName, efn.Kind().String(), safereflect.TypeFor[T]().Kind().String())
	}
	return efn.Set(safereflect.ValueOf(fieldValue))
}

// GetStructFieldValue is a generic function. It accepts a typeInstance, a fieldName, and the expected return type T. Returns the value of
// the field specified in fieldName. If typeInstance is not a struct, the fieldName doesn't exist on the struct, or the underlying type of
// the field can not be type asserted to the type T, this function returns an error.
func GetStructFieldValue[T any](typeInstance any, fieldName string) (T, error) {
	var err error
	val := safereflect.ValueOf(typeInstance)
	if val.Kind() == safereflect.Pointer {
		val, err = val.Elem()
		if err != nil {
			return safereflect.ZeroGeneric[T](), err
		}
	}
	if val.Kind() != safereflect.Struct {
		return safereflect.ZeroGeneric[T](), errors.New("expected a struct instance")
	}
	fn, err := val.FieldByName(fieldName)
	if err != nil {
		return safereflect.ZeroGeneric[T](), errors.New("field not valid for struct instance")
	}
	if !fn.IsValid() {
		return safereflect.ZeroGeneric[T](), errors.New("field not valid for struct instance")
	}
	if fn.Kind() != safereflect.TypeFor[T]().Kind() {
		return safereflect.ZeroGeneric[T](), fmt.Errorf("field with name: \"%s\" has underlying type: %s, but generic type assertion was for type: %s", fieldName, fn.Kind().String(), safereflect.TypeFor[T]().Kind().String())
	}
	fni, err := fn.Interface()
	if err != nil {
		return safereflect.ZeroGeneric[T](), err
	}
	return fni.(T), nil
}

// GetStructFieldValueAny is like GetStructFieldValue, however, it doesn't take a generic T. This function returns The value on the typeInstance of the fieldName specified.
// If fieldName is not present on the typeInstance or the typeInstance is not a struct this function returns an error. To use the value returned from this function,
// it should be type asserted.
func GetStructFieldValueAny(typeInstance any, fieldName string) (any, error) {
	var err error
	val := safereflect.ValueOf(typeInstance)
	if val.Kind() == safereflect.Pointer {
		val, err = val.Elem()
		if err != nil {
			return nil, err
		}
	}
	if val.Kind() != safereflect.Struct {
		return nil, errors.New("expected a struct instance")
	}
	fn, err := val.FieldByName(fieldName)
	if err != nil {
		return nil, errors.New("field not valid for struct instance")
	}
	if !fn.IsValid() {
		return nil, errors.New("field not valid for struct instance")
	}
	return fn.Interface()
}
