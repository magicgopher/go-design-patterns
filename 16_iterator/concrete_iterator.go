package iterator

// 迭代器模式

// UserIterator 是具体迭代器
type UserIterator struct {
	users []*User
	index int
}

// HasNext 检查是否有下一个元素
func (u *UserIterator) HasNext() bool {
	return u.index < len(u.users)
}

// Next 获取下一个元素
func (u *UserIterator) Next() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
