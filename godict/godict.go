package godict

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/gcottom/refract/refractutils"
)

type GoDict struct {
	dict     JSONDict
	slice    JSONDictSlice
	val      any
	dataType reflect.Type
	isNull   bool
}

type JSONDict map[string]GoDict
type JSONDictSlice []GoDict

type SingleLevelJSONDict map[string]json.RawMessage
type SingleLevelJSONDictSlice []json.RawMessage

func (r *GoDict) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		r.val = nil
		r.dataType = nil
		r.isNull = true
		return nil
	}
	// Attempt to unmarshal into map
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err == nil {
		r.dict = make(JSONDict)
		for key, rawMsg := range tempMap {
			var nested GoDict
			if err := nested.UnmarshalJSON(rawMsg); err != nil {
				return err
			}
			r.dict[key] = nested
		}
		r.dataType = reflect.TypeOf(r.dict)
		return nil
	}

	// Attempt to unmarshal into a slice
	var tempSlice []json.RawMessage
	if err := json.Unmarshal(data, &tempSlice); err == nil {
		r.slice = make(JSONDictSlice, len(tempSlice))
		for i, rawMsg := range tempSlice {
			var nested GoDict
			if err := nested.UnmarshalJSON(rawMsg); err != nil {
				return err
			}
			r.slice[i] = nested
		}
		r.dataType = reflect.TypeOf(r.slice)
		return nil
	}

	// Fallback: Unmarshal into a generic value
	var tempVal any
	if err := json.Unmarshal(data, &tempVal); err == nil {
		r.val = tempVal
		r.dataType = reflect.TypeOf(tempVal)
		return nil
	}

	return errors.New("unable to unmarshal JSON into RefractDict")
}

func (r GoDict) MarshalJSON() ([]byte, error) {
	if r.isNull {
		return []byte("null"), nil
	}

	switch r.dataType {
	case reflect.TypeOf(JSONDict{}):
		return json.Marshal(r.dict)
	case reflect.TypeOf(JSONDictSlice{}):
		return json.Marshal(r.slice)
	default:
		return json.Marshal(r.val)
	}
}

func (r *GoDict) Get(index any) *GoDict {
	switch r.dataType {
	case reflect.TypeFor[JSONDict]():
		s, ok := index.(string)
		if !ok {
			return &GoDict{}
		}
		ret := r.dict[s]
		return &ret

	case reflect.TypeFor[JSONDictSlice]():
		i, ok := index.(int)
		if !ok {
			return &GoDict{}
		}
		if i >= len(r.slice) {
			return &GoDict{}
		}
		ret := r.slice[i]
		return &ret
	default:
		return &GoDict{}
	}
}

func (r *GoDict) GetValue() (any, error) {
	if r.dataType == nil {
		if r.isNull {
			return nil, nil
		}
		return nil, errors.New("unable to get value at current level, went too deep or missed a key at higher level")
	} else if r.dataType != reflect.TypeFor[JSONDict]() && r.dataType != reflect.TypeFor[JSONDictSlice]() {
		return r.val, nil
	} else {
		if r.dataType == reflect.TypeFor[JSONDict]() {
			return nil, errors.New("unable to get value at this level, value at this level is of type: map")
		}
		return nil, errors.New("unable to get value at this level, value at this level is of type: slice")
	}
}

// v should be a ptr
func (dict SingleLevelJSONDict) UnmarshalFromKey(key string, v any) error {
	d, ok := dict[key]
	if !ok {
		return errors.New("key not found in dict")
	}
	return json.Unmarshal(d, v)
}

func (dict SingleLevelJSONDict) GetKey(key string) (any, error) {
	d, ok := dict[key]
	if !ok {
		return nil, errors.New("key not found in dict")
	}
	var out any
	err := json.Unmarshal(d, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (dict SingleLevelJSONDict) GetDict(key string) SingleLevelJSONDict {
	d, ok := dict[key]
	if !ok {
		return nil
	}
	var out SingleLevelJSONDict
	err := json.Unmarshal(d, &out)
	if err != nil {
		return nil
	}
	return out
}

func (dict SingleLevelJSONDict) GetSlice(key string) SingleLevelJSONDictSlice {
	d, ok := dict[key]
	if !ok {
		return nil
	}
	var out SingleLevelJSONDictSlice
	err := json.Unmarshal(d, &out)
	if err != nil {
		return nil
	}
	return out
}

// v should be a ptr
func (dictSlice SingleLevelJSONDictSlice) UnmarshalJSONFromIndex(index int, v any) error {
	val, err := refractutils.GetSliceIndex(dictSlice, index)
	if err != nil {
		return err
	}
	return json.Unmarshal(val.(json.RawMessage), v)
}

func (dictSlice SingleLevelJSONDictSlice) GetIndex(index int) (any, error) {
	val, err := refractutils.GetSliceIndex(dictSlice, index)
	if err != nil {
		return nil, err
	}
	var out any
	err = json.Unmarshal(val.(json.RawMessage), &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (dictSlice SingleLevelJSONDictSlice) GetDictAtIndex(index int) SingleLevelJSONDict {
	val, err := refractutils.GetSliceIndex(dictSlice, index)
	if err != nil {
		return nil
	}
	var out SingleLevelJSONDict
	err = json.Unmarshal(val.(json.RawMessage), &out)
	if err != nil {
		return nil
	}
	return out
}
