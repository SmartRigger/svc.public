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

import "time"

type ApiHTTPPostAuthLogin struct {
	ID       string `json:"id" yaml:"id" xml:"id"`
	Username string `json:"username" yaml:"username" xml:"username"`
	Name     string `json:"name" yaml:"name" xml:"name"`
	RealName string `json:"real_name" yaml:"real_name" xml:"real_name"`
	Avatar   string `json:"avatar" yaml:"avatar" xml:"avatar"`
	Status   int    `json:"status" yaml:"status" xml:"status"`
	Token    string `json:"token" yaml:"token" xml:"token"`
	Expired  int64  `json:"expired" yaml:"expired" xml:"expired"`
}

type ApiHTTPGetAccountGet struct {
	ID               string `json:"id" yaml:"id" xml:"id"`
	GroupID          string `json:"group_id" yaml:"group_id" xml:"group_id"`
	GroupName        string `json:"group_name" yaml:"group_name" xml:"group_name"`
	RealName         string `json:"real_name" yaml:"real_name" xml:"real_name"`
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
	RealName  string `json:"real_name" yaml:"real_name" xml:"real_name"`
	Name      string `json:"name" yaml:"name" xml:"name"`
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
	TemplateNumber  string `json:"template_number" yaml:"template_number" xml:"template_number"`
	CraftID         string `json:"craft_id" yaml:"craft_id" xml:"craft_id"`
	Craft           any    `json:"craft" yaml:"craft" xml:"craft"`
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
	TemplateNumber  string `json:"template_number" yaml:"template_number" xml:"template_number"`
	CraftID         string `json:"craft_id" yaml:"craft_id" xml:"craft_id"`
	Craft           any    `json:"craft" yaml:"craft" xml:"craft"`
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
	ID           string `json:"id"`
	TemplateID   string `json:"template_id"`
	TemplateName string `json:"template_name"`
	Template     any    `json:"template"`
	CreatorID    string `json:"creator_id"`
	CreatorName  string `json:"creator_name"`
	Creator      any    `json:"creator"`
	//OperatorID   []string                `json:"operator_id"`
	//OperatorName []string                `json:"operator_name"`
	Operators    []*ApiHTTPGetAccountGet `json:"operators"`
	TaskDetail   string                  `json:"task_detail"`
	ReMatch      string                  `json:"re_match"`
	TaskType     string                  `json:"task_type"`
	GroupCode    string                  `json:"group_code"`
	ItemCount    int32                   `json:"item_count"`
	IsComplete   bool                    `json:"is_complete"`
	CompleteTime string                  `json:"complete_time"`
	CreatedAt    string                  `json:"created_at"`
	UpdatedAt    string                  `json:"updated_at"`
	Addition     string                  `json:"addition"`
}

type ApiHTTPGetDetectionTask struct {
	ID           string `json:"id"`
	TemplateID   string `json:"template_id"`
	TemplateName string `json:"template_name"`
	Template     any    `json:"template"`
	CreatorID    string `json:"creator_id"`
	CreatorName  string `json:"creator_name"`
	Creator      any    `json:"creator"`
	//OperatorID   []string `json:"operator_id"`
	//OperatorName []string `json:"operator_name"`
	Operators    []*ApiHTTPGetAccountGet `json:"operators"`
	TaskDetail   string                  `json:"task_detail"`
	ReMatch      string                  `json:"re_match"`
	TaskType     string                  `json:"task_type"`
	GroupCode    string                  `json:"group_code"`
	ItemCount    int32                   `json:"item_count"`
	IsComplete   bool                    `json:"is_complete"`
	CompleteTime string                  `json:"complete_time"`
	CreatedAt    string                  `json:"created_at"`
	Addition     string                  `json:"addition"`
}

