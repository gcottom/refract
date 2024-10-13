package safereflect

import (
	"errors"
	"reflect"
	"unsafe"
)

type Value struct {
	v reflect.Value
}

// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i. ValueOf(nil) returns the zero Value.
func ValueOf(i any) Value {
	if i == nil {
		return Value{}
	}
	return Value{reflect.ValueOf(i)}
}

func (v Value) Addr() (Value, error) {
	if v.v.CanAddr() {
		return Value{v.v.Addr()}, nil
	}
	return Value{}, errors.New("reflect.Value.Addr of unaddressable value")
}

func (v Value) Bool() (bool, error) {
	if v.v.Kind() == reflect.Bool {
		return v.v.Bool(), nil
	}
	return false, errors.New("reflect.Value.Bool of non-bool value")
}

func (v Value) Bytes() ([]byte, error) {
	if v.v.Kind() == reflect.Slice && v.v.Type().Elem().Kind() == reflect.Uint8 {
		return v.v.Bytes(), nil
	}
	return nil, errors.New("reflect.Value.Bytes of non-[]byte value")
}

func (v Value) CanAddr() bool {
	return v.v.CanAddr()
}

func (v Value) CanSet() bool {
	return v.v.CanSet()
}

func (v Value) Call(args []Value) ([]Value, error) {
	if v.v.Kind() != reflect.Func {
		return nil, errors.New("reflect.Value.Call of non-function")
	}
	var in []reflect.Value
	for _, arg := range args {
		in = append(in, arg.v)
	}
	out := v.v.Call(in)
	var outv []Value
	for _, v := range out {
		outv = append(outv, Value{v})
	}
	return outv, nil
}

func (v Value) CallSlice(args []Value) ([]Value, error) {
	if v.v.Kind() != reflect.Func {
		return nil, errors.New("reflect.Value.Call of non-function")
	}
	var in []reflect.Value
	for _, arg := range args {
		in = append(in, arg.v)
	}
	out := v.v.Call(in)
	var outv []Value
	for _, v := range out {
		outv = append(outv, Value{v})
	}
	return outv, nil
}

func (v Value) Close() {
	if v.v.Kind() != reflect.Chan {
		return
	}
	if v.v.Type().ChanDir()&reflect.ChanDir(SendDir) == 0 {
		return
	}
	v.v.Close()
}

func (v Value) CanComplex() bool {
	return v.v.CanComplex()
}

func (v Value) Complex() (complex128, error) {
	if v.v.Kind() == reflect.Complex64 || v.v.Kind() == reflect.Complex128 {
		return v.v.Complex(), nil
	}
	return 0, errors.New("reflect.Value.Complex of non-complex value")
}

// Can still panic if this is a pointer to a not-in-heap object
func (v Value) Elem() (Value, error) {
	if v.v.Kind() == reflect.Ptr || v.v.Kind() == reflect.Interface {
		return Value{v.v.Elem()}, nil
	}
	return Value{}, errors.New("reflect.Value.Elem of non-pointer")

}

func (v Value) Field(i int) (Value, error) {
	if v.v.Kind() != reflect.Struct {
		return Value{}, errors.New("reflect.Value.Field of non-struct value")
	}
	if i < 0 || i >= v.v.NumField() {
		return Value{}, errors.New("reflect.Value.Field index out of range")
	}
	return Value{v.v.Field(i)}, nil
}

func (v Value) FieldByIndex(index []int) (Value, error) {
	if v.v.Kind() != reflect.Struct {
		return Value{}, errors.New("reflect.Value.FieldByIndex of non-struct value")
	}
	if len(index) == 0 {
		return Value{}, errors.New("reflect.Value.FieldByIndex with no index")
	}
	return Value{v.v.FieldByIndex(index)}, nil
}

// Can still panic if indirection through a nil pointer to embedded struct field occurs
func (v Value) FieldByIndexErr(index []int) (Value, error) {
	if v.v.Kind() != reflect.Struct {
		return Value{}, errors.New("reflect.Value.FieldByIndex of non-struct value")
	}
	if len(index) == 0 {
		return Value{}, errors.New("reflect.Value.FieldByIndex with no index")
	}
	return Value{v.v.FieldByIndex(index)}, nil
}

