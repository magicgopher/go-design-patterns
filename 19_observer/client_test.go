package observer

import "testing"

// 单元测试
// 模拟客户端调用

// TestObserverPattern 模拟真实客户端使用观察者模式的完整流程
func TestObserverPattern(t *testing.T) {
	// 1. 创建被观察者（主题）
	newsAgency := NewConcreteSubject()

	// 2. 创建多个具体观察者（比如不同的新闻订阅者）
	tvStation := NewConcreteObserver("电视台")
	wechatUser := NewConcreteObserver("微信用户")
	appUser := NewConcreteObserver("手机APP用户")

	// 3. 订阅（注册观察者）
	newsAgency.Attach(tvStation)
	newsAgency.Attach(wechatUser)
	newsAgency.Attach(appUser)

	t.Log("第一次发布新闻，所有人都能收到...")
	// 4. 主题状态改变 → 自动通知所有观察者（推模型）
	newsAgency.Notify("紧急：最新天气预报发布！")

	t.Log("\n微信用户取关了...")
	// 5. 某个观察者取消订阅
	newsAgency.Detach(wechatUser)

	t.Log("第二次发布新闻，只有剩余订阅者收到...")
	newsAgency.Notify("股市大涨，指数突破4000点！")
}
