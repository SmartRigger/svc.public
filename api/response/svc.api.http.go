/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file svc.api.http.go
 * @package response
 * @author Dr.NP <conan.np@gmail.com>
 * @since 01/14/2025
 */

package response

type ApiHTTPPostAuthLogin struct {
	ID       string `json:"id" yaml:"id" xml:"id"`
	Username string `json:"username" yaml:"username" xml:"username"`
	Name     string `json:"name" yaml:"name" xml:"name"`
	Avatar   string `json:"avatar" yaml:"avatar" xml:"avatar"`
	Status   int    `json:"status" yaml:"status" xml:"status"`
	Token    string `json:"token" yaml:"token" xml:"token"`
	Expired  int64  `json:"expired" yaml:"expired" xml:"expired"`
}

type ApiHTTPGetAccountGet struct {
	ID               string `json:"id" yaml:"id" xml:"id"`
	GroupID          string `json:"group_id" yaml:"group_id" xml:"group_id"`
	GroupName        string `json:"group_name" yaml:"group_name" xml:"group_name"`
	Realname         string `json:"realname" yaml:"realname" xml:"realname"`
	GroupPermissions int64  `json:"group_permissions" yaml:"group_permissions" xml:"group_permissions"`
	Username         string `json:"username" yaml:"username" xml:"username"`
	Name             string `json:"name" yaml:"name" xml:"name"`
	DisplayID        string `json:"display_id" yaml:"display_id" xml:"display_id"`
	Status           int    `json:"status" yaml:"status" xml:"status"`
	Avatar           string `json:"avatar" yaml:"avatar" xml:"avatar"`
}

type ApiHTTPGetGroupGet struct {
	ID          string `json:"id" yaml:"id" xml:"id"`
	Name        string `json:"name" yaml:"name" xml:"name"`
	Permissions int64  `json:"permissions" yaml:"permissions" xml:"permissions"`
	Memo        string `json:"memo" yaml:"memo" xml:"memo"`
}

type ApiHTTPGetAccount struct {
	ID        string `json:"id" yaml:"id" xml:"id"`
	Username  string `json:"username" yaml:"username" xml:"username"`
	Realname  string `json:"realname" yaml:"realname" xml:"realname"`
	DisplayID string `json:"display_id" yaml:"display_id" xml:"display_id"`
	GroupID   string `json:"group_id" yaml:"group_id" xml:"group_id"`
	GroupName string `json:"group_name" yaml:"group_name" xml:"group_name"`
	Status    int    `json:"status" yaml:"status" xml:"status"`
	CreatedAt string `json:"created_at" yaml:"created_at" xml:"created_at"`
}

type ApiHTTPGetPermission struct {
	Name        string `json:"name" yaml:"name" xml:"name"`
	Value       int    `json:"value" yaml:"value" xml:"value"`
	Description string `json:"description" yaml:"description" xml:"description"`
}

type ApiHTTPGetGroupedPermission struct {
	Group       string                 `json:"group" yaml:"group" xml:"group"`
	Permissions []ApiHTTPGetPermission `json:"permissions" yaml:"permissions" xml:"permissions"`
}

type ApiHTTPGetIssueGet struct {
	ID                  string `json:"id" yaml:"id" xml:"id"`
	ToolID              string `json:"tool_id" yaml:"tool_id" xml:"tool_id"`
	ToolCustomNumber    string `json:"tool_custom_number" yaml:"tool_custom_number" xml:"tool_custom_number"`
	ToolType            string `json:"tool_type" yaml:"tool_type" xml:"tool_type"`
	ToolName            string `json:"tool_name" yaml:"tool_name" xml:"tool_name"`
	ToolSpecification   string `json:"tool_specification" yaml:"tool_specification" xml:"tool_specification"`
	ToolMeasurementUnit string `json:"tool_measurement_unit" yaml:"tool_measurement_unit" xml:"tool_measurement_unit"`
	StaffId             string `json:"staff_id" yaml:"staff_id" xml:"staff_id"`
	StaffName           string `json:"staff_name" yaml:"staff_name" xml:"staff_name"`
	IssueTime           string `json:"issue_time" yaml:"issue_time" xml:"issue_time"`
	ReturnTime          string `json:"return_time" yaml:"return_time" xml:"return_time"`
}

