package chain

import "fmt"

// Manager 是具体处理者，处理 1000 元以下的采购
type Manager struct {
	next Approver
}

// SetNext 设置下一个审批者
func (m *Manager) SetNext(approver Approver) {
	m.next = approver
}

// ProcessRequest 处理采购请求，若金额小于 1000 元则批准，否则传递给下一个审批者
func (m *Manager) ProcessRequest(request *PurchaseRequest) string {
	if request.Amount < 1000 {
		return fmt.Sprintf("经理批准了采购 '%s'，金额：%f", request.Purpose, request.Amount)
	}
	if m.next != nil {
		return m.next.ProcessRequest(request)
	}
	return "无人能处理此采购请求"
}

// Director 是具体处理者，处理 10000 元以下的采购
type Director struct {
	next Approver
}

// SetNext 设置下一个审批者
func (d *Director) SetNext(approver Approver) {
	d.next = approver
}

// ProcessRequest 处理采购请求，若金额小于 10000 元则批准，否则传递给下一个审批者
func (d *Director) ProcessRequest(request *PurchaseRequest) string {
	if request.Amount < 10000 {
		return fmt.Sprintf("总监批准了采购 '%s'，金额：%f", request.Purpose, request.Amount)
	}
	if d.next != nil {
		return d.next.ProcessRequest(request)
	}
	return "无人能处理此采购请求"
}

// CEO 是具体处理者，处理所有金额的采购
type CEO struct {
	next Approver
}

// SetNext 设置下一个审批者
func (c *CEO) SetNext(approver Approver) {
	c.next = approver
}

// ProcessRequest 处理采购请求，CEO 批准所有采购请求
func (c *CEO) ProcessRequest(request *PurchaseRequest) string {
	return fmt.Sprintf("CEO 批准了采购 '%s'，金额：%f", request.Purpose, request.Amount)
}
