package iterator

// 迭代器模式

// UserCollection 是具体聚合
type UserCollection struct {
	users []*User
}

// CreateIterator 创建迭代器
func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

// Add 添加用户
func (u *UserCollection) Add(user *User) {
	u.users = append(u.users, user)
}
