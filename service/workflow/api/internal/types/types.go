// Code generated by goctl. DO NOT EDIT.
package types

type PageReq struct {
	PageIndex int64 `form:"pageIndex,default=1"`
	PageSize  int64 `form:"pageSize,default=20"`
}

type Pagination struct {
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
	Count     int64 `json:"count"`
}

type ClassifyListReq struct {
	PageReq
}

type ClassifyListData struct {
	ClassifyId int64  `json:"classifyId"`
	Name       string `json:"name"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	CreateBy   int64  `json:"createBy"`
	UpdateBy   int64  `json:"updateBy"`
}

type ClassifyListResp struct {
	List []ClassifyListData `json:"list"`
	Pagination
}

type ClassifyRetrieveReq struct {
	ClassifyId int64 `path:"classifyId"`
}

type ClassifyRetrieveResp struct {
	ClassifyId int64  `json:"classifyId"`
	Name       string `json:"name"`
}

type ClassifyAddReq struct {
	Name     string `json:"name"`
	CreateBy int64  `json:"createBy,optional"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type ClassifyUpdateReq struct {
	ClassifyId int64  `json:"classifyId"`
	Name       string `json:"name"`
	UpdateBy   int64  `json:"updateBy,optional"`
}

type ClassifyDeleteReq struct {
	ClassifyId int64 `path:"classifyId"`
}

type TemplateListReq struct {
	PageReq
}

type TemplateListData struct {
	TemplateId    int64  `json:"templateId"`
	Name          string `json:"name"`
	FormStructure string `json:"form_structure"`
	Remark        string `json:"remark"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	CreateBy      int64  `json:"createBy"`
	UpdateBy      int64  `json:"updateBy"`
}

type TemplateListResp struct {
	List []TemplateListData `json:"list"`
	Pagination
}

type TemplateRetrieveReq struct {
	TemplateId int64 `path:"templateId"`
}

type TemplateRetrieveResp struct {
	TemplateId    int64  `json:"templateId"`
	Name          string `json:"name"`
	FormStructure string `json:"form_structure"`
	Remark        string `json:"remark"`
}

type TemplateAddReq struct {
	Name          string `json:"name"`
	FormStructure string `json:"form_structure"`
	Remark        string `json:"remark"`
	CreateBy      int64  `json:"createBy,optional"`
	UpdateBy      int64  `json:"updateBy,optional"`
}

type TemplateUpdateReq struct {
	TemplateId    int64  `json:"templateId"`
	Name          string `json:"name"`
	FormStructure string `json:"form_structure"`
	Remark        string `json:"remark"`
	UpdateBy      int64  `json:"updateBy,optional"`
}

type TemplateDeleteReq struct {
	TemplateId int64 `path:"templateId"`
}

type TaskListReq struct {
	PageReq
}

type TaskListData struct {
	TaskId    int64  `json:"taskId"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Content   string `json:"content"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
}

type TaskListResp struct {
	List []TaskListData `json:"list"`
	Pagination
}

type TaskRetrieveReq struct {
	TaskId int64 `path:"taskId"`
}

type TaskRetrieveResp struct {
	TaskId   int64  `json:"taskId"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Remark   string `json:"remark"`
}

type TaskAddReq struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Remark   string `json:"remark"`
	CreateBy int64  `json:"createBy,optional"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type TaskUpdateReq struct {
	TaskId   int64  `json:"taskId"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Remark   string `json:"remark"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type TaskDeleteReq struct {
	TaskId int64 `path:"taskId"`
}

type ProcessListReq struct {
	PageReq
}

type ProcessListData struct {
	ProcessId int64  `json:"processId"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Structure string `json:"structure"`
	Classify  int64  `json:"classify"`
	Template  string `json:"template"`
	Task      string `json:"task"`
	Notice    string `json:"notice"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
}

type ProcessListResp struct {
	List []ProcessListData `json:"list"`
	Pagination
}

type ProcessRetrieveReq struct {
	ProcessId int64 `path:"processId"`
}

type ProcessRetrieveResp struct {
	ProcessId int64  `json:"processId"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Structure string `json:"structure"`
	Classify  int64  `json:"classify"`
	Template  string `json:"template"`
	Task      string `json:"task"`
	Notice    string `json:"notice"`
	Remark    string `json:"remark"`
}

