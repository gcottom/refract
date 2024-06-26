package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gcottom/refract"
)

func main() {
	d := refract.NewStructDefinition(refract.NewStructField("TestField", "string", ""))
	s := refract.NewSliceOfType(d)

	m := refract.NewTypeInstance(d)
	err := refract.SetStructFieldValue(m, "TestField", "testVal1")
	if err != nil {
		panic(err)
	}
	s, err = refract.Append(s, m)
	if err != nil {
		panic(err)
	}
	m = refract.NewTypeInstance(d)
	err = refract.SetStructFieldValue(m, "TestField", "testVal2")
	if err != nil {
		panic(err)
	}
	s, err = refract.Append(s, m)
	if err != nil {
		panic(err)
	}
	m = refract.NewTypeInstance(d)
	err = refract.SetStructFieldValue(m, "TestField", "banana")
	if err != nil {
		panic(err)
	}
	s, err = refract.Prepend(s, m)
	if err != nil {
		panic(err)
	}
	l, err := refract.Len(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v, %d\n", s, l)

	refract.RangeOverSlice(s, func(counter int, sliceItem any) {
		fmt.Println(sliceItem)
	})

	j := make([]int, 0)

	k, err := refract.Append(j, 1, 2, 3, []int{9, 8, 7})
	if err != nil {
		panic(err)
	}

	k, err = refract.Prepend(k, []int{10, 10, 10}, 9, 9, 9, []int{4, 4, 4})
	if err != nil {
		panic(err)
	}

	fmt.Println(k)

	si, err := refract.GetSliceIndex(s, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(si)

	si, err = refract.GetSliceIndex(s, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(si)

	tf, err := refract.GetStructFieldValue[string](si, "TestField")
	if err != nil {
		panic(err)
	}
	fmt.Println(tf)

	err = refract.SetSliceIndex(s, m, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	si, err = refract.GetSliceIndex(s, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(si)
	err = refract.SetStructFieldValue(si, "TestField", "not a banana")
	if err != nil {
		panic(err)
	}
	fmt.Println(si)

	err = refract.SetSliceIndex(s, si, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	s, err = refract.GetSliceIndexValue(s, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	mappy := refract.NewMapOfType("string", d)
	mpyval := refract.NewTypeInstance(d)
	err = refract.SetStructFieldValue(mpyval, "TestField", "I'm in the mappy")
	if err != nil {
		panic(err)
	}
	err = refract.PutMapIndex(mappy, "ind1", mpyval)
	if err != nil {
		panic(err)
	}
	mpyvaleval, err := refract.GetMapIndex(mappy, "ind1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("MapVal: %v\n", mpyvaleval)
	mpvalevalcopy, err := refract.GetMapIndexValue(mappy, "ind1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("MapVal: %v\n", mpvalevalcopy)

	typ := refract.GetReflectType(mpyval)
	fmt.Println(typ)

	ntype := refract.NewTypeInstance(typ)

	ntype2 := refract.GetReflectType(ntype)
	fmt.Println(ntype2)

	mappytype := refract.GetReflectType(s)
	fmt.Println("type:", mappytype)
	mappykind := reflect.ValueOf(s).Kind()
	fmt.Println("kind:", mappykind)
	mappytypekind := reflect.ValueOf(s).Type().Kind()
	fmt.Println("typekind:", mappytypekind)

	sl2 := make([]int, 0)
	sl2 = append(sl2, 1)
	sl1, err := refract.GetSliceIndexValue(sl2, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(sl1)

	var dct refract.JSONDict
	err = json.Unmarshal([]byte(`{"a":{"b":"c"},"d":[1,2,"e","f"],"g":[{"h":"i"},{"j":"k"}]}`), &dct)
	if err != nil {
		panic(err)
	}

	dk, err := dct.GetSlice("d").GetIndex(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dk)
}