func (v Value) FieldByName(name string) (Value, error) {
	if v.v.Kind() != reflect.Struct {
		return Value{}, errors.New("reflect.Value.FieldByName of non-struct value")
	}
	return Value{v.v.FieldByName(name)}, nil
}

func (v Value) FieldByNameFunc(match func(string) bool) (Value, error) {
	if v.v.Kind() != reflect.Struct {
		return Value{}, errors.New("reflect.Value.FieldByNameFunc of non-struct value")
	}
	return Value{v.v.FieldByNameFunc(match)}, nil
}

func (v Value) CanFloat() bool {
	return v.v.CanFloat()
}

func (v Value) Float() (float64, error) {
	if v.v.Kind() == reflect.Float32 || v.v.Kind() == reflect.Float64 {
		return v.v.Float(), nil
	}
	return 0, errors.New("reflect.Value.Float of non-float value")
}

func (v Value) Index(i int) (Value, error) {
	if v.v.Kind() != reflect.Array && v.v.Kind() != reflect.Slice && v.v.Kind() != reflect.String {
		return Value{}, errors.New("reflect.Value.Index of non-array, non-slice, non-string value")
	}
	if i < 0 || i >= v.v.Len() {
		return Value{}, errors.New("reflect.Value.Index index out of range")
	}
	return Value{v.v.Index(i)}, nil
}

func (v Value) CanInt() bool {
	return v.v.CanInt()
}

func (v Value) Int() (int64, error) {
	if v.v.Kind() == reflect.Int || v.v.Kind() == reflect.Int8 || v.v.Kind() == reflect.Int16 || v.v.Kind() == reflect.Int32 || v.v.Kind() == reflect.Int64 {
		return v.v.Int(), nil
	}
	return 0, errors.New("reflect.Value.Int of non-int value")
}

func (v Value) CanInterface() bool {
	return v.v.CanInterface()
}

func (v Value) Interface() (any, error) {
	if !v.v.CanInterface() {
		return nil, errors.New("reflect.Value.Interface of unexported value")
	}
	return v.v.Interface(), nil
}

func (v Value) IsNil() (bool, error) {
	if v.v.Kind() != reflect.Chan && v.v.Kind() != reflect.Func && v.v.Kind() != reflect.Interface && v.v.Kind() != reflect.Map && v.v.Kind() != reflect.Ptr && v.v.Kind() != reflect.Slice {
		return false, errors.New("reflect.Value.IsNil of non-pointer")
	}
	return v.v.IsNil(), nil
}

func (v Value) IsValid() bool {
	return v.v.IsValid()
}

func (v Value) IsZero() (bool, error) {
	switch v.v.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Array, reflect.String, reflect.Struct:
		return v.v.IsZero(), nil
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	default:
		return false, errors.New("reflect.Value.IsZero of invalid")
	}
}

func (v Value) SetZero() error {
	if !v.v.CanSet() {
		return errors.New("reflect.Value.SetZero of unaddressable value")
	}
	v.v.SetZero()
	return nil
}

func (v Value) Kind() Kind {
	return Kind(v.v.Kind())
}

func (v Value) Len() (int, error) {
	if v.v.Kind() == reflect.Array || v.v.Kind() == reflect.Slice || v.v.Kind() == reflect.String {
		return v.v.Len(), nil
	}
	if v.v.Kind() == reflect.Pointer {
		t, err := v.Elem()
		if err != nil {
			return 0, err
		}
		if t.Kind() == Array {
			return v.v.Len(), nil
		}
	}
	return 0, errors.New("reflect.Value.Len of non-array, non-slice, non-string value")
}

func (v Value) MapIndex(key Value) (Value, error) {
	if v.v.Kind() != reflect.Map {
		return Value{}, errors.New("reflect.Value.MapIndex of non-map value")
	}
	return Value{v.v.MapIndex(key.v)}, nil
}

func (v Value) MapKeys() ([]Value, error) {
	if v.v.Kind() != reflect.Map {
		return nil, errors.New("reflect.Value.MapKeys of non-map value")
	}
	var keys []Value
	for _, key := range v.v.MapKeys() {
		keys = append(keys, Value{key})
	}
	return keys, nil
}

