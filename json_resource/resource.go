package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strconv"
)

// Resource provides a wrapper for arbitrary JSON objects that adds methods to access properties.
type Resource struct {
	// TODO: restrict data to map[string]interface{}
	data interface{}
}

// NewResource reads data from input and returns resources if it's not empty.
func NewResource(data []byte) (*Resource, error) {
	if len(data) == 0 {
		return EmptyResource(), nil
	}
	r := new(Resource)
	if err := r.FromBytes(data); err != nil {
		return nil, err
	}
	return r, nil
}

// EmptyResource returns an empty resource wrapper whose data is a empty map.
func EmptyResource() *Resource {
	r := new(Resource)
	r.data = make(map[string]interface{})
	return r
}

// NewResourceFromString returns a *Resource by decoding from an string
func NewResourceFromString(data string) (*Resource, error) {
	return NewResource([]byte(data))
}

// NewResourceFromReader returns a *Resource by decoding from an io.Reader
func NewResourceFromReader(r io.Reader) (*Resource, error) {
	rs := new(Resource)
	dec := json.NewDecoder(r)
	dec.UseNumber()
	err := dec.Decode(&rs.data)
	return rs, err
}

// FromBytes reads data from json bytes.
func (r *Resource) FromBytes(data []byte) error {
	r.data = make(map[string]interface{})

	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.UseNumber()
	return dec.Decode(&r.data)
}

// ToBytes marshal resource to json bytes.
func (r *Resource) ToBytes() ([]byte, error) {
	if r == nil || r.data == nil {
		return nil, nil
	}

	return r.Encode()
}

// ToString marshals data to json bytes and convert it to string
func (r *Resource) ToString() (string, error) {
	b, err := r.ToBytes()
	if err != nil {
		return "", err
	}
	return string(b), err
}

// Encode marshals data to json.
func (r *Resource) Encode() ([]byte, error) {
	return r.MarshalJSON()
}

// EncodePretty returns its marshaled data as `[]byte` with indentation
func (r *Resource) EncodePretty() ([]byte, error) {
	return json.MarshalIndent(&r.data, "", "  ")
}

// MarshalJSON implements the json.Marshaler interface.
func (r *Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (r *Resource) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(p))
	dec.UseNumber()
	return dec.Decode(&r.data)
}

// Set modifies `Json` map by `key` and `value`
// Useful for changing single key/value in a `Json` object easily.
func (r *Resource) Set(key string, val interface{}) {
	m, err := r.Map()
	if err != nil {
		return
	}
	m[key] = val
}

// SetPath modifies `Json`, recursively checking/creating map keys for the supplied path,
// and then finally writing in the value
func (r *Resource) SetPath(branch []string, val interface{}) {
	if len(branch) == 0 {
		r.data = val
		return
	}

	// in order to insert our branch, we need map[string]interface{}
	if _, ok := (r.data).(map[string]interface{}); !ok {
		// have to replace with something suitable
		r.data = make(map[string]interface{})
	}
	curr := r.data.(map[string]interface{})

	for i := 0; i < len(branch)-1; i++ {
		b := branch[i]
		// key exists?
		if _, ok := curr[b]; !ok {
			n := make(map[string]interface{})
			curr[b] = n
			curr = n
			continue
		}

		// make sure the value is the right sort of thing
		if _, ok := curr[b].(map[string]interface{}); !ok {
			// have to replace with something suitable
			n := make(map[string]interface{})
			curr[b] = n
		}

		curr = curr[b].(map[string]interface{})
	}

	// add remaining k/v
	curr[branch[len(branch)-1]] = val
}

// Del modifies `Json` map by deleting `key` if it is present.
func (r *Resource) Del(key string) {
	m, err := r.Map()
	if err != nil {
		return
	}
	delete(m, key)
}

// Get returns a pointer to a new `Json` object
// for `key` in its `map` representation
//
// useful for chaining operations (to traverse a nested JSON):
//    resource.Get("top_level").Get("dict").Get("value").Int()
func (r *Resource) Get(key string) *Resource {
	m, err := r.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Resource{val}
		}
	}
	return &Resource{nil}
}

// GetPath searches for the item as specified by the branch
// without the need to deep dive using Get()'s.
//
//   resource.GetPath("top_level", "dict")
func (r *Resource) GetPath(branch ...string) *Resource {
	// TODO: returns error if current node not found
	rin := r
	for _, p := range branch {
		rin = rin.Get(p)
	}
	return rin
}

