package refract

import (
	"encoding/json"
	"errors"
)

type JSONDict map[string]json.RawMessage
type JSONDictSlice []json.RawMessage

func AssertJSONDict(dict any) (JSONDict, error) {
	d, ok := dict.(JSONDict)
	if !ok {
		return nil, errors.New("unable to assert as JSONDict")
	}
	return d, nil
}

func AssertJSONDictSlice(val any) (JSONDictSlice, error) {
	s, ok := val.(JSONDictSlice)
	if !ok {
		return nil, errors.New("unable to assert as JSONDictSlice")
	}
	return s, nil
}

func (dict JSONDict) GetKey(key string) (any, error) {
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

func (dict JSONDict) GetDict(key string) JSONDict {
	d, ok := dict[key]
	if !ok {
		return nil
	}
	var out JSONDict
	err := json.Unmarshal(d, &out)
	if err != nil {
		return nil
	}
	return out
}

func (dict JSONDict) GetSlice(key string) JSONDictSlice {
	d, ok := dict[key]
	if !ok {
		return nil
	}
	var out JSONDictSlice
	err := json.Unmarshal(d, &out)
	if err != nil {
		return nil
	}
	return out
}

func (dictSlice JSONDictSlice) GetIndex(index int) (any, error) {
	val, err := GetSliceIndex(dictSlice, index)
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

func (dictSlice JSONDictSlice) GetDictAtIndex(index int) JSONDict {
	val, err := GetSliceIndex(dictSlice, index)
	if err != nil {
		return nil
	}
	var out JSONDict
	err = json.Unmarshal(val.(json.RawMessage), &out)
	if err != nil {
		return nil
	}
	return out
}
