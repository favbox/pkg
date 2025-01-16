package object

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MergeStringMap(t *testing.T) {
	base := &StringMap{
		"key1": "123",
		"key2": "xxx",
		"key3": "",
		"key4": "101112",
	}

	toMap := &StringMap{
		"key1": "",
		"key2": "456",
		"key3": "789",
	}

	toMap = MergeMap(toMap, base)

	assert.EqualValues(t, &StringMap{
		"key1": "123",
		"key2": "456",
		"key3": "789",
		"key4": "101112",
	}, toMap)

}

func Test_ReplaceStringMapRecursive(t *testing.T) {
	base := &StringMap{
		"key1": "123",
		"key2": "456",
		"key3": "789",
		"key4": "nil",
	}

	base2 := &StringMap{
		"key1": "456",
		"key2": "base456",
		"key3": "",
		"key4": "nil",
		"key5": "&StringMap{}",
	}

	toMap := &StringMap{
		"key2": "",
		"key3": "123",
		"key4": "",
		"key5": "nil",
		"key6": "&StringMap{}",
	}

	toMap = ReplaceMap(toMap, base, base2)

	assert.EqualValues(t, &StringMap{
		"key1": "456",
		"key2": "base456",
		"key3": "",
		"key4": "nil",
		"key5": "&StringMap{}",
		"key6": "&StringMap{}",
	}, toMap)

}

