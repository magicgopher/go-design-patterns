package bridge

// 控制器接口（Abstraction），定义控制设备的操作
type Controller interface {
	打开设备()
	关闭设备()
}

// 遥控器（Refined Abstraction），扩展抽象部分
type RemoteController struct {
	device Device
}

// NewRemoteController 创建遥控器实例
func NewRemoteController(device Device) *RemoteController {
	return &RemoteController{device: device}
}

// 打开设备 调用设备的打开方法
func (r *RemoteController) 打开设备() {
	r.device.打开()
}

// 关闭设备 调用设备的关闭方法
func (r *RemoteController) 关闭设备() {
	r.device.关闭()
}

// 手机应用（Refined Abstraction），扩展抽象部分
type MobileApp struct {
	device Device
}

// NewMobileApp 创建手机应用实例
func NewMobileApp(device Device) *MobileApp {
	return &MobileApp{device: device}
}

// 打开设备 调用设备的打开方法
func (m *MobileApp) 打开设备() {
	m.device.打开()
}

// 关闭设备 调用设备的关闭方法
func (m *MobileApp) 关闭设备() {
	m.device.关闭()
}
