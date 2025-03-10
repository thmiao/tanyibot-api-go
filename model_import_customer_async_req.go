package tanyibot

// ImportCustomerAsyncReq struct for ImportCustomerAsyncReq
type ImportCustomerAsyncReq struct {
	// 任务Id
	RobotCallJobId int64 `json:"robotCallJobId"`
	// 请求批次号
	BatchId string `json:"batchId"`
	// 是否已呼列表去重 默认false
	CallRecordDup   bool             `json:"callRecordDup,omitempty"`
	CustomerPersons []CustomerPerson `json:"customerPersons,omitempty"`
}
