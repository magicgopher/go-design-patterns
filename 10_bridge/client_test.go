package bridge

import "testing"

// 单元测试
// 模拟客户端调用

// TestRemoteController 测试使用遥控器控制设备
func TestRemoteController(t *testing.T) {
	// 创建电视设备
	tv := &Television{}
	// 使用遥控器控制电视
	remoteTV := NewRemoteController(tv)
	remoteTV.打开设备()
	remoteTV.关闭设备()

	// 创建空调设备
	ac := &AirConditioner{}
	// 使用遥控器控制空调
	remoteAC := NewRemoteController(ac)
	remoteAC.打开设备()
	remoteAC.关闭设备()
}

// TestMobileApp 测试使用手机应用控制设备
func TestMobileApp(t *testing.T) {
	// 创建电视设备
	tv := &Television{}
	// 使用手机应用控制电视
	appTV := NewMobileApp(tv)
	appTV.打开设备()
	appTV.关闭设备()

	// 创建空调设备
	ac := &AirConditioner{}
	// 使用手机应用控制空调
	appAC := NewMobileApp(ac)
	appAC.打开设备()
	appAC.关闭设备()
}