type ApiHTTPGetIssue struct {
	ID                  string `json:"id" yaml:"id" xml:"id"`
	ToolID              string `json:"tool_id" yaml:"tool_id" xml:"tool_id"`
	ToolCustomNumber    string `json:"tool_custom_number" yaml:"tool_custom_number" xml:"tool_custom_number"`
	ToolType            string `json:"tool_type" yaml:"tool_type" xml:"tool_type"`
	ToolName            string `json:"tool_name" yaml:"tool_name" xml:"tool_name"`
	ToolSpecification   string `json:"tool_specification" yaml:"tool_specification" xml:"tool_specification"`
	ToolMeasurementUnit string `json:"tool_measurement_unit" yaml:"tool_measurement_unit" xml:"tool_measurement_unit"`
	StaffId             string `json:"staff_id" yaml:"staff_id" xml:"staff_id"`
	StaffName           string `json:"staff_name" yaml:"staff_name" xml:"staff_name"`
	IssueTime           string `json:"issue_time" yaml:"issue_time" xml:"issue_time"`
	ReturnTime          string `json:"return_time" yaml:"return_time" xml:"return_time"`
}

type ApiHTTPGetCraftGet struct {
	ID                   string `json:"id" yaml:"id" xml:"id"`
	CraftNumber          string `json:"craft_number" yaml:"craft_number" xml:"craft_number"`
	CraftName            string `json:"craft_name" yaml:"craft_name" xml:"craft_name"`
	MeasurementToolID    string `json:"measurement_tool_id" yaml:"measurement_tool_id" xml:"measurement_tool_id"`
	ProductionObjectName string `json:"production_object_name" yaml:"production_object_name" xml:"production_object_name"`
	CompilerID           string `json:"compiler_id" yaml:"compiler_id" xml:"compiler_id"`
	CompilerName         string `json:"compiler_name" yaml:"compiler_name" xml:"compiler_name"`
	Version              string `json:"version" yaml:"version" xml:"version"`
	Remarks              string `json:"remarks" yaml:"remarks" xml:"remarks"`
	CraftProcesses       string `json:"craft_processes" yaml:"craft_processes" xml:"craft_processes"`
	CreatedAt            string `json:"created_at" yaml:"created_at" xml:"created_at"`
	Params               string `json:"params" yaml:"params" xml:"params"`
}

type ApiHTTPGetCraft struct {
	ID                   string `json:"id" yaml:"id" xml:"id"`
	CraftNumber          string `json:"craft_number" yaml:"craft_number" xml:"craft_number"`
	CraftName            string `json:"craft_name" yaml:"craft_name" xml:"craft_name"`
	MeasurementToolID    string `json:"measurement_tool_id" yaml:"measurement_tool_id" xml:"measurement_tool_id"`
	ProductionObjectName string `json:"production_object_name" yaml:"production_object_name" xml:"production_object_name"`
	CraftProcesses       string `bun:"craft_processes" json:"craft_processes" xml:"craft_processes" yaml:"craft_processes"`
	CompilerID           string `json:"compiler_id" yaml:"compiler_id" xml:"compiler_id"`
	CompilerName         string `json:"compiler_name" yaml:"compiler_name" xml:"compiler_name"`
	Version              string `json:"version" yaml:"version" xml:"version"`
	Remarks              string `json:"remarks" yaml:"remarks" xml:"remarks"`
	CreatedAt            string `json:"created_at" yaml:"created_at" xml:"created_at"`
	Params               string `json:"params" yaml:"params" xml:"params"`
}

type ApiHTTPGetTemplateGet struct {
	ID              string `json:"id" yaml:"id" xml:"id"`
	TemplateName    string `json:"template_name" yaml:"template_name" xml:"template_name"`
	CraftID         string `json:"craft_id" yaml:"craft_id" xml:"craft_id"`
	CompilerID      string `json:"compiler_id" yaml:"compiler_id" xml:"compiler_id"`
	CompilerName    string `json:"compiler_name" yaml:"compiler_name" xml:"compiler_name"`
	Version         string `json:"version" yaml:"version" xml:"version"`
	Remarks         string `json:"remarks" yaml:"remarks" xml:"remarks"`
	ParameterConfig string `json:"parameter_config" yaml:"parameter_config" xml:"parameter_config"`
	CreatedAt       string `json:"created_at" yaml:"created_at" xml:"created_at"`
}

type ApiHTTPGetTemplate struct {
	ID              string `json:"id" yaml:"id" xml:"id"`
	TemplateName    string `json:"template_name" yaml:"template_name" xml:"template_name"`
	CraftID         string `json:"craft_id" yaml:"craft_id" xml:"craft_id"`
	CompilerID      string `json:"compiler_id" yaml:"compiler_id" xml:"compiler_id"`
	CompilerName    string `json:"compiler_name" yaml:"compiler_name" xml:"compiler_name"`
	Version         string `json:"version" yaml:"version" xml:"version"`
	Remarks         string `json:"remarks" yaml:"remarks" xml:"remarks"`
	ParameterConfig string `json:"parameter_config" yaml:"parameter_config" xml:"parameter_config"`
	CreatedAt       string `json:"created_at" yaml:"created_at" xml:"created_at"`
}

