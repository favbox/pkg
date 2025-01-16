package object

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/bytedance/sonic"
)

// Map 是基础泛型 Map 类型
type Map[V any] map[string]V

// HashMap 是 Map[any] 的类型别名
type HashMap = Map[any]

// StringMap 是 Map[string] 的类型别名
type StringMap = Map[string]

// MergeMap 通用的合并方法
func MergeMap[V any](target *Map[V], sources ...*Map[V]) *Map[V] {
	if target == nil {
		target = &Map[V]{}
	}
	for _, source := range sources {
		if source == nil {
			continue
		}
		for k, v := range *source {
			if IsZero((*target)[k]) && !IsZero(v) {
				(*target)[k] = v
			}
		}
	}
	return target
}

// ReplaceMap 通用的替换方法
func ReplaceMap[V any](target *Map[V], sources ...*Map[V]) *Map[V] {
	if target == nil {
		target = &Map[V]{}
	}
	for _, source := range sources {
		if source != nil {
			for k, v := range *source {
				(*target)[k] = v
			}
		}
	}
	return target
}

// FilterEmptyMap 通用的过滤方法
func FilterEmptyMap[V any](mapData *Map[V]) *Map[V] {
	filtered := &Map[V]{}
	for k, v := range *mapData {
		if !IsZero(v) {
			(*filtered)[k] = v
		}
	}
	return filtered
}

// HashMapToStringMap HashMap 转 StringMap
func HashMapToStringMap(obj *HashMap) (*StringMap, error) {
	newMap := &StringMap{}

	if obj == nil {
		return newMap, nil
	}

	for k, v := range *obj {
		(*newMap)[k] = fmt.Sprintf("%v", v)
	}

	return newMap, nil
}

// StructToHashMap 结构体转 HashMap
func StructToHashMap(obj any) (*HashMap, error) {
	data, err := sonic.Marshal(obj)
	if err != nil {
		return nil, err
	}

	newMap := &HashMap{}
	err = sonic.Unmarshal(data, newMap)
	return newMap, err
}

// HashMapToStructure HashMap 转结构体
func HashMapToStructure(mapObj *HashMap, obj any) error {
	data, err := sonic.Marshal(mapObj)
	if err != nil {
		return err
	}
	return sonic.Unmarshal(data, obj)
}

// StructToHashMapWithXML 结构体转 HashMap (使用 XML 标签)
func StructToHashMapWithXML(obj any) (*HashMap, error) {
	newMap := &HashMap{}
	if obj == nil {
		return newMap, nil
	}

	e := reflect.ValueOf(obj).Elem()
	for i := 0; i < e.NumField(); i++ {
		field := e.Field(i).Interface()
		key := e.Type().Field(i).Tag.Get("xml")
		(*newMap)[key] = field
	}

	return newMap, nil
}

// InHash 在 HashMap 中查找值
func InHash(val any, hash *HashMap) (exists bool, key string) {
	for k, v := range *hash {
		if reflect.DeepEqual(val, v) {
			return true, k
		}
	}
	return false, ""
}

// Get 获取 Map 的值。
func (m *Map[V]) Get(key string) V {
	return (*m)[key]
}

func (m *Map[V]) Has(key string) bool {
	_, exists := (*m)[key]
	return exists
}

// GetMapKV 获取 Map 的键值对
func GetMapKV[V any](maps Map[V]) (keys []string, values []any) {
	mapLen := len(maps)
	keys = make([]string, 0, mapLen)
	values = make([]any, 0, mapLen)

	for k, v := range maps {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

// GetJoinedWithKSort 将 StringMap 中的键值对按照键的字母顺序排序后拼接成字符串
func GetJoinedWithKSort[V any](params *Map[V], keepEmpty ...bool) string {
	if params == nil || len(*params) == 0 {
		return ""
	}

	if len(keepEmpty) == 0 || !keepEmpty[0] {
		params = FilterEmptyMap(params)
	}

	// ksort
	var keys []string
	for k := range *params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// join
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(fmt.Sprintf("%v", (*params)[k]))
		sb.WriteString("&")
	}
	strJoined := sb.String()
	return strJoined[:len(strJoined)-1]
	//
	//var strJoined string
	//
	//// ksort
	//var keys []string
	//for k := range *params {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys)
	//
	//// join
	//for _, k := range keys {
	//	strJoined += k + "=" + (*params)[k] + "&"
	//}
	//
	//strJoined = strJoined[0 : len(strJoined)-1]
	//
	//return strJoined
}