// GetIndex returns a pointer to a new `Json` object
// for `index` in its `array` representation
//
// this is the analog to Get when accessing elements of
// a json array instead of a json object:
//    resource.Get("top_level").Get("array").GetIndex(1).Get("key").Int()
func (r *Resource) GetIndex(index int) *Resource {
	a, err := r.Array()
	if err == nil {
		if len(a) > index {
			return &Resource{a[index]}
		}
	}
	return &Resource{nil}
}

// CheckGet returns a pointer to a new `Json` object and
// a `bool` identifying success or failure
//
// useful for chained operations when success is important:
//    if data, ok := resource.Get("top_level").CheckGet("inner"); ok {
//        log.Println(data)
//    }
func (r *Resource) CheckGet(key string) (*Resource, bool) {
	m, err := r.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Resource{val}, true
		}
	}
	return nil, false
}

// Map type asserts to `map`
func (r *Resource) Map() (map[string]interface{}, error) {
	if m, ok := (r.data).(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map[string]interface{} failed")
}

// Array type asserts to an `array`
func (r *Resource) Array() ([]interface{}, error) {
	if a, ok := (r.data).([]interface{}); ok {
		return a, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

// Bool type asserts to `bool`
func (r *Resource) Bool() (bool, error) {
	if s, ok := (r.data).(bool); ok {
		return s, nil
	}
	return false, errors.New("type assertion to bool failed")
}

// String type asserts to `string`
func (r *Resource) String() (string, error) {
	if s, ok := (r.data).(string); ok {
		return s, nil
	}
	return "", errors.New("type assertion to string failed")
}

// Bytes type asserts to `[]byte`
func (r *Resource) Bytes() ([]byte, error) {
	if s, ok := (r.data).(string); ok {
		return []byte(s), nil
	}
	return nil, errors.New("type assertion to []byte failed")
}

// StringArray type asserts to an `array` of `string`
func (r *Resource) StringArray() ([]string, error) {
	arr, err := r.Array()
	if err != nil {
		return nil, err
	}
	retArr := make([]string, 0, len(arr))
	for _, a := range arr {
		if a == nil {
			retArr = append(retArr, "")
			continue
		}
		s, ok := a.(string)
		if !ok {
			return nil, err
		}
		retArr = append(retArr, s)
	}
	return retArr, nil
}

// IntArray type asserts to an `array` of `int`
func (r *Resource) IntArray() ([]int, error) {
	arr, err := r.Array()
	if err != nil {
		return nil, err
	}
	retArr := make([]int, 0, len(arr))
	for _, a := range arr {
		if a == nil {
			retArr = append(retArr, 0)
			continue
		}
		var s int
		switch n := a.(type) {
		case json.Number:
			v, err := n.Int64()
			if err != nil {
				return nil, err
			}
			s = int(v)
		case int:
			s = int(reflect.ValueOf(a).Int())
		}
		retArr = append(retArr, s)
	}
	return retArr, nil
}

// Float64 coerces into a float64
func (r *Resource) Float64() (float64, error) {
	switch n := r.data.(type) {
	case json.Number:
		return n.Float64()
	case float32, float64:
		return reflect.ValueOf(r.data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(r.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(r.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Int coerces into an int
func (r *Resource) Int() (int, error) {
	switch n := r.data.(type) {
	case json.Number:
		i, err := n.Int64()
		if err != nil {
			return 0, err
		}
		return int(i), nil
	case float32, float64:
		return int(reflect.ValueOf(r.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(r.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(r.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Int64 coerces into an int64
func (r *Resource) Int64() (int64, error) {
	switch n := r.data.(type) {
	case json.Number:
		return n.Int64()
	case float32, float64:
		return int64(reflect.ValueOf(r.data).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(r.data).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(r.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Uint64 coerces into an uint64
func (r *Resource) Uint64() (uint64, error) {
	switch n := r.data.(type) {
	case json.Number:
		return strconv.ParseUint(n.String(), 10, 64)
	case float32, float64:
		return uint64(reflect.ValueOf(r.data).Float()), nil
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(r.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(r.data).Uint(), nil
	}
	return 0, errors.New("invalid value type")
}

// MustString try to get value from key, and value must be string
func (r *Resource) MustString(key string) (string, error) {
	res, exist := r.CheckGet(key)
	if !exist {
		return "", errors.New("key not found")
	}

	return res.String()
}

func (r *Resource) Length() (int, error) {
	arrs, err := r.Array()
	if err != nil {
		return 0, errors.New("key is not array")
	}

	return len(arrs), nil
}
