/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file svc.toolshed.http.go
 * @package response
 * @author Dr.NP <conan.np@gmail.com>
 * @since 07/13/2025
 */

package response

import "time"

type ToolshedHTTPEssential struct {
	ID   string `json:"id" xml:"id" yaml:"id"`
	Name string `json:"name" xml:"name" yaml:"name"`
}

type ToolshedHTTPStaff struct {
	ID           string `json:"id" xml:"id" yaml:"id"`
	Name         string `json:"name" xml:"name" yaml:"name"`
	DisplayNo    string `json:"display_no" xml:"display_no" yaml:"display_no"`
	CardNo       string `json:"card_no" xml:"card_no" yaml:"card_no"`
	TeamID       string `json:"team_id" xml:"team_id" yaml:"team_id"`
	Team         any    `json:"team,omitempty" xml:"team,omitempty" yaml:"team,omitempty"`
	DepartmentID string `json:"department_id" xml:"department_id" yaml:"department_id"`
	Department   any    `json:"department,omitempty" xml:"department,omitempty" yaml:"department,omitempty"`
	JobTitle     string `json:"job_title" xml:"job_title" yaml:"job_title"`
	Telephone    string `json:"telephone" xml:"telephone" yaml:"telephone"`
	Status       int    `json:"status" xml:"status" yaml:"status"`
}

type ToolshedHTTPAccount struct {
	ID       string `json:"id" xml:"id" yaml:"id"`
	Username string `json:"username" xml:"username" yaml:"username"`
	Name     string `json:"name" xml:"name" yaml:"name"`
	GroupID  string `json:"group_id" xml:"group_id" yaml:"group_id"`
	Group    any    `json:"group,omitempty" xml:"group,omitempty" yaml:"group,omitempty"`
}

type ToolshedHTTPTool struct {
	ID             string     `json:"id" xml:"id" yaml:"id"`
	Name           string     `json:"name" xml:"name" yaml:"name"`
	Serial         string     `json:"serial" xml:"serial" yaml:"serial"`
	RFIDTag        string     `json:"rfid_tag" xml:"rfid_tag" yaml:"rfid_tag"`
	Model          string     `json:"model" xml:"model" yaml:"model"`
	Accuracy       string     `json:"accuracy" xml:"accuracy" yaml:"accuracy"`
	Range          string     `json:"range" xml:"range" yaml:"range"`
	Position       string     `json:"position" xml:"position" yaml:"position"`
	Description    string     `json:"description" xml:"description" yaml:"description"`
	Image          string     `json:"image" xml:"image" yaml:"image"`
	CategoryID     string     `json:"category_id" xml:"category_id" yaml:"category_id"`
	Category       any        `json:"category,omitempty" xml:"category,omitempty" yaml:"category,omitempty"`
	ManufacturerID string     `json:"manufacturer_id" xml:"manufacturer_id" yaml:"manufactureer_id"`
	Manufacturer   any        `json:"manufacturer,omitempty" xml:"manufacturer,omitempty" yaml:"manufacturer,omitempty"`
	PlaceID        string     `json:"place_id" xml:"place_id" yaml:"place_id"`
	Place          any        `json:"place,omitempty" xml:"place,omitempty" yaml:"place,omitempty"`
	VerifierID     string     `json:"verifier_id" xml:"verifier_id" yaml:"verifier_id"`
	Verifier       any        `json:"verifier,omitempty" xml:"verifier,omitempty" yaml:"verifier,omitempty"`
	SubmitterID    string     `json:"submitter_id" xml:"submitter_id" yaml:"submitter_id"`
	Submitter      any        `json:"submitter,omitempty" xml:"submitter,omitempty" yaml:"submitter,omitempty"`
	WorkshopID     string     `json:"workshop_id" xml:"workshop_id" yaml:"workshop_id"`
	Workshop       any        `json:"workshop,omitempty" xml:"workshop,omitempty" yaml:"workshop,omitempty"`
	HolderID       string     `json:"holder_id" xml:"holder_id" yaml:"holder_id"`
	Holder         any        `json:"holder,omitempty" xml:"holder,omitempty" yaml:"holder,omitempty"`
	Status         int        `json:"status" xml:"status" yaml:"status"`
	ReleasedAt     string     `json:"released_at" xml:"released_at" yaml:"released_at"`
	ExpiryDate     *time.Time `json:"expiry_date" xml:"expiry_date" yaml:"expiry_date"`
	VerifyDate     *time.Time `json:"verify_date" xml:"verify_date" yaml:"verify_date"`
	VerifyPeriod   int        `json:"verify_period" xml:"verify_period" yaml:"verify_period"`
	VerifyStatus   int        `json:"verify_status" xml:"verify_status" yaml:"verify_status"`
}

