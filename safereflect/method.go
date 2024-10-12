package safereflect

type Method struct {
	Name    string
	PkgPath string
	Type    Type
	Func    Value
	Index   int
}

func (m Method) IsExported() bool {
	return m.PkgPath == ""
}
