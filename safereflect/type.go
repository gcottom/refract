package safereflect

import (
	"errors"
	"fmt"
	"reflect"
)

type ChanDir int

const (
	RecvDir ChanDir             = 1 << iota // <-chan
	SendDir                                 // chan<-
	BothDir = RecvDir | SendDir             // chan
)

type Type interface {
	Align() int
	FieldAlign() int
	Method(int) (Method, error)
	MethodByName(string) (Method, bool)
	NumMethod() int
	Name() string
	PkgPath() string
	Size() uintptr
	String() string
	Kind() Kind
	Implements(u Type) (bool, error)
	AssignableTo(u Type) (bool, error)
	ConvertibleTo(u Type) (bool, error)
	Comparable() bool
	Bits() int
	ChanDir() (ChanDir, error)
	IsVariadic() (bool, error)
	Elem() Type
	Field(int) (StructField, error)
	FieldByIndex(index []int) (StructField, error)
	FieldByName(name string) (StructField, bool)
	FieldByNameFunc(match func(string) bool) (StructField, bool)
	In(i int) (Type, error)
	Key() (Type, error)
	Len() (int, error)
	NumField() (int, error)
	NumIn() (int, error)
	NumOut() (int, error)
	Out(i int) (Type, error)
	ReflectType() reflect.Type
}

type RefractType struct {
	T reflect.Type
}

func TypeOf(i any) Type {
	return &RefractType{reflect.TypeOf(i)}
}

func (t *RefractType) ReflectType() reflect.Type {
	return t.T
}

func (t *RefractType) Name() string {
	return t.T.Name()
}

func (t *RefractType) PkgPath() string {
	return t.T.PkgPath()
}

func (t *RefractType) Size() uintptr {
	return t.T.Size()
}

func (t *RefractType) String() string {
	return t.T.String()
}

func (t *RefractType) Implements(u Type) (bool, error) {
	if u == nil {
		return false, errors.New("reflect: nil type passed to Type.Implements")
	}
	if u.Kind() != Interface {
		return false, fmt.Errorf("reflect: non-interface type passed to Type.Implements")
	}
	return t.T.Implements(u.ReflectType()), nil
}

func (t *RefractType) Align() int {
	return int(t.T.Align())
}

func (t *RefractType) FieldAlign() int {
	return int(t.T.FieldAlign())
}

func (t *RefractType) Kind() Kind {
	return Kind(t.T.Kind())
}

func (t *RefractType) NumMethod() int {
	return t.T.NumMethod()
}

func (t *RefractType) Method(i int) (Method, error) {
	if t.Kind() == Interface {
		m := t.T.Method(i)
		return Method{
			Name:    m.Name,
			PkgPath: m.PkgPath,
			Type:    TypeOf(m.Type),
			Func:    ValueOf(m.Func),
			Index:   m.Index,
		}, nil
	}
	if i < 0 || i >= t.T.NumMethod() {
		return Method{}, errors.New("reflect: Method index out of range")
	}
	m := t.T.Method(i)
	return Method{
		Name:    m.Name,
		PkgPath: m.PkgPath,
		Type:    TypeOf(m.Type),
		Func:    ValueOf(m.Func),
		Index:   m.Index,
	}, nil
}

func (t *RefractType) MethodByName(name string) (Method, bool) {
	m, ok := t.T.MethodByName(name)
	if !ok {
		return Method{}, false
	}
	return Method{
		Name:    m.Name,
		PkgPath: m.PkgPath,
		Type:    TypeOf(m.Type),
		Func:    ValueOf(m.Func),
		Index:   m.Index,
	}, true
}

func (t *RefractType) ChanDir() (ChanDir, error) {
	if t.Kind() != Chan {
		return 0, fmt.Errorf("reflect: ChanDir of non-chan type %s", t.String())
	}
	return ChanDir(t.T.ChanDir()), nil
}

func (t *RefractType) AssignableTo(u Type) (bool, error) {
	if u == nil {
		return false, errors.New("reflect: nil type passed to Type.AssignableTo")
	}
	return t.T.AssignableTo(u.ReflectType()), nil
}

func (t *RefractType) ConvertibleTo(u Type) (bool, error) {
	if u == nil {
		return false, errors.New("reflect: nil type passed to Type.ConvertibleTo")
	}
	return t.T.ConvertibleTo(u.ReflectType()), nil
}

func (t *RefractType) Comparable() bool {
	return t.T.Comparable()
}

func (t *RefractType) Bits() int {
	return int(t.T.Bits())
}

func (t *RefractType) IsVariadic() (bool, error) {
	if t.Kind() != Func {
		return false, fmt.Errorf("reflect: IsVariadic of non-func type %s", t.String())
	}
	return t.T.IsVariadic(), nil
}

func (t *RefractType) Elem() Type {
	return TypeOf(t.T.Elem())
}

