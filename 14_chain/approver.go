package chain

// PurchaseRequest 结构体定义了采购请求的内容
type PurchaseRequest struct {
	Amount  float64
	Purpose string
}

// Approver 是处理者接口，定义了审批行为
type Approver interface {
	SetNext(approver Approver)
	ProcessRequest(request *PurchaseRequest) string
}
