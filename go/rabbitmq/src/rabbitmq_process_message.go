package main

// MessageProcess 消息处理接口
type MessageProcess interface {
	ProcessDevice() error
	ProcessInspection() error
	ProcessNetworkIP() error
}

// Message 具体从xmdb broker上拿到的消息
type Message struct {
	DataType    string      `json:"dataType"`
	CIClassName string      `json:"-"`
	DataObject  interface{} `json:"dataObject"`
}

// MessageCI 具体从xmdb broker上拿到的消息
type MessageCI struct {
	DataType    string `json:"dataType"`
	CIClassName string `json:"-"`
	DataObject  CIData `json:"dataObject"`
}

// MessageRel 具体从xmdb broker上拿到的消息
type MessageRel struct {
	DataType    string  `json:"dataType"`
	CIClassName string  `json:"-"`
	DataObject  RelData `json:"dataObject"`
}

// CIData CI类型数据
type CIData struct {
	// CI类型显示名称
	CIClassDisplayName string `json:"ciClassDisplayName"`
	// CI类型id
	CIClassID string `json:"ciClassId"`
	// CI类型名称
	CIClassName string `json:"ciClassName"`
	// 被谁创建
	CreateBy string `json:"createBy"`
	// 创建时间
	CreateDate uint64 `json:"createDate"`
	// 数据信息
	DataFieldMap map[string]interface{} `json:"dataFieldMap"`
	// 标志ci是新增、修改、删除
	Flag string `json:"flag"`
	// ID
	ID string `json:"id"`
	// 是否审批类
	IsApproval bool `json:"isApproval"`
	// 最后一次修改人
	LastModifiedBy string `json:"lastModifiedBy"`
	// 最后一次修改时间
	LastModifiedDate uint64 `json:"lastModifiedDate"`
	// 最后一次修改版本ID
	LastVersionID string `json:"lastVersionId"`
	// pk名称
	PkName string `json:"pkName"`
	// pk值
	PkValue string `json:"pkValue"`
	// tag标签（用于搜索）
	PropertyTag string `json:"propertyTag"`
	// 来源
	Source string `json:"source"`
	// 租户id
	TenantID string `json:"tenantID"`
}

// RelData 关系类型数据
type RelData struct {
	// 关系类型名称
	RelTypeName string `json:"relTypeName"`
	// 源ci类显示名称
	SrcCiClassDisplayName string `json:"srcCiClassDisplayName"`
	// 源ci类名称
	SrcCiClassName string `json:"srcCiClassName"`
	// 源ci 主键
	SrcCiPkValue string `json:"srcCiPkValue"`
	// 目标ci类显示名称
	DestCiClassDisplayName string `json:"destCiClassDisplayName"`
	// 目标ci类名称
	DestCiClassName string `json:"destCiClassName"`
	// 目标ci 主键
	DestCiPkValue string `json:"destCiPkValue"`
	// 标志关系数据是新增、修改、删除
	Flag string `json:"flag"`
	// 租户id
	TenantID string `json:"tenantID"`
	// 关系类型id
	RelTypeID string `json:"relTypeID"`
	// 源CI数据id
	SrcCiID string `json:"srcCiID"`
	// 源CI类型id
	SrcClsID string `json:"srcClsID"`
	// 目标CI数据id
	DestCiID string `json:"destCiID"`
	// 目标CI类型id
	DestClsID string `json:"destClsID"`
	// 数据信息
	DataFieldMap map[string]interface{} `json:"dataFieldMap"`
	// 数据来源
	Source string `json:"source"`
	// 上一版本ID
	LastVersionID string `json:"lastVersionID"`
}

// MessagePool 消息池
type MessagePool struct {
	Message chan *Message
}