func (v Value) MapRange() (*reflect.MapIter, error) {
	if v.v.Kind() != reflect.Map {
		return nil, errors.New("reflect.Value.MapRange of non-map value")
	}
	return v.v.MapRange(), nil
}

func (v Value) Method(i int) (Value, error) {
	if v.v.Kind() != reflect.Struct {
		return Value{}, errors.New("reflect.Value.Method of non-struct value")
	}
	if i < 0 || i >= v.v.NumMethod() {
		return Value{}, errors.New("reflect.Value.Method index out of range")
	}
	return Value{v.v.Method(i)}, nil
}

func (v Value) NumMethod() (int, error) {
	if v.v.Kind() != reflect.Struct {
		return 0, errors.New("reflect.Value.NumMethod of non-struct value")
	}
	return v.v.NumMethod(), nil
}

func (v Value) MethodByName(name string) (Value, error) {
	if v.v.Kind() != reflect.Struct {
		return Value{}, errors.New("reflect.Value.MethodByName of non-struct value")
	}
	return Value{v.v.MethodByName(name)}, nil
}

func (v Value) OverflowComplex(x complex128) (bool, error) {
	if v.v.Kind() != reflect.Complex64 && v.v.Kind() != reflect.Complex128 {
		return false, errors.New("reflect.Value.OverflowComplex of non-complex value")
	}
	return v.v.OverflowComplex(x), nil
}

func (v Value) OverflowFloat(x float64) (bool, error) {
	if v.v.Kind() != reflect.Float32 && v.v.Kind() != reflect.Float64 {
		return false, errors.New("reflect.Value.OverflowFloat of non-float value")
	}
	return v.v.OverflowFloat(x), nil
}

func (v Value) OverflowInt(x int64) (bool, error) {
	if v.v.Kind() != reflect.Int && v.v.Kind() != reflect.Int8 && v.v.Kind() != reflect.Int16 && v.v.Kind() != reflect.Int32 && v.v.Kind() != reflect.Int64 {
		return false, errors.New("reflect.Value.OverflowInt of non-int value")
	}
	return v.v.OverflowInt(x), nil
}

func (v Value) OverflowUint(x uint64) (bool, error) {
	if v.v.Kind() != reflect.Uint && v.v.Kind() != reflect.Uint8 && v.v.Kind() != reflect.Uint16 && v.v.Kind() != reflect.Uint32 && v.v.Kind() != reflect.Uint64 && v.v.Kind() != reflect.Uintptr {
		return false, errors.New("reflect.Value.OverflowUint of non-uint value")
	}
	return v.v.OverflowUint(x), nil
}

func (v Value) Pointer() (uintptr, error) {
	k := v.v.Kind()
	switch k {
	case reflect.Pointer, reflect.Chan, reflect.Map, reflect.UnsafePointer, reflect.Func, reflect.Slice:
		return v.v.Pointer(), nil
	}
	return 0, errors.New("reflect.Value.Pointer of non-pointer value")
}

func (v Value) Recv() (Value, bool, error) {
	if v.v.Kind() != reflect.Chan {
		return Value{}, false, errors.New("reflect.Value.Recv of non-chan value")
	}
	if v.v.Type().ChanDir()&reflect.RecvDir == 0 {
		return Value{}, false, errors.New("reflect.Value.Recv of non-receivable chan")
	}
	x, b := v.v.Recv()
	return Value{x}, b, nil
}

func (v Value) Send(x Value) error {
	if v.v.Kind() != reflect.Chan {
		return errors.New("reflect.Value.Send of non-chan value")
	}
	if v.v.Type().ChanDir()&reflect.SendDir == 0 {
		return errors.New("reflect.Value.Send of non-sendable chan")
	}
	if v.v.Type().Elem() != x.v.Type() {
		return errors.New("reflect.Value.Send with wrong type")
	}
	v.v.Send(x.v)
	return nil
}

func (v Value) Set(x Value) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.Set of unaddressable value")
	}
	if v.v.Type() != x.v.Type() {
		return errors.New("reflect.Value.Set with wrong type")
	}
	v.v.Set(x.v)
	return nil
}

func (v Value) SetBool(x bool) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetBool of unaddressable value")
	}
	if v.v.Kind() != reflect.Bool {
		return errors.New("reflect.Value.SetBool of non-bool value")
	}
	v.v.SetBool(x)
	return nil
}

