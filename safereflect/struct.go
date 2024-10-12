package safereflect

import "reflect"

type StructField struct {
	// Name is the field name.
	Name string

	// PkgPath is the package path that qualifies a lower case (unexported)
	// field name. It is empty for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath string

	Type      Type      // field type
	Tag       StructTag // field tag string
	Offset    uintptr   // offset within struct, in bytes
	Index     []int     // index sequence for Type.FieldByIndex
	Anonymous bool      // is an embedded field
}

// IsExported reports whether the field is exported.
func (f StructField) IsExported() bool {
	return f.PkgPath == ""
}

type StructTag string

func (tag StructTag) Get(key string) (value string) {
	return reflect.StructTag(tag).Get(key)
}

func (tag StructTag) Lookup(key string) (value string, ok bool) {
	return reflect.StructTag(tag).Lookup(key)
}
