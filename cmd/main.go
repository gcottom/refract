package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gcottom/refract/genericdynamic"
	"github.com/gcottom/refract/godict"

	"github.com/gcottom/refract/refractutils"
)

func main() {
	d, err := genericdynamic.NewStructDefinition(genericdynamic.NewStructField("TestField", "string", ""))
	if err != nil {
		panic(err)
	}
	s, err := genericdynamic.NewSliceOfType(d)
	if err != nil {
		panic(err)
	}
	m, err := genericdynamic.NewTypeInstance(d)
	if err != nil {
		panic(err)
	}
	err = genericdynamic.SetStructFieldValue(m, "TestField", "testVal1")
	if err != nil {
		panic(err)
	}
	s, err = refractutils.Append(s, m)
	if err != nil {
		panic(err)
	}
	m, err = genericdynamic.NewTypeInstance(d)
	if err != nil {
		panic(err)
	}
	err = genericdynamic.SetStructFieldValue(m, "TestField", "testVal2")
	if err != nil {
		panic(err)
	}
	s, err = refractutils.Append(s, m)
	if err != nil {
		panic(err)
	}
	m, err = genericdynamic.NewTypeInstance(d)
	if err != nil {
		panic(err)
	}
	err = genericdynamic.SetStructFieldValue(m, "TestField", "banana")
	if err != nil {
		panic(err)
	}
	s, err = refractutils.Prepend(s, m)
	if err != nil {
		panic(err)
	}
	l, err := refractutils.Len(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v, %d\n", s, l)

	refractutils.RangeOverSlice(s, func(counter int, sliceItem any) {
		fmt.Println(sliceItem)
	})

	j := make([]int, 0)

	k, err := refractutils.Append(j, 1, 2, 3, []int{9, 8, 7})
	if err != nil {
		panic(err)
	}

	k, err = refractutils.Prepend(k, []int{10, 10, 10}, 9, 9, 9, []int{4, 4, 4})
	if err != nil {
		panic(err)
	}

	fmt.Println(k)

	si, err := refractutils.GetSliceIndex(s, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(si)

	si, err = refractutils.GetSliceIndex(s, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(si)

	tf, err := genericdynamic.GetStructFieldValue[string](si, "TestField")
	if err != nil {
		panic(err)
	}
	fmt.Println(tf)

	err = refractutils.SetSliceIndex(s, m, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	si, err = refractutils.GetSliceIndex(s, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(si)
	err = genericdynamic.SetStructFieldValue(si, "TestField", "not a banana")
	if err != nil {
		panic(err)
	}
	fmt.Println(si)

	err = refractutils.SetSliceIndex(s, si, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	s, err = refractutils.GetSliceIndexValue(s, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	mappy, err := genericdynamic.NewMapOfType("string", d)
	if err != nil {
		panic(err)
	}
	mpyval, err := genericdynamic.NewTypeInstance(d)
	if err != nil {
		panic(err)
	}
	err = genericdynamic.SetStructFieldValue(mpyval, "TestField", "I'm in the mappy")
	if err != nil {
		panic(err)
	}
	err = refractutils.PutMapIndex(mappy, "ind1", mpyval)
	if err != nil {
		panic(err)
	}
	mpyvaleval, err := refractutils.GetMapIndex(mappy, "ind1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("MapVal: %v\n", mpyvaleval)
	mpvalevalcopy, err := refractutils.GetMapIndexValue(mappy, "ind1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("MapVal: %v\n", mpvalevalcopy)

	typ := genericdynamic.GetReflectType(mpyval)
	fmt.Println(typ)

	ntype, err := genericdynamic.NewTypeInstance(typ)
	if err != nil {
		panic(err)
	}

	ntype2 := genericdynamic.GetReflectType(ntype)
	fmt.Println(ntype2)

	mappytype := genericdynamic.GetReflectType(s)
	fmt.Println("type:", mappytype)
	mappykind := reflect.ValueOf(s).Kind()
	fmt.Println("kind:", mappykind)
	mappytypekind := reflect.ValueOf(s).Type().Kind()
	fmt.Println("typekind:", mappytypekind)

	sl2 := make([]int, 0)
	sl2 = append(sl2, 1)
	sl1, err := refractutils.GetSliceIndexValue(sl2, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(sl1)

	var dct godict.SingleLevelJSONDict
	err = json.Unmarshal([]byte(`{"a":{"b":"c"},"d":[1,2,"e","f"],"g":[{"h":"i"},{"j":"k"}]}`), &dct)
	if err != nil {
		panic(err)
	}

	dk, err := dct.GetSlice("d").GetIndex(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dk)

	var dct2 godict.GoDict
	err = json.Unmarshal([]byte(`{"a":{"b":"c"},"d":[1,2,"e","f"],"g":[{"h":"i"},{"j":null}]}`), &dct2)
	if err != nil {
		panic(err)
	}

	dk2, err := dct2.Get("g").Get(1).Get("j").GetValue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dk2)

	jtxt, err := json.Marshal(dct2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jtxt))

	singleText, err := godict.UnmarshalSingleJSONKey[string]("a", []byte(`{"a":"b"}`))
	if err != nil {
		panic(err)
	}
	fmt.Println(singleText)

	var btext string
	if err = godict.UnmarshalSingleJSONKeyIntoPtr("a", []byte(`{"a":"j"}`), &btext); err != nil {
		panic(err)
	}

	fmt.Println(btext)
}
