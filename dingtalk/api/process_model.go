package api

// Request Data

type FormValues []FormComponentValuesVo

type CreateProcessInstanceReq struct {
	AgentID          int64  `json:"agent_id"`
	ProcessCode      string `json:"process_code"`
	OriginatorUserID string `json:"originator_user_id"`
	DeptID           int64  `json:"dept_id"`
	// Approvers           string                      `json:"approvers"`
	// ApproversV2         []ProcessInstanceApproverVo `json:"approvers_v2"`
	// CcList              string                      `json:"cc_list"`
	// CcPosition          string                      `json:"cc_position"`
	FormComponentValues FormValues `json:"form_component_values"`
}

type ProcessInstanceApproverVo struct {
	UserIds        []string `json:"user_ids"`
	TaskActionType string   `json:"task_action_type"`
}

type FormComponentValuesVo struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	ExtValue string `json:"ext_value"`
}

type CreateProcessInstanceResp struct {
	BaseResp
	ProcessInstanceID string `json:"process_instance_id" gorm:"column:process_id"`
}

type BaseResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type FormComponentValue struct {
	ComponentType string `json:"component_type"`
	ID            string `json:"id"`
	Value         string `json:"value"`
	Name          string `json:"name,omitempty"`
	ExtValue      string `json:"ext_value,omitempty"`
}

type Task struct {
	TaskStatus string `json:"task_status"`
	CreateTime string `json:"create_time"`
	TaskResult string `json:"task_result"`
	Userid     string `json:"userid"`
	Taskid     string `json:"taskid"`
	URL        string `json:"url"`
}

type OperationRecord struct {
	Date            string `json:"date"`
	OperationType   string `json:"operation_type"`
	OperationResult string `json:"operation_result"`
	Userid          string `json:"userid"`
}

type ProcessInstance struct {
	Result                     string               `json:"result"`
	OriginatorDeptID           string               `json:"originator_dept_id"`
	CreateTime                 string               `json:"create_time"`
	OriginatorUserid           string               `json:"originator_userid"`
	Title                      string               `json:"title"`
	BusinessID                 string               `json:"business_id"`
	BizAction                  string               `json:"biz_action"`
	Status                     string               `json:"status"`
	OperationRecords           []OperationRecord    `json:"operation_records"`
	FormComponentValues        []FormComponentValue `json:"form_component_values"`
	Tasks                      []Task               `json:"tasks"`
	AttachedProcessInstanceIds []interface{}        `json:"attached_process_instance_ids"`
}

type ProcessInstanceDetail struct {
	Errcode        int             `json:"errcode"`
	ProcessInstanc ProcessInstance `json:"process_instance"`
	RequestID      string          `json:"request_id"`
}