func (v Value) SetBytes(x []byte) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetBytes of unaddressable value")
	}
	if v.v.Kind() != reflect.Slice || v.v.Type().Elem().Kind() != reflect.Uint8 {
		return errors.New("reflect.Value.SetBytes of non-[]byte value")
	}
	v.v.SetBytes(x)
	return nil
}

func (v Value) SetComplex(x complex128) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetComplex of unaddressable value")
	}
	if v.v.Kind() != reflect.Complex64 && v.v.Kind() != reflect.Complex128 {
		return errors.New("reflect.Value.SetComplex of non-complex value")
	}
	v.v.SetComplex(x)
	return nil
}

func (v Value) SetFloat(x float64) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetFloat of unaddressable value")
	}
	if v.v.Kind() != reflect.Float32 && v.v.Kind() != reflect.Float64 {
		return errors.New("reflect.Value.SetFloat of non-float value")
	}
	v.v.SetFloat(x)
	return nil
}

func (v Value) SetInt(x int64) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetInt of unaddressable value")
	}
	if v.v.Kind() != reflect.Int && v.v.Kind() != reflect.Int8 && v.v.Kind() != reflect.Int16 && v.v.Kind() != reflect.Int32 && v.v.Kind() != reflect.Int64 {
		return errors.New("reflect.Value.SetInt of non-int value")
	}
	v.v.SetInt(x)
	return nil
}

func (v Value) SetLen(n int) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetLen of unaddressable value")
	}
	if v.v.Kind() != reflect.Slice {
		return errors.New("reflect.Value.SetLen of non-slice value")
	}
	v.v.SetLen(n)
	return nil
}

func (v Value) SetCap(n int) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetCap of unaddressable value")
	}
	if v.v.Kind() != reflect.Slice {
		return errors.New("reflect.Value.SetCap of non-slice value")
	}
	v.v.SetCap(n)
	return nil
}

func (v Value) SetMapIndex(key Value, elem Value) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetMapIndex of unaddressable value")
	}
	if v.v.Kind() != reflect.Map {
		return errors.New("reflect.Value.SetMapIndex of non-map value")
	}
	if v.v.Type().Key() != key.v.Type() {
		return errors.New("reflect.Value.SetMapIndex with wrong key type")
	}
	if v.v.Type().Elem() != elem.v.Type() {
		return errors.New("reflect.Value.SetMapIndex with wrong elem type")
	}
	v.v.SetMapIndex(key.v, elem.v)
	return nil
}

func (v Value) SetUint(x uint64) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetUint of unaddressable value")
	}
	if v.v.Kind() != reflect.Uint && v.v.Kind() != reflect.Uint8 && v.v.Kind() != reflect.Uint16 && v.v.Kind() != reflect.Uint32 && v.v.Kind() != reflect.Uint64 && v.v.Kind() != reflect.Uintptr {
		return errors.New("reflect.Value.SetUint of non-uint value")
	}
	v.v.SetUint(x)
	return nil
}

func (v Value) SetPointer(x unsafe.Pointer) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetPointer of unaddressable value")
	}
	if v.v.Kind() != reflect.Ptr {
		return errors.New("reflect.Value.SetPointer of non-pointer value")
	}
	v.v.SetPointer(x)
	return nil
}

func (v Value) SetString(x string) error {
	if !v.CanSet() {
		return errors.New("reflect.Value.SetString of unaddressable value")
	}
	if v.v.Kind() != reflect.String {
		return errors.New("reflect.Value.SetString of non-string value")
	}
	v.v.SetString(x)
	return nil
}

func (v Value) Slice(i int, j int) (Value, error) {
	if v.v.Kind() != reflect.Array && v.v.Kind() != reflect.Slice && v.v.Kind() != reflect.String {
		return Value{}, errors.New("reflect.Value.Slice of non-array, non-slice, non-string value")
	}
	if i < 0 || i >= v.v.Len() {
		return Value{}, errors.New("reflect.Value.Slice low out of range")
	}
	if j < i || j > v.v.Len() {
		return Value{}, errors.New("reflect.Value.Slice high out of range")
	}
	return Value{v.v.Slice(i, j)}, nil
}

