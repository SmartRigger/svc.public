/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file svc.api.http.go
 * @package request
 * @author Dr.NP <conan.np@gmail.com>
 * @since 01/14/2025
 */

package request

/* {{{ [svc.api/auth] */
type ApiHTTPPostAuthLogin struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type ApiHTTPPostEditAccount struct {
	ID        string `json:"id" xml:"id" form:"id"`
	Username  string `json:"user_name" xml:"user_name" form:"user_name"`
	GroupID   string `json:"group_id" xml:"group_id" form:"group_id"`
	Status    *int   `json:"status" xml:"status" form:"status"`
	DisplayID string `json:"display_id"`
}

type ApiHTTPPostCreateAccount struct {
	UserName  string `json:"user_name" binding:"required"`
	GroupID   string `json:"group_id" binding:"required"`
	Status    int    `json:"status"`
	RealName  string `json:"real_name"`
	DisplayID string `json:"display_id"`
}

type ApiHTTPPostEditCraft struct {
	ID                   string `json:"id"`
	CraftNumber          string `json:"craft_number"`
	CraftName            string `json:"craft_name"`
	MeasurementToolID    string `json:"measurement_tool_id"`
	ProductionObjectName string `json:"production_object_name"`
	Version              string `json:"version"`
	Remarks              string `json:"remarks"`
	CraftProcesses       string `json:"craft_processes"`
	Params               string `json:"params"`
}

type ApiHTTPPostCreateCraft struct {
	CraftNumber          string `json:"craft_number"`
	CraftName            string `json:"craft_name"`
	MeasurementToolID    string `json:"measurement_tool_id"`
	ProductionObjectName string `json:"production_object_name"`
	Version              string `json:"version"`
	Remarks              string `json:"remarks"`
	CraftProcesses       string `json:"craft_processes"`
	Params               string `json:"params"`
}

type ApiHTTPPostEditGroup struct {
	ID          string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string `json:"name" example:"Updated Group Name"`
	Permissions *int64 `json:"permissions" example:"2"`
	Memo        string `json:"memo" example:"Updated memo"`
}

type ApiHTTPPostCreateGroup struct {
	Name        string `json:"name" example:"New Group"`
	Permissions int64  `json:"permissions" example:"1"`
	Memo        string `json:"memo" example:"Group memo"`
}

type ApiHTTPPostReturnToolRequest struct {
	IssueID string `json:"issue_id"`
}

type ApiHTTPPostIssueToolRequest struct {
	ToolID    string `json:"tool_id"`
	StaffName string `json:"staff_name" example:"John Doe"`
}

type ApiHTTPPutMeasurementRequest struct {
	ID               string `json:"id"`
	TemplateID       string `json:"template_id"`
	WorkgroupNumber  string `json:"workgroup_number"`
	MeasurementCount int    `json:"measurement_count"`
}

type ApiHTTPPostMeasurementRequest struct {
	TemplateID       string `json:"template_id"`
	WorkgroupNumber  string `json:"workgroup_number"`
	MeasurementCount int    `json:"measurement_count"`
}

type ApiHTTPPostTemplateRequest struct {
	ID              string `json:"id"`
	TemplateName    string `json:"template_name"`
	CraftID         string `json:"craft_id"`
	Version         string `json:"version"`
	Remarks         string `json:"remarks"`
	ParameterConfig string `json:"parameter_config"`
}

type ApiHTTPPutToolRequest struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	Category                 string `json:"category"`
	Range                    string `json:"range"`
	InspectionDate           string `json:"inspection_date"`
	InspectionPersonnel      string `json:"inspection_personnel"`
	DailyStartupVerification bool   `json:"daily_startup_verification"`
	SelfMutualInspection     bool   `json:"self_mutual_inspection"`
	MeasurementUnit          string `json:"measurement_unit"`
	ExpireDate               string `json:"expire_date"`
	Detail                   string `json:"detail"`
}

type ApiHTTPPostToolRequest struct {
	Name                     string `json:"name"`
	Category                 string `json:"category"`
	Range                    string `json:"range"`
	InspectionDate           string `json:"inspection_date"`
	InspectionPersonnel      string `json:"inspection_personnel"`
	DailyStartupVerification string `json:"daily_startup_verification"`
	SelfMutualInspection     string `json:"self_mutual_inspection"`
	MeasurementUnit          string `json:"measurement_unit"`
	RegistrationDate         string `json:"registration_date"`
	ExpireDate               string `json:"expire_date"`
	Detail                   string `json:"detail"`
}

type ApiHTTPPostChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ApiHTTPPostChangePasswordByID struct {
	ID          string `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ApiHTTPPostCreateDetectionTask struct {
	TemplateID string `json:"template_id" binding:"required"`
	CreatorID  string `json:"creator_id" binding:"required"`
	OperatorID string `json:"operator_id"`
	TaskDetail string `json:"task_detail"`
	ReMatch    string `json:"re_match"`
	TaskType   string `json:"task_type"`
	GroupCode  string `json:"group_code" binding:"required"`
	ItemCount  int32  `json:"item_count"`
	IsComplete bool   `json:"is_complete"`
}

type ApiHTTPPostEditDetectionTask struct {
	ID         string `json:"id" binding:"required"`
	TemplateID string `json:"template_id"`
	OperatorID string `json:"operator_id"`
	TaskDetail string `json:"task_detail"`
	ReMatch    string `json:"re_match"`
	TaskType   string `json:"task_type"`
	GroupCode  string `json:"group_code"`
	ItemCount  *int32 `json:"item_count"`
	IsComplete *bool  `json:"is_complete"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