type ToolshedHTTPStatisticsSummary struct {
	TotalTools     int64 `json:"total_tools" xml:"total_tools" yaml:"total_tools"`
	TotalStaff     int64 `json:"total_staff" xml:"total_staff" yaml:"total_staff"`
	TotalBorrow    int64 `json:"total_borrow" xml:"total_borrow" yaml:"total_borrow"`
	TotalReturn    int64 `json:"total_return" xml:"total_return" yaml:"total_return"`
	AvailableTools int64 `json:"available_tools" xml:"available_tools" yaml:"available_tools"`
	BorrowedTools  int64 `json:"borrowed_tools" xml:"borrowed_tools" yaml:"borrowed_tools"`
	OverdueTools   int64 `json:"overdue_tools" xml:"overdue_tools" yaml:"overdue_tools"`
}

type ToolshedHTTPStatisticsTool struct {
	Total     int64 `json:"total" xml:"total" yaml:"total"`
	Available int64 `json:"available" xml:"available" yaml:"available"`
	Borrowed  int64 `json:"borrowed" xml:"borrowed" yaml:"borrowed"`
	Overdue   int64 `json:"overdue" xml:"overdue" yaml:"overdue"`
	AboutTo   int64 `json:"about_to" xml:"about_to" yaml:"about_to"`
	Awaiting  int64 `json:"awaiting" xml:"awaiting" yaml:"awaiting"`
}

type ToolshedHTTPStatisticsStaff struct {
	Total         int64 `json:"total" xml:"total" yaml:"total"`
	TotalActive   int64 `json:"total_active" xml:"total_active" yaml:"total_active"`
	TotalInactive int64 `json:"total_inactive" xml:"total_inactive" yaml:"total_inactive"`
}

type ToolshedHTTPStatisticsVerify struct {
	Total int64 `json:"total" xml:"total" yaml:"total"`
	Pass  int64 `json:"pass" xml:"pass" yaml:"pass"`
	Fail  int64 `json:"fail" xml:"fail" yaml:"fail"`
}

type ToolshedHTTPStatisticsUsageItem struct {
	ToolID   string `json:"tool_id" xml:"tool_id" yaml:"tool_id"`
	ToolName string `json:"tool_name" xml:"tool_name" yaml:"tool_name"`
	Count    int64  `json:"count" xml:"count" yaml:"count"`
}

type ToolshedHTTPStatisticsUsage struct {
	Tools []ToolshedHTTPStatisticsUsageItem `json:"items" xml:"items" yaml:"items"`
}

type ToolshedHTTPStatisticsBase struct {
	Key   string `json:"key" xml:"key" yaml:"key"`
	Count int64  `json:"count" xml:"count" yaml:"count"`
}

type ToolshedHTTPPostAuthLogin struct {
	ID        string `json:"id" yaml:"id" xml:"id"`
	Username  string `json:"username" yaml:"username" xml:"username"`
	Name      string `json:"name" yaml:"name" xml:"name"`
	GroupName string `json:"group_name" yaml:"group_name" xml:"group_name"`
	Token     string `json:"token" yaml:"token" xml:"token"`
	Expired   int64  `json:"expired" yaml:"expired" xml:"expired"`
}

type ToolshedHTTPPostBorrow struct {
	StaffID         string   `json:"staff_id" xml:"staff_id" form:"staff_id"`
	BorrowedToolIDs []string `json:"borrowed_tool_ids" xml:"borrowed_tool_ids" form:"borrowed_tool_ids"`
}

type ToolshedHTTPPostRevert struct {
	StaffID         string   `json:"staff_id" xml:"staff_id" form:"staff_id"`
	RevertedToolIDs []string `json:"reverted_tool_ids" xml:"reverted_tool_ids" form:"reverted_tool_ids"`
}

type ToolshedHTTPActivity struct {
	ID        string         `json:"id" xml:"id" yaml:"id"`
	AccountID string         `json:"account_id" xml:"account_id" yaml:"account_id"`
	Account   any            `json:"account,omitempty" xml:"account,omitempty" yaml:"account,omitempty"`
	ToolID    string         `json:"tool_id" xml:"tool_id" yaml:"tool_id"`
	Tool      any            `json:"tool,omitempty" xml:"tool,omitempty" yaml:"tool,omitempty"`
	StaffID   string         `json:"staff_id" xml:"staff_id" yaml:"staff_id"`
	Staff     any            `json:"staff,omitempty" xml:"staff,omitempty" yaml:"staff,omitempty"`
	Action    string         `json:"action" xml:"action" yaml:"action"`
	Detail    map[string]any `json:"detail,omitempty" xml:"detail,omitempty" yaml:"detail,omitempty"`
	CreatedAt string         `json:"created_at" xml:"created_at" yaml:"created_at"`
}

type ToolshedHTTPItemCount struct {
	ID    string `json:"id" xml:"id" yaml:"id"`
	Name  string `json:"name" xml:"name" yaml:"name"`
	Count int64  `json:"count" xml:"count" yaml:"count"`
}

type ToolshedHTTPCategories struct {
	Categories []*ToolshedHTTPItemCount `json:"categories" xml:"categories" yaml:"categories"`
}

type ToolshedHTTPGroup struct {
	ID   string `json:"id" xml:"id" yaml:"id"`
	Name string `json:"name" xml:"name" yaml:"name"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