func (v Value) Slice3(i int, j int, k int) (Value, error) {
	if v.v.Kind() != reflect.Array && v.v.Kind() != reflect.Slice && v.v.Kind() != reflect.String {
		return Value{}, errors.New("reflect.Value.Slice3 of non-array, non-slice, non-string value")
	}
	if i < 0 || i >= v.v.Len() {
		return Value{}, errors.New("reflect.Value.Slice3 low out of range")
	}
	if j < i || j > v.v.Len() {
		return Value{}, errors.New("reflect.Value.Slice3 mid out of range")
	}
	if k < j || k > v.v.Len() {
		return Value{}, errors.New("reflect.Value.Slice3 high out of range")
	}
	return Value{v.v.Slice3(i, j, k)}, nil
}

func (v Value) String() string {
	return v.v.String()
}

func (v Value) TryRecv() (Value, bool, error) {
	if v.v.Kind() != reflect.Chan {
		return Value{}, false, errors.New("reflect.Value.TryRecv of non-chan value")
	}
	if v.v.Type().ChanDir()&reflect.RecvDir == 0 {
		return Value{}, false, errors.New("reflect.Value.TryRecv of non-receivable chan")
	}
	x, b := v.v.TryRecv()
	return Value{x}, b, nil
}

func (v Value) TrySend(x Value) (bool, error) {
	if v.v.Kind() != reflect.Chan {
		return false, errors.New("reflect.Value.TrySend of non-chan value")
	}
	if v.v.Type().ChanDir()&reflect.SendDir == 0 {
		return false, errors.New("reflect.Value.TrySend of non-sendable chan")
	}
	if v.v.Type().Elem() != x.v.Type() {
		return false, errors.New("reflect.Value.TrySend with wrong type")
	}
	return v.v.TrySend(x.v), nil
}

func (v Value) Type() Type {
	t := v.v.Type()
	return &RefractType{t}
}

func (v Value) CanUint() bool {
	return v.v.CanUint()
}

func (v Value) Uint() (uint64, error) {
	if v.v.Kind() == reflect.Uint || v.v.Kind() == reflect.Uint8 || v.v.Kind() == reflect.Uint16 || v.v.Kind() == reflect.Uint32 || v.v.Kind() == reflect.Uint64 || v.v.Kind() == reflect.Uintptr {
		return v.v.Uint(), nil
	}
	return 0, errors.New("reflect.Value.Uint of non-uint value")
}

func (v Value) UnsafeAddr() (uintptr, error) {
	if v.v.CanAddr() {
		return v.v.UnsafeAddr(), nil
	}
	return 0, errors.New("reflect.Value.UnsafeAddr of unaddressable value")
}

// this can panic if the value is a not in heap pointer
func (v Value) UnsafePointer() (unsafe.Pointer, error) {
	switch v.v.Kind() {
	case reflect.Ptr, reflect.UnsafePointer, reflect.Chan, reflect.Map, reflect.Func, reflect.Slice:
		return v.v.UnsafePointer(), nil
	default:
		return nil, errors.New("reflect.Value.UnsafePointer of non-pointer value")
	}
}

func (v Value) Grow(n int) error {
	if v.v.Kind() != reflect.Slice {
		return errors.New("reflect.Value.Grow of non-slice value")
	}
	v.v.Grow(n)
	return nil
}

func (v Value) Clear() error {
	if v.v.Kind() != reflect.Slice && v.v.Kind() != reflect.Map {
		return errors.New("reflect.Value.Clear of non-slice, non-map value")
	}
	v.v.Clear()
	return nil
}

func Append(s Value, x ...Value) (Value, error) {
	if s.v.Kind() != reflect.Slice {
		return Value{}, errors.New("reflect.Append of non-slice value")
	}
	var in []reflect.Value
	for _, v := range x {
		if v.v.Type() != s.v.Type().Elem() {
			return Value{}, errors.New("reflect.Append with wrong type")
		}
		in = append(in, v.v)
	}
	out := reflect.Append(s.v, in...)
	return Value{out}, nil
}

