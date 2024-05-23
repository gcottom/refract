package main

import (
	"fmt"

	"github.com/gcottom/refract"
)

func main() {
	d := refract.NewStructDefinition(refract.NewStructField("TestField", "string", ""))
	s := refract.NewSliceOfStruct(d)

	m := refract.NewStructInstance(d)
	err := refract.SetStructFieldValue(m, "TestField", "testVal1")
	if err != nil {
		panic(err)
	}
	s, err = refract.Append(s, m)
	if err != nil {
		panic(err)
	}
	m = refract.NewStructInstance(d)
	err = refract.SetStructFieldValue(m, "TestField", "testVal2")
	if err != nil {
		panic(err)
	}
	s, err = refract.Append(s, m)
	if err != nil {
		panic(err)
	}
	m = refract.NewStructInstance(d)
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

	mappy := refract.NewMapOfStruct("string", d)
	mpyval := refract.NewStructInstance(d)
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
}