type ProcessAddReq struct {
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Structure string `json:"structure"`
	Classify  int64  `json:"classify"`
	Template  string `json:"template"`
	Task      string `json:"task"`
	Notice    string `json:"notice"`
	Remark    string `json:"remark"`
	CreateBy  int64  `json:"createBy,optional"`
	UpdateBy  int64  `json:"updateBy,optional"`
}

type ProcessUpdateReq struct {
	ProcessId int64  `json:"processId"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Structure string `json:"structure"`
	Classify  int64  `json:"classify"`
	Template  string `json:"template"`
	Task      string `json:"task"`
	Notice    string `json:"notice"`
	Remark    string `json:"remark"`
	UpdateBy  int64  `json:"updateBy,optional"`
}

type ProcessDeleteReq struct {
	ProcessId int64 `path:"processId"`
}

type ProcessClassifyReq struct {
	Name string `form:"name,optional"`
}

type ProcessClassifyData struct {
	ClassifyId int64             `json:"classifyId"`
	Name       string            `json:"name"`
	Process    []ProcessListData `json:"process"`
}

type ProcessClassifyResp struct {
	List []ProcessClassifyData `json:"list"`
}

type TicketInfoReq struct {
	ProcessId int64 `form:"processId,default=0"`
	TicketId  int64 `form:"ticketId,default=0"`
}

type TicketInfoResp struct {
	Process     string `json:"process"`
	Template    string `json:"template"`
	Circulation string `json:"circulation"`
	Nodes       string `json:"nodes"`
	Edges       string `json:"edges"`
	Ticket      string `json:"ticket"`
	FormData    string `json:"form_data"`
}

type TicketAddReq struct {
	ProcessId     int64  `json:"process_id"`
	ClassifyId    int64  `json:"classify_id"`
	ProcessMethod string `json:"process_method"`
	Source        string `json:"source"`
	SourceState   string `json:"source_state"`
	State         string `json:"state"`
	Tasks         string `json:"tasks"`
	Template      string `json:"template"`
	CreateBy      int64  `json:"createBy,optional"`
	UpdateBy      int64  `json:"updateBy,optional"`
}

type TicketListReq struct {
	PageReq
	Category int64 `form:"category"`
}

type TicketListData struct {
	TicketId      int64  `json:"ticket_id"`
	ProcessId     int64  `json:"process_id"`
	ProcessName   string `json:"process_name"`
	StateName     string `json:"state_name"`
	ProcessMethod string `json:"process_method"`
	Principals    string `json:"principals"`
	IsEnd         int64  `json:"is_end"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	CreateBy      int64  `json:"createBy"`
	UpdateBy      int64  `json:"updateBy"`
}

type TicketListResp struct {
	List []TicketListData `json:"list"`
	Pagination
}

type TicketHandleReq struct {
	TicketId       int64  `json:"ticket_id"`
	FlowProperties int64  `json:"flow_properties"`
	Remark         string `json:"remark"`
	UpdateBy       int64  `json:"updateBy,optional"`
}

type TicketDeleteReq struct {
	TicketId int64 `path:"ticketId"`
}

type TicketFinishReq struct {
	TicketId int64 `json:"ticket_id"`
	UpdateBy int64 `json:"updateBy,optional"`
}

type TicketTransferReq struct {
	TicketId int64  `json:"ticket_id"`
	UserId   int64  `json:"user_id"`
	Remark   string `json:"remark"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type TicketUrgeReq struct {
	TicketId int64 `json:"ticket_id"`
	UpdateBy int64 `json:"updateBy,optional"`
}
