package iterator

// 迭代器模式

// User 结构体
type User struct {
	Name string
	Age  int
}

// Iterator 是迭代器接口
type Iterator interface {
	HasNext() bool
	Next() *User
}

// Aggregate 是聚合接口
type Aggregate interface {
	CreateIterator() Iterator
}