func AppendSlice(s Value, t Value) (Value, error) {
	if s.v.Kind() != reflect.Slice || t.v.Kind() != reflect.Slice {
		return Value{}, errors.New("reflect.AppendSlice of non-slice value")
	}
	if t.v.Type().Elem() != s.v.Type().Elem() {
		return Value{}, errors.New("reflect.AppendSlice with wrong type")
	}
	out := reflect.AppendSlice(s.v, t.v)
	return Value{out}, nil
}

func Copy(dst Value, src Value) (int, error) {
	if dst.v.Kind() != reflect.Slice && dst.v.Kind() != reflect.Array {
		return 0, errors.New("reflect.Copy to non-slice, non-array value")
	}
	var stringCopy bool
	if src.v.Kind() != reflect.Slice && src.v.Kind() != reflect.Array {
		stringCopy = src.v.Kind() == reflect.String && dst.v.Type().Elem().Kind() == reflect.Uint8
		if !stringCopy {
			return 0, errors.New("reflect.Copy from non-slice, non-array value")
		}
	}
	return reflect.Copy(dst.v, src.v), nil
}

func MakeSlice(t Type, len int, cap int) (Value, error) {
	if t.Kind() != Slice {
		return Value{}, errors.New("reflect.MakeSlice with non-slice type")
	}
	if len < 0 {
		return Value{}, errors.New("reflect.MakeSlice with negative len")
	}
	if cap < 0 {
		return Value{}, errors.New("reflect.MakeSlice with negative cap")
	}
	if len > cap {
		return Value{}, errors.New("reflect.MakeSlice with len > cap")
	}
	out := reflect.MakeSlice(t.ReflectType(), len, cap)
	return Value{out}, nil
}

func MakeChan(t Type, buffer int) (Value, error) {
	if t.Kind() != Chan {
		return Value{}, errors.New("reflect.MakeChan with non-chan type")
	}
	if buffer < 0 {
		return Value{}, errors.New("reflect.MakeChan with negative buffer")
	}
	d, err := t.ChanDir()
	if err != nil {
		return Value{}, err
	}
	if d != BothDir {
		return Value{}, errors.New("reflect.MakeChan with unidirectional chan type")
	}
	out := reflect.MakeChan(t.ReflectType(), buffer)
	return Value{out}, nil
}

func MakeMap(t Type) (Value, error) {
	if t.Kind() != Map {
		return Value{}, errors.New("reflect.MakeMap with non-map type")
	}
	out := reflect.MakeMap(t.ReflectType())
	return Value{out}, nil
}

func Indirect(v Value) Value {
	return Value{reflect.Indirect(v.v)}
}

func Zero(t Type) (Value, error) {
	if t.ReflectType() == nil {
		return Value{}, errors.New("reflect.Zero with nil type")
	}
	out := reflect.Zero(t.ReflectType())
	return Value{out}, nil
}

func ZeroGeneric[T any]() T {
	return reflect.Zero(reflect.TypeFor[T]()).Interface().(T)
}

// may panic if t is a new type that may not be allocated in heap (possibly undefined cgo C type)
func New(t Type) (Value, error) {
	if t.ReflectType() == nil {
		return Value{}, errors.New("reflect.New with nil type")
	}
	out := reflect.New(t.ReflectType())
	return Value{out}, nil
}

// may panic if t is a new type that may not be allocated in heap (possibly undefined cgo C type)
func NewAt(t Type, p unsafe.Pointer) (Value, error) {
	if t.ReflectType() == nil {
		return Value{}, errors.New("reflect.NewAt with nil type")
	}
	out := reflect.NewAt(t.ReflectType(), p).Elem()
	return Value{out}, nil
}

func (v Value) CanConvert(t Type) bool {
	return v.v.CanConvert(t.ReflectType())
}

func (v Value) Convert(t Type) (Value, error) {
	if !v.v.CanConvert(t.ReflectType()) {
		return Value{}, errors.New("reflect.Value.Convert with incompatible type")
	}
	out := v.v.Convert(t.ReflectType())
	return Value{out}, nil
}

func (v Value) Comparable() bool {
	return v.v.Comparable()
}

func (v Value) Equal(x Value) (bool, error) {
	if !v.Comparable() || !x.Comparable() {
		return false, errors.New("reflect.Value.Equal with incomparable value")
	}
	return v.v.Equal(x.v), nil
}
