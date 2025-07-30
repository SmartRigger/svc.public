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
	Name      string `json:"name" xml:"name" form:"name"`
	GroupID   string `json:"group_id" xml:"group_id" form:"group_id"`
	Status    *int   `json:"status" xml:"status" form:"status"`
	RealName  string `json:"real_name" xml:"real_name" form:"real_name"`
	DisplayID string `json:"display_id" xml:"display_id" form:"display_id"`
}

type ApiHTTPPostCreateAccount struct {
	Username  string `json:"username" xml:"username" form:"username" binding:"required"`
	GroupID   string `json:"group_id" xml:"group_id" form:"group_id" binding:"required"`
	Password  string `json:"password" xml:"password" form:"password"`
	Status    int    `json:"status" xml:"status" form:"status"`
	RealName  string `json:"real_name" xml:"real_name" form:"real_name"`
	Name      string `json:"name" xml:"name" form:"name"`
	DisplayID string `json:"display_id" xml:"display_id" form:"display_id"`
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
	ID              string `json:"id" xml:"id" form:"id"`
	TemplateName    string `json:"template_name" xml:"template_name" form:"template_name"`
	TemplateNumber  string `json:"template_number" xml:"template_number" form:"template_number"`
	CraftID         string `json:"craft_id" xml:"craft_id" form:"craft_id"`
	Version         string `json:"version" xml:"version" form:"version"`
	Remarks         string `json:"remarks" xml:"remarks" form:"remarks"`
	ParameterConfig string `json:"parameter_config" xml:"parameter_config" form:"parameter_config"`
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
	ID string `json:"id" xml:"id" form:"id"`
	//OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password" xml:"new_password"`
}

type ApiHTTPPostCreateDetectionTask struct {
	TemplateID  string   `json:"template_id" binding:"required"`
	CreatorID   string   `json:"creator_id" binding:"required"`
	OperatorIDs []string `json:"operator_ids"`
	TaskDetail  string   `json:"task_detail"`
	ReMatch     string   `json:"re_match"`
	TaskType    string   `json:"task_type"`
	GroupCode   string   `json:"group_code" binding:"required"`
	ItemCount   int32    `json:"item_count"`
	IsComplete  bool     `json:"is_complete"`
}

type ApiHTTPPostEditDetectionTask struct {
	ID          string   `json:"id" binding:"required"`
	TemplateID  string   `json:"template_id"`
	OperatorIDs []string `json:"operator_ids"`
	TaskDetail  string   `json:"task_detail"`
	ReMatch     string   `json:"re_match"`
	TaskType    string   `json:"task_type"`
	GroupCode   string   `json:"group_code"`
	ItemCount   *int32   `json:"item_count"`
	IsComplete  *bool    `json:"is_complete"`
	ReportData  string   `json:"reportData"`
}

type ApiHTTPPostIOTServe struct {
	Type string `json:"type"`
	Data struct {
		OrderNo    string   `json:"orderNo"`
		SrcCode    string   `json:"srcCode"`
		ModelName  string   `json:"modelName"`
		VehicleNo  string   `json:"vehicleNo"`
		TerminalNo []string `json:"terminalNo"`
		ProcessNo  string   `json:"processNo"`
	} `json:"data"`
}

type ApiHTTPPostInspection struct {
	AccountID string `json:"account_id"`
	ToolID    string `json:"tool_id"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}

type ApiHTTPPostToolCreate struct {
	Name                     string `json:"name" xml:"name" form:"name"`
	Category                 string `json:"category" xml:"category" form:"category"`
	Range                    string `json:"range" xml:"range" form:"range"`
	InspectionDate           string `json:"inspection_date" xml:"inspection_date" form:"inspection_date"`
	InspectionPersonnel      string `json:"inspection_personnel" xml:"inspection_personnel" form:"inspection_personnel"`
	DailyStartupVerification bool   `json:"daily_startup_verification" xml:"daily_startup_verification" form:"daily_startup_verification"`
	SelfMutualInspection     bool   `json:"self_mutual_inspection" xml:"self_mutual_inspection" form:"self_mutual_inspection"`
	MeasurementUnit          string `json:"measurement_unit" xml:"measurement_unit" form:"measurement_unit"`
	RegistrationDate         string `json:"registration_date" xml:"registration_date" form:"registration_date"`
}

type ApiHTTPPostToolUpdate struct {
	Name                     string `json:"name" xml:"name" form:"name"`
	Category                 string `json:"category" xml:"category" form:"category"`
	Range                    string `json:"range" xml:"range" form:"range"`
	InspectionDate           string `json:"inspection_date" xml:"inspection_date" form:"inspection_date"`
	InspectionPersonnel      string `json:"inspection_personnel" xml:"inspection_personnel" form:"inspection_personnel"`
	DailyStartupVerification bool   `json:"daily_startup_verification" xml:"daily_startup_verification" form:"daily_startup_verification"`
	SelfMutualInspection     bool   `json:"self_mutual_inspection" xml:"self_mutual_inspection" form:"self_mutual_inspection"`
	MeasurementUnit          string `json:"measurement_unit" xml:"measurement_unit" form:"measurement_unit"`
	RegistrationDate         string `json:"registration_date" xml:"registration_date" form:"registration_date"`
	OwnerID                  string `json:"owner_id" xml:"owner_id" form:"owner_id"`
}

type ApiHTTPGetToolList struct {
	Page  int `json:"page" xml:"page" form:"page"`
	Limit int `json:"limit" xml:"limit" form:"limit"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
