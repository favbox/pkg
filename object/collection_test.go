package object

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCollection(t *testing.T) {
	// 创建集合
	collection := NewCollection(&HashMap{
		"user": &HashMap{
			"name": "张三",
			"age":  25,
		},
	})

	// 获取嵌套值
	name := collection.Get("user.name")
	assert.Equal(t, "张三", name)

	// 设置嵌套值
	collection.Set("user.email", "zhangsan@example.com")
	// 获取嵌套值
	email := collection.Get("user.email")
	assert.Equal(t, "zhangsan@example.com", email)
}