func (t *RefractType) Field(i int) (StructField, error) {
	if t.Kind() != Struct {
		return StructField{}, fmt.Errorf("reflect: Field of non-struct type %s", t.String())
	}
	if i < 0 || i >= t.T.NumField() {
		return StructField{}, errors.New("reflect: Field index out of range")
	}
	f := t.T.Field(i)
	return StructField{
		Name:      f.Name,
		PkgPath:   f.PkgPath,
		Type:      TypeOf(f.Type),
		Tag:       StructTag(f.Tag),
		Offset:    f.Offset,
		Index:     f.Index,
		Anonymous: f.Anonymous,
	}, nil
}

func (t *RefractType) FieldByIndex(index []int) (StructField, error) {
	if t.Kind() != Struct {
		return StructField{}, fmt.Errorf("reflect: FieldByIndex of non-struct type %s", t.String())
	}
	if len(index) == 0 {
		return StructField{}, errors.New("reflect: FieldByIndex index is empty")
	}
	tt := t.T.FieldByIndex(index)
	return StructField{
		Name:      tt.Name,
		PkgPath:   tt.PkgPath,
		Type:      TypeOf(tt.Type),
		Tag:       StructTag(tt.Tag),
		Offset:    tt.Offset,
		Index:     tt.Index,
		Anonymous: tt.Anonymous,
	}, nil
}

func (t *RefractType) FieldByName(name string) (StructField, bool) {
	if t.Kind() != Struct {
		return StructField{}, false
	}
	f, ok := t.T.FieldByName(name)
	if !ok {
		return StructField{}, false
	}
	return StructField{
		Name:      f.Name,
		PkgPath:   f.PkgPath,
		Type:      TypeOf(f.Type),
		Tag:       StructTag(f.Tag),
		Offset:    f.Offset,
		Index:     f.Index,
		Anonymous: f.Anonymous,
	}, true
}

func (t *RefractType) FieldByNameFunc(match func(string) bool) (StructField, bool) {
	if t.Kind() != Struct {
		return StructField{}, false
	}
	f, ok := t.T.FieldByNameFunc(match)
	if !ok {
		return StructField{}, false
	}
	return StructField{
		Name:      f.Name,
		PkgPath:   f.PkgPath,
		Type:      TypeOf(f.Type),
		Tag:       StructTag(f.Tag),
		Offset:    f.Offset,
		Index:     f.Index,
		Anonymous: f.Anonymous,
	}, true
}

func (t *RefractType) Key() (Type, error) {
	if t.Kind() != Map {
		return nil, fmt.Errorf("reflect: Key of non-map type %s", t.String())
	}
	return TypeOf(t.T.Key()), nil
}

func (t *RefractType) Len() (int, error) {
	if t.Kind() != Array {
		return 0, fmt.Errorf("reflect: Len of non-array type %s", t.String())
	}
	return t.T.Len(), nil
}

func (t *RefractType) NumField() (int, error) {
	if t.Kind() != Struct {
		return 0, fmt.Errorf("reflect: NumField of non-struct type %s", t.String())
	}
	return t.T.NumField(), nil
}

func (t *RefractType) NumIn() (int, error) {
	if t.Kind() != Func {
		return 0, fmt.Errorf("reflect: NumIn of non-func type %s", t.String())
	}
	return t.T.NumIn(), nil
}

func (t *RefractType) NumOut() (int, error) {
	if t.Kind() != Func {
		return 0, fmt.Errorf("reflect: NumOut of non-func type %s", t.String())
	}
	return t.T.NumOut(), nil
}

func (t *RefractType) In(i int) (Type, error) {
	if t.Kind() != Func {
		return nil, fmt.Errorf("reflect: In of non-func type %s", t.String())
	}
	return TypeOf(t.T.In(i)), nil
}

func (t *RefractType) Out(i int) (Type, error) {
	if t.Kind() != Func {
		return nil, fmt.Errorf("reflect: Out of non-func type %s", t.String())
	}
	return TypeOf(t.T.Out(i)), nil
}

func TypeFor[T any]() Type {
	return TypeOf((*T)(nil)).Elem()
}

func MapOf(key, elem Type) (Type, error) {
	if !ValueOf(key).Comparable() {
		return nil, errors.New("reflect.MapOf with incomparable key type")
	}
	return &RefractType{reflect.MapOf(key.ReflectType(), elem.ReflectType())}, nil
}

func SliceOf(t Type) Type {
	return &RefractType{reflect.SliceOf(t.ReflectType())}
}

// StructOf returns a new struct type with the given fields. StructOf can panic due to the complexity of requirements for structs (unexported fields, etc).
func StructOf(fields []StructField) (Type, error) {
	var f []reflect.StructField
	for _, field := range fields {
		if field.Name == "" {
			return nil, errors.New("reflect.StructOf: field " + field.Type.String() + " has no name")
		}
		if field.Type == nil {
			return nil, errors.New("reflect.StructOf: field " + field.Name + " has no type")
		}
		f = append(f, reflect.StructField{
			Name:      field.Name,
			PkgPath:   field.PkgPath,
			Type:      field.Type.ReflectType(),
			Tag:       reflect.StructTag(field.Tag),
			Offset:    field.Offset,
			Index:     field.Index,
			Anonymous: field.Anonymous,
		})
	}
	return &RefractType{reflect.StructOf(f)}, nil
}
