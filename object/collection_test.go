package object

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewCollection(t *testing.T) {
	// 创建集合
	collection := NewCollection(&HashMap{
		"user": &HashMap{
			"name": "张三",
			"age":  25,
			"dob":  time.Date(2020, time.April, 10, 0, 0, 0, 0, time.UTC),
		},
		"notes": &HashMap{
			"note1": "这是第一条笔记",
			"note2": "这是第二条笔记",
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

	// 获取嵌套值
	dob := collection.Get("user.dob").(time.Time)
	assert.Equal(t, 2020, dob.Year())

	// 获取不存在的值
	assert.Equal(t, "默认值", collection.Get("user.xxx", "默认值"))

	// 判断集合条目数量
	assert.Equal(t, 2, collection.Count())
}

func Test_Collection_Set_AND_Get(t *testing.T) {
	collectionTest := NewCollection(&HashMap{
		"gun": "model",
	})

	collectionTest.Set("weapon.bullet", 100)
	collectionTest.Set("weapon.shield.strength", "strong")

	bulletCount := collectionTest.Get("weapon.bullet", 0)
	if bulletCount != 100 {
		t.Error("get bullet error")
		fmt.Println(bulletCount)
	}

	shieldStrength := collectionTest.Get("weapon.shield.strength", "")
	if shieldStrength != "strong" {
		t.Error("get shield error")
		fmt.Println(shieldStrength)
	}
}