type ApiHTTPGetMeasurementGet struct {
	ID               string `json:"id" yaml:"id" xml:"id"`
	TemplateID       string `json:"template_id" yaml:"template_id" xml:"template_id"`
	TemplateName     string `json:"template_name" yaml:"template_name" xml:"template_name"`
	WorkgroupNumber  string `json:"workgroup_number" yaml:"workgroup_number" xml:"workgroup_number"`
	MeasurementCount int16  `json:"measurement_count" yaml:"measurement_count" xml:"measurement_count"`
	CreatorID        string `json:"creator_id" yaml:"creator_id" xml:"creator_id"`
	CreatorName      string `json:"creator_name" yaml:"creator_name" xml:"creator_name"`
	CreatedAt        string `json:"created_at" yaml:"created_at" xml:"created_at"`
	Number           string `json:"number" yaml:"number" xml:"number"`
}

type ApiHTTPGetMeasurement struct {
	SerialNumber     int64  `json:"serial_number" yaml:"serial_number" xml:"serial_number"`
	ID               string `json:"id" yaml:"id" xml:"id"`
	TemplateName     string `json:"template_name" yaml:"template_name" xml:"template_name"`
	Number           string `json:"number" yaml:"number" xml:"number"`
	WorkgroupNumber  string `json:"workgroup_number" yaml:"workgroup_number" xml:"workgroup_number"`
	MeasurementCount int16  `json:"measurement_count" yaml:"measurement_count" xml:"measurement_count"`
	CreatorName      string `json:"creator_name" yaml:"creator_name" xml:"creator_name"`
	CreatedAt        string `json:"created_at" yaml:"created_at" xml:"created_at"`
}

type ApiHTTPGetLog struct {
	SerialNumber int64  `json:"serial_number"`
	ID           string `json:"id"`
	SessionID    string `json:"session_id"`
	Account      string `json:"account"`
	AccountName  string `json:"account_name"`
	Operation    int    `json:"operation"`
	Success      bool   `json:"success"`
	Error        string `json:"error,omitempty"`
	Detail       string `json:"detail"`
	CreatedAt    string `json:"created_at"`
}

type ApiHTTPGetToolCategory struct {
	Category string `json:"category"`
	Count    int64  `json:"count"`
}

type ApiHTTPGetToolUsage struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	ImageURL   string `json:"image_url"`
	UsageCount int64  `json:"usage_count"`
}

type ApiHTTPGetDetectionTaskGet struct {
	ID           string      `json:"id"`
	TemplateID   string      `json:"template_id"`
	TemplateName string      `json:"template_name"`
	Template     interface{} `json:"template"`
	CreatorID    string      `json:"creator_id"`
	CreatorName  string      `json:"creator_name"`
	Creator      interface{} `json:"creator"`
	OperatorID   string      `json:"operator_id"`
	OperatorName string      `json:"operator_name"`
	Operator     interface{} `json:"operator"`
	TaskDetail   string      `json:"task_detail"`
	ReMatch      string      `json:"re_match"`
	TaskType     string      `json:"task_type"`
	GroupCode    string      `json:"group_code"`
	ItemCount    int32       `json:"item_count"`
	IsComplete   bool        `json:"is_complete"`
	CompleteTime string      `json:"complete_time"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	Addition     string      `json:"addition"`
}

type ApiHTTPGetDetectionTask struct {
	ID           string      `json:"id"`
	TemplateID   string      `json:"template_id"`
	TemplateName string      `json:"template_name"`
	Template     interface{} `json:"template"`
	CreatorID    string      `json:"creator_id"`
	CreatorName  string      `json:"creator_name"`
	Creator      interface{} `json:"creator"`
	OperatorID   string      `json:"operator_id"`
	OperatorName string      `json:"operator_name"`
	Operator     interface{} `json:"operator"`
	TaskDetail   string      `json:"task_detail"`
	ReMatch      string      `json:"re_match"`
	TaskType     string      `json:"task_type"`
	GroupCode    string      `json:"group_code"`
	ItemCount    int32       `json:"item_count"`
	IsComplete   bool        `json:"is_complete"`
	CompleteTime string      `json:"complete_time"`
	CreatedAt    string      `json:"created_at"`
	Addition     string      `json:"addition"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
