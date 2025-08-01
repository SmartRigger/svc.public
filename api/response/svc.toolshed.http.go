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
	Team         any    `json:"team" xml:"team" yaml:"team"`
	DepartmentID string `json:"department_id" xml:"department_id" yaml:"department_id"`
	Department   any    `json:"department" xml:"department" yaml:"department"`
	JobTitle     string `json:"job_title" xml:"job_title" yaml:"job_title"`
	Telephone    string `json:"telephone" xml:"telephone" yaml:"telephone"`
	Status       int    `json:"status" xml:"status" yaml:"status"`
}

type ToolshedHTTPAccount struct {
	ID       string `json:"id" xml:"id" yaml:"id"`
	Username string `json:"username" xml:"username" yaml:"username"`
	Name     string `json:"name" xml:"name" yaml:"name"`
}

type ToolshedHTTPTool struct {
	ID             string `json:"id" xml:"id" yaml:"id"`
	Name           string `json:"name" xml:"name" yaml:"name"`
	Serial         string `json:"serial" xml:"serial" yaml:"serial"`
	RFIDTag        string `json:"rfid_tag" xml:"rfid_tag" yaml:"rfid_tag"`
	Model          string `json:"model" xml:"model" yaml:"model"`
	Accuracy       string `json:"accuracy" xml:"accuracy" yaml:"accuracy"`
	Range          string `json:"range" xml:"range" yaml:"range"`
	CategoryID     string `json:"category_id" xml:"category_id" yaml:"category_id"`
	Category       any    `json:"category" xml:"category" yaml:"category"`
	ManufacturerID string `json:"manufacturer_id" xml:"manufacturer_id" yaml:"manufactureer_id"`
	Manufacturer   any    `json:"manufacturer" xml:"manufacturer" yaml:"manufacturer"`
	PlaceID        string `json:"place_id" xml:"place_id" yaml:"place_id"`
	Place          any    `json:"place" xml:"place" yaml:"place"`
	VerifierID     string `json:"verifier_id" xml:"verifier_id" yaml:"verifier_id"`
	Verifier       any    `json:"verifier" xml:"verifier" yaml:"verifier"`
	SubmitterID    string `json:"submitter_id" xml:"submitter_id" yaml:"submitter_id"`
	Submitter      any    `json:"submitter" xml:"submitter" yaml:"submitter"`
	ReleasedAt     string `json:"released_at" xml:"released_at" yaml:"released_at"`
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

type ToolshedHTTPPostAuthLogin struct {
	ID       string `json:"id" yaml:"id" xml:"id"`
	Username string `json:"username" yaml:"username" xml:"username"`
	Name     string `json:"name" yaml:"name" xml:"name"`
	RealName string `json:"real_name" yaml:"real_name" xml:"real_name"`
	Avatar   string `json:"avatar" yaml:"avatar" xml:"avatar"`
	Status   int    `json:"status" yaml:"status" xml:"status"`
	Token    string `json:"token" yaml:"token" xml:"token"`
	Expired  int64  `json:"expired" yaml:"expired" xml:"expired"`
}

type ToolshedHTTPPostBorrow struct {
	StaffID         string   `json:"staff_id" xml:"staff_id" form:"staff_id"`
	BorrowedToolIDs []string `json:"borrowed_tool_ids" xml:"borrowed_tool_ids" form:"borrowed_tool_ids"`
}

type ToolshedHTTPPostRevert struct {
	StaffID         string   `json:"staff_id" xml:"staff_id" form:"staff_id"`
	RevertedToolIDs []string `json:"reverted_tool_ids" xml:"reverted_tool_ids" form:"reverted_tool_ids"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
