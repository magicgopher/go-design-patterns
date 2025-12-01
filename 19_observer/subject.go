package observer

// 观察者模式（主题接口与具体实现）

// Subject 定义主题（被观察者）接口
type Subject interface {
	// Attach 注册/添加一个观察者
	Attach(observer Observer)

	// Detach 注销/移除一个观察者
	Detach(observer Observer)

	// Notify 通知所有已注册的观察者状态已变更
	Notify(message string)
}

// ConcreteSubject 具体的主题实现（被观察的对象）
type ConcreteSubject struct {
	observers []Observer // 保存所有注册的观察者
	state     string     // 主题当前的状态（可选，用于业务扩展）
}

// NewConcreteSubject 创建一个具体的主题实例
func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
	}
}

// Attach 添加观察者到列表中
func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Detach 从列表中移除指定的观察者
func (s *ConcreteSubject) Detach(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer { // 引用相同即认为相等
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

// Notify 遍历所有观察者，调用它们的 Update 方法进行通知
func (s *ConcreteSubject) Notify(message string) {
	s.state = message // 更新主题状态（实际项目中可能有更多字段）
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// GetState （可选）获取当前状态，用于观察者主动拉取（拉模型扩展）
func (s *ConcreteSubject) GetState() string {
	return s.state
}
