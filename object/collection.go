package object

import (
	"strings"
)

// Collection 集合类型
type Collection struct {
	items *HashMap
}

// NewCollection 创建新的集合
func NewCollection(items *HashMap) *Collection {
	if items == nil {
		items = &HashMap{}
	}
	return &Collection{
		items: items,
	}
}

// All 返回所有项目
func (c *Collection) All() *HashMap {
	return c.items
}

// Get 使用"点"表示法从集合中获取项目
func (c *Collection) Get(key string, defaultValue ...any) any {
	if key == "" {
		return GetDefaultValue(defaultValue...)
	}

	// 处理简单键
	if !strings.Contains(key, ".") {
		val := c.items.Get(key)
		if IsZero(val) {
			return GetDefaultValue(defaultValue...)
		}
		return val
	}

	// 处理嵌套键
	segments := strings.Split(key, ".")
	current := c.items

	for _, segment := range segments[:len(segments)-1] {
		val := current.Get(segment)
		if IsZero(val) {
			return GetDefaultValue(defaultValue...)
		}

		// 尝试将值转换为 Map
		if nestedMap, ok := any(val).(*HashMap); ok {
			current = nestedMap
		} else {
			return GetDefaultValue(defaultValue...)
		}
	}

	val := current.Get(segments[len(segments)-1])
	if IsZero(val) {
		return GetDefaultValue(defaultValue...)
	}
	return val
}

// Set 设置集合中的值
func (c *Collection) Set(key string, value any) {
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
			newMap := &HashMap{}
			(*current)[segment] = any(newMap)
			current = newMap
			continue
		}

		if nestedMap, ok := any(val).(*HashMap); ok {
			current = nestedMap
		} else {
			newMap := &HashMap{}
			(*current)[segment] = any(newMap)
			current = newMap
		}
	}

	(*current)[segments[len(segments)-1]] = value
}

// Has 检查键是否存在
func (c *Collection) Has(key string) bool {
	return c.items.Has(key)
}

// Count 返回集合中的项目数
func (c *Collection) Count() int {
	return len(*c.items)
}

// ToMap 将集合转换为 Map
func (c *Collection) ToMap() *HashMap {
	return c.All()
}

func (c *Collection) ToJson() (string, error) {
	return JsonEncode(c.items)
}

func (c *Collection) String() string {
	strJson, _ := c.ToJson()
	return strJson
}
