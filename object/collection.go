package object

import (
	"strings"
)

// Collection 集合类型
type Collection[V any] struct {
	items *Map[V]
}

// NewCollection 创建新的集合
func NewCollection[V any](items *Map[V]) *Collection[V] {
	if items == nil {
		items = &Map[V]{}
	}
	return &Collection[V]{
		items: items,
	}
}

// All 返回所有项目
func (c *Collection[V]) All() *Map[V] {
	return c.items
}

// Get 使用"点"表示法从集合中获取项目
func (c *Collection[V]) Get(key string, defaultValue ...V) V {
	if key == "" {
		return getDefaultValue(defaultValue...)
	}

	// 处理简单键
	if !strings.Contains(key, ".") {
		val := c.items.Get(key)
		if IsZero(val) {
			return getDefaultValue(defaultValue...)
		}
		return val
	}

	// 处理嵌套键
	segments := strings.Split(key, ".")
	current := c.items

	for _, segment := range segments[:len(segments)-1] {
		val := current.Get(segment)
		if IsZero(val) {
			return getDefaultValue(defaultValue...)
		}

		// 尝试将值转换为 Map
		if nestedMap, ok := any(val).(*Map[V]); ok {
			current = nestedMap
		} else {
			return getDefaultValue(defaultValue...)
		}
	}

	val := current.Get(segments[len(segments)-1])
	if IsZero(val) {
		return getDefaultValue(defaultValue...)
	}
	return val
}

// Set 设置集合中的值
func (c *Collection[V]) Set(key string, value V) {
	if key == "" {
		return
	}

	segments := strings.Split(key, ".")
	current := c.items

	// 处理嵌套路径
	for i := 0; i < len(segments)-1; i++ {
		segment := segments[i]
		val := current.Get(segment)

		if IsZero(val) {
			newMap := &Map[V]{}
			(*current)[segment] = any(newMap).(V)
			current = newMap
			continue
		}

		if nestedMap, ok := any(val).(*Map[V]); ok {
			current = nestedMap
		} else {
			newMap := &Map[V]{}
			(*current)[segment] = any(newMap).(V)
			current = newMap
		}
	}

	(*current)[segments[len(segments)-1]] = value
}

// Has 检查键是否存在
func (c *Collection[V]) Has(key string) bool {
	return c.items.Has(key)
}

// Count 返回集合中的项目数
func (c *Collection[V]) Count() int {
	return len(*c.items)
}

// ToMap 将集合转换为 Map
func (c *Collection[V]) ToMap() *Map[V] {
	return c.All()
}

func (c *Collection[V]) ToJson() (string, error) {
	return JsonEncode(c.items)
}

func (c *Collection[V]) String() string {
	strJson, _ := c.ToJson()
	return strJson
}
