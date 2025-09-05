package iterator

import (
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestIterator 测试迭代器模式的功能
func TestIterator(t *testing.T) {
	// 创建具体聚合对象
	userCollection := &UserCollection{}
	userCollection.Add(&User{Name: "Alice", Age: 30})
	userCollection.Add(&User{Name: "Bob", Age: 25})
	userCollection.Add(&User{Name: "Charlie", Age: 35})

	// 从聚合对象获取迭代器
	iterator := userCollection.CreateIterator()

	// 使用迭代器遍历集合
	for iterator.HasNext() {
		user := iterator.Next()
		if user != nil {
			t.Logf("User:{Name %s, Age: %d}", user.Name, user.Age)
		}
	}
}