type ApiHTTPGetInspection struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	ToolID    string `json:"tool_id"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}

type ApiHTTPGetToolGet struct {
	ID                       string    `json:"id" yaml:"id" xml:"id"`
	Name                     string    `json:"name" yaml:"name" xml:"name"`
	Category                 string    `json:"category" yaml:"category" xml:"category"`
	Range                    string    `json:"range" yaml:"range" xml:"range"`
	InspectionDate           string    `json:"inspection_date" yaml:"inspection_date" xml:"inspection_date"`
	InspectionPersonnel      string    `json:"inspection_personnel" yaml:"inspection_personnel" xml:"inspection_personnel"`
	DailyStartupVerification time.Time `json:"daily_startup_verification" yaml:"daily_startup_verification" xml:"daily_startup_verification"`
	SelfMutualInspection     bool      `json:"self_mutual_inspection" yaml:"self_mutual_inspection" xml:"self_mutual_inspection"`
	MeasurementUnit          string    `json:"measurement_unit" yaml:"measurement_unit" xml:"measurement_unit"`
	RegistrationDate         string    `json:"registration_date" yaml:"registration_date" xml:"registration_date"`
	Detail                   string    `json:"detail" yaml:"detail" xml:"detail"`
	ExpireDate               string    `json:"expire_date" yaml:"expire_date" xml:"expire_date"`
	OwnerID                  string    `json:"owner_id" yaml:"owner_id" xml:"owner_id"`
}

type ApiHTTPGetTool struct {
	ID                       string    `json:"id" yaml:"id" xml:"id"`
	Name                     string    `json:"name" yaml:"name" xml:"name"`
	Category                 string    `json:"category" yaml:"category" xml:"category"`
	Range                    string    `json:"range" yaml:"range" xml:"range"`
	InspectionDate           string    `json:"inspection_date" yaml:"inspection_date" xml:"inspection_date"`
	MeasurementUnit          string    `json:"measurement_unit" yaml:"measurement_unit" xml:"measurement_unit"`
	RegistrationDate         string    `json:"registration_date" yaml:"registration_date" xml:"registration_date"`
	CreatedAt                string    `json:"created_at" yaml:"created_at" xml:"created_at"`
	DailyStartupVerification time.Time `json:"daily_startup_verification" yaml:"daily_startup_verification" xml:"daily_startup_verification"`
	SelfMutualInspection     bool      `json:"self_mutual_inspection" yaml:"self_mutual_inspection" xml:"self_mutual_inspection"`
	InspectionPersonnel      string    `json:"inspection_personnel" yaml:"inspection_personnel" xml:"inspection_personnel"`
	Detail                   string    `json:"detail" yaml:"detail" xml:"detail"`
	ExpireDate               string    `json:"expire_date" yaml:"expire_date" xml:"expire_date"`
	OwnerID                  string    `json:"owner_id" yaml:"owner_id" xml:"owner_id"`
}

type ApiHTTPGetMeasure struct {
	ID        string               `json:"id" yaml:"id" xml:"id"`
	TaskID    string               `json:"task_id" yaml:"task_id" xml:"task_id"`
	ItemID    string               `json:"item_id" yaml:"item_id" xml:"item_id"`
	PointID   string               `json:"point_id" yaml:"point_id" xml:"point_id"`
	Index     string               `json:"index" yaml:"index" xml:"index"`
	IndexNo   string               `json:"index_no" yaml:"index_no" xml:"index_no"`
	SerialNo  string               `json:"serial_no" yaml:"serial_no" xml:"serial_no"`
	Status    int                  `json:"status" yaml:"status" xml:"status"`
	Value     string               `json:"value" yaml:"value" xml:"value"`
	PT        string               `json:"pt" yaml:"pt" xml:"pt"`
	Values    []*ApiHTTPGetDataset `json:"values" yaml:"values" xml:"values"`
	CreatedAt time.Time            `json:"created_at" yaml:"created_at" xml:"created_at"`
}

type ApiHTTPGetDataset struct {
	ID         string    `json:"id" yaml:"id" xml:"id"`
	MeasureID  string    `json:"measure_id" yaml:"measure_id" xml:"measure_id"`
	DeviceID   string    `json:"device_id" yaml:"device_id" xml:"device_id"`
	CreatorID  string    `json:"creator_id" yaml:"creator_id" xml:"creator_id"`
	Creator    any       `json:"creator" yaml:"creator" xml:"creator"`
	Value      string    `json:"value" yaml:"value" xml:"value"`
	Detail     string    `json:"detail" yaml:"detail" xml:"detail"`
	Index      int       `json:"index" yaml:"index" xml:"index"`
	CreateTime time.Time `json:"create_time" yaml:"create_time" xml:"create_time"`
}

type ApiHTTPGetTechnology struct {
	ID            string    `json:"id" yaml:"id" xml:"id"`
	Number        string    `json:"number" yaml:"number" xml:"number"`
	Name          string    `json:"name" yaml:"name" xml:"name"`
	TrainType     string    `json:"train_type" yaml:"train_type"`
	CreatorID     string    `json:"creator_id" yaml:"creator_id" xml:"creator_id"`
	Creator       any       `json:"creator" yaml:"creator" xml:"creator"`
	ItemList      string    `json:"item_list" yaml:"item_list" xml:"item_list"`
	MatchRelation string    `json:"match_relation" yaml:"match_relation" xml:"match_relation"`
	CreatedAt     time.Time `json:"created_at" yaml:"created_at" xml:"created_at"`

	//Version              string    `json:"version" yaml:"version" xml:"version"`
	//Remarks              string    `json:"remarks" yaml:"remarks" xml:"remarks"`
	//Params               string    `json:"params" yaml:"params" xml:"params"`
	//Processes            string    `json:"processes" xml:"processes" yaml:"processes"`
	//MeasurementToolID    string    `json:"measurement_tool_id" yaml:"measurement_tool_id" xml:"measurement_tool_id"`
	//ProductionObjectName string    `json:"production_object_name" yaml:"production_object_name" xml:"production_object_name"`
}

type ApiHTTPGetTechnologyItem struct {
	ID             string            `json:"id" yaml:"id" xml:"id"`
	Name           string            `json:"name" yaml:"name" xml:"name"`
	Type           string            `json:"type" yaml:"type" xml:"type"`
	TechnologyID   string            `json:"technology_id" yaml:"technology_id" xml:"technology_id"`
	TaskID         string            `json:"task_id" yaml:"task_id" xml:"task_id"`
	Description    string            `json:"description" yaml:"description" xml:"description"`
	StandardValue  string            `json:"standard_value" yaml:"standard_value" xml:"standard_value"`
	MaxValue       string            `json:"max_value" yaml:"max_value" xml:"max_value"`
	MinValue       string            `json:"min_value" yaml:"min_value" xml:"min_value"`
	EnableScanCode bool              `json:"enable_scan_code" yaml:"enable_scan_code" xml:"enable_scan_code"`
	ScanCodeRegex  string            `json:"scan_code_regex" yaml:"scan_code_regex" xml:"scan_code_regex"`
	Total          int               `json:"total" yaml:"total" xml:"total"`
	PalletCount    int               `json:"pallet_count" yaml:"pallet_count" xml:"pallet_count"`
	PalletRows     int               `json:"pallet_rows" yaml:"pallet_rows" xml:"pallet_rows"`
	PalletColumns  int               `json:"pallet_columns" yaml:"pallet_columns" xml:"pallet_columns"`
	PointCount     int               `json:"point_count" yaml:"point_count" xml:"point_count"`
	Items          []*ApiHTTPGetItem `json:"items,omitempty" yaml:"items,omitempty" xml:"items,omitempty"`
}

type ApiHTTPGetTask struct {
	ID           string    `json:"id" yaml:"id" xml:"id"`
	TechnologyID string    `json:"technology_id" yaml:"technology_id" xml:"technology_id"`
	Technology   any       `json:"technology" yaml:"technology" xml:"technology"`
	CreatorID    string    `json:"creator_id" yaml:"creator_id" xml:"creator_id"`
	Creator      any       `json:"creator" yaml:"creator" xml:"creator"`
	GroupCode    string    `json:"group_code" yaml:"group_code" xml:"group_code"`
	ItemCount    int       `json:"item_count" yaml:"item_count" xml:"item_count"`
	Detail       string    `json:"detail" yaml:"detail" xml:"detail"`
	ReMatch      string    `json:"re_match" yaml:"re_match" xml:"re_match"`
	Type         string    `json:"type" yaml:"type" xml:"type"`
	IsComplete   bool      `json:"is_complete" yaml:"is_complete" xml:"is_complete"`
	CompleteTime time.Time `json:"complete_time" yaml:"complete_time" xml:"complete_time"`
	Addition     string    `json:"addition" yaml:"addition" xml:"addition"`

	Operators []*ApiHTTPGetAccountGet              `json:"operators" yaml:"operators" xml:"operators"`
	Group     map[string]*ApiHTTPGetTechnologyItem `json:"group" yaml:"group" xml:"group"`
}

type ApiHTTPGetItem struct {
	ID       string `json:"id" yaml:"id" xml:"id"`
	OriginID string `json:"origin_id" yaml:"origin_id" xml:"origin_id"`
	TaskID   string `json:"task_id" yaml:"task_id" xml:"task_id"`
	SerialNo string `json:"serial_no" yaml:"serial_no" xml:"serial_no"`
	Index    int    `json:"index" yaml:"index" xml:"index"`
	Pallet   int    `json:"pallet" yaml:"pallet" xml:"pallet"`
	Row      int    `json:"row" yaml:"row" xml:"row"`
	Column   int    `json:"column" yaml:"column" xml:"column"`

	Points []*ApiHTTPGetPoint `json:"points,omitempty" yaml:"points,omitempty" xml:"points,omitempty"`
}

type ApiHTTPGetPoint struct {
	ID           string `json:"id" yaml:"id" xml:"id"`
	ItemID       string `json:"item_id" yaml:"item_id" xml:"item_id"`
	TaskID       string `json:"task_id" yaml:"task_id" xml:"task_id"`
	Name         string `json:"name" yaml:"name" xml:"name"`
	Index        int    `json:"index" yaml:"index" xml:"index"`
	MeasureTimes int    `json:"measure_times" yaml:"measure_times" xml:"measure_times"`
	MeasureID    string `json:"measure_id" yaml:"measure_id" xml:"measure_id"`
	Measure      any    `json:"measure,omitempty" yaml:"measure,omitempty" xml:"measure,omitempty"`
}

type ApiHTTPGetStat struct {
	Total int64 `json:"total" yaml:"total" xml:"total"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