func TestMergeMap(t *testing.T) {
	tests := []struct {
		name     string
		target   *Map[string]
		sources  []*Map[string]
		expected *Map[string]
	}{
		{
			name:   "合并空 map",
			target: &Map[string]{"a": "1"},
			sources: []*Map[string]{
				{"b": "2"},
				nil,
				{"c": "3"},
			},
			expected: &Map[string]{"a": "1", "b": "2", "c": "3"},
		},
		{
			name:   "已存在的值不覆盖",
			target: &Map[string]{"a": "1", "b": "2"},
			sources: []*Map[string]{
				{"b": "3", "c": "4"},
			},
			expected: &Map[string]{"a": "1", "b": "2", "c": "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeMap(tt.target, tt.sources...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReplaceMap(t *testing.T) {
	tests := []struct {
		name     string
		target   *Map[string]
		sources  []*Map[string]
		expected *Map[string]
	}{
		{
			name:   "替换现有值",
			target: &Map[string]{"a": "1", "b": "2"},
			sources: []*Map[string]{
				{"b": "3", "c": "4"},
			},
			expected: &Map[string]{"a": "1", "b": "3", "c": "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceMap(tt.target, tt.sources...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFilterEmptyMap(t *testing.T) {
	input := &HashMap{
		"empty_string": "",
		"valid_string": "hello",
		"zero_int":     0,
		"valid_int":    42,
		"nil_value":    nil,
		"empty_slice":  []string{},
		"valid_slice":  []string{"hello"},
	}

	expected := &HashMap{
		"valid_string": "hello",
		"valid_int":    42,
		"valid_slice":  []string{"hello"},
	}

	result := FilterEmptyMap(input)
	assert.Equal(t, expected, result)
}

type TestStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestStructToHashMap(t *testing.T) {
	input := &TestStruct{
		Name:  "test",
		Value: 123,
	}

	expected := &HashMap{
		"name":  "test",
		"value": float64(123), // JSON 数字会被解析为 float64
	}

	result, err := StructToHashMap(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestHashMapToStructure(t *testing.T) {
	input := &HashMap{
		"name":  "test",
		"value": float64(123),
	}

	expected := &TestStruct{
		Name:  "test",
		Value: 123,
	}

	var result TestStruct
	err := HashMapToStructure(input, &result)
	assert.NoError(t, err)
	assert.Equal(t, expected, &result)
}

type XMLTestStruct struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func TestStructToHashMapWithXML(t *testing.T) {
	input := &XMLTestStruct{
		Name:  "test",
		Value: "123",
	}

	expected := &HashMap{
		"name":  "test",
		"value": "123",
	}

	result, err := StructToHashMapWithXML(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestInHash(t *testing.T) {
	hash := &HashMap{
		"key1": "value1",
		"key2": 123,
	}

	tests := []struct {
		name          string
		searchValue   any
		expectedFound bool
		expectedKey   string
	}{
		{
			name:          "查找存在的字符串值",
			searchValue:   "value1",
			expectedFound: true,
			expectedKey:   "key1",
		},
		{
			name:          "查找存在的数字值",
			searchValue:   123,
			expectedFound: true,
			expectedKey:   "key2",
		},
		{
			name:          "查找不存在的值",
			searchValue:   "nonexistent",
			expectedFound: false,
			expectedKey:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, key := InHash(tt.searchValue, hash)
			assert.Equal(t, tt.expectedFound, found)
			assert.Equal(t, tt.expectedKey, key)
		})
	}
}

func TestGetHashMapKV(t *testing.T) {
	input := StringMap{
		"key1": "value1",
		"key2": "value2",
	}

	keys, values := GetMapKV(input)

	// 由于 map 迭代顺序是随机的，我们需要检查结果集合而不是具体顺序
	assert.ElementsMatch(t, []string{"key1", "key2"}, keys)
	assert.ElementsMatch(t, []any{"value1", "value2"}, values)
	assert.Equal(t, len(keys), len(values))
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected bool
	}{
		{
			name:     "nil 值",
			value:    nil,
			expected: true,
		},
		{
			name:     "空字符串",
			value:    "",
			expected: true,
		},
		{
			name:     "非空字符串",
			value:    "hello",
			expected: false,
		},
		{
			name:     "零值整数",
			value:    0,
			expected: true,
		},
		{
			name:     "非零值整数",
			value:    42,
			expected: false,
		},
		{
			name:     "空切片",
			value:    []string{},
			expected: true,
		},
		{
			name:     "非空切片",
			value:    []string{"hello"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsZero(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetJoinedWithKSort(t *testing.T) {
	input := &Map[any]{
		"key1": "value1",
		"key3": 123.23,
		"key0": true,
		"key4": nil,
		"key7": 0,
		"key2": "value2",
	}

	s := GetJoinedWithKSort(input, true)
	fmt.Println(s)
}

func TestMap_Get(t *testing.T) {
	m := HashMap{
		"key1": "value1",
		"key2": 2,
		"b":    true,
	}
	assert.Equal(t, "value1", m.Get("key1"))
	assert.Equal(t, 2, m.Get("key2"))
	assert.Equal(t, true, m.Get("b"))
}

func Test_MergeHashMap(t *testing.T) {
	base := &HashMap{
		"key1": 123,
		"key2": "456",
		"key3": "",
		"key4": nil,
		"key5": &StringMap{},
		"key6": &HashMap{},
	}

	toMap := &HashMap{
		"key2": "",
		"key3": "123",
		"key4": &HashMap{},
		"key5": nil,
		"key6": &StringMap{},
	}

	toMap = MergeMap(toMap, base)

	assert.EqualValues(t, &HashMap{
		"key1": 123,
		"key2": "456",
		"key3": "123",
		"key4": &HashMap{},
		"key5": &StringMap{},
		"key6": &StringMap{},
	}, toMap)

}

func Test_ReplaceHashMapRecursive(t *testing.T) {
	base := &HashMap{
		"key1": 123,
		"key2": "456",
		"key3": "789",
		"key4": nil,
		"key5": map[string]int{},
		"key6": &map[string]float32{},
	}

	base2 := &HashMap{
		"key1": 456,
		"key2": "base456",
		"key3": "",
		"key4": nil,
		"key5": &StringMap{},
	}

	toMap := &HashMap{
		"key2": "",
		"key3": "123",
		"key4": &HashMap{},
		"key5": nil,
		"key6": &StringMap{},
	}

	toMap = ReplaceMap(toMap, base, base2)

	assert.EqualValues(t, &HashMap{
		"key1": 456,
		"key2": "base456",
		"key3": "",
		"key4": nil,
		"key5": &StringMap{},
		"key6": &map[string]float32{},
	}, toMap)

}
