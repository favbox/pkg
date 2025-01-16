package object

import "reflect"

// IsZero 工具函数：判断空值
func IsZero(v any) bool {
	if v == nil {
		return true
	}
	//if lo.IsEmpty(v) {
	//	return true
	//}
	switch t := v.(type) {
	case string:
		return t == ""
	case bool:
		return !t
	case int, int8, int16, int32, int64:
		return t == 0
	case uint, uint8, uint16, uint32, uint64:
		return t == 0
	case float32, float64:
		return t == 0.0
	case map[any]any:
		return len(t) == 0
	case []any:
		return len(t) == 0
	case []byte:
		return len(t) == 0
	case chan any:
		return false // channel 不判断 None
	default:
		// 对于其他类型，通过 reflect 检查零值
		val := reflect.ValueOf(v)
		if !val.IsValid() {
			return true
		}
		switch val.Kind() {
		case reflect.Map, reflect.Slice, reflect.Array, reflect.Chan:
			return val.Len() == 0
		case reflect.Ptr, reflect.Interface:
			return val.IsNil()
		default:
			return val.IsZero()
		}
	}
}
