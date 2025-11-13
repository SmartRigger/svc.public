/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file svc.toolshed.http.go
 * @package request
 * @author Dr.NP <conan.np@gmail.com>
 * @since 07/13/2025
 */

package request

type ToolshedHTTPEssential struct {
	Name string `json:"name" xml:"name" form:"name"`
}

type ToolshedHTTPStaff struct {
	Name         string `json:"name" xml:"name" form:"name"`
	DisplayNo    string `json:"display_no" xml:"display_no" form:"display_no"`
	CardNo       string `json:"card_no" xml:"card_no" form:"card_no"`
	TeamID       string `json:"team_id" xml:"team_id" form:"team_id"`
	DepartmentID string `json:"department_id" xml:"department_id" form:"department_id"`
	JobTitle     string `json:"job_title" xml:"job_title" form:"job_title"`
	Telephone    string `json:"telephone" xml:"telephone" form:"telephone"`
	Status       int    `json:"status" xml:"status" form:"status"`
}

type ToolshedHTTPAccount struct {
	Username string `json:"username" xml:"username" form:"username"`
	Name     string `json:"name" xml:"name" form:"name"`
	Password string `json:"password" xml:"password" form:"password"`
}

type ToolshedHTTPTool struct {
	Name           string `json:"name" xml:"name" form:"name"`
	Serial         string `json:"serial" xml:"serial" form:"serial"`
	RFIDTag        string `json:"rfid_tag" xml:"rfid_tag" form:"rfid_tag"`
	Model          string `json:"model" xml:"model" form:"model"`
	Accuracy       string `json:"accuracy" xml:"accuracy" form:"accuracy"`
	Range          string `json:"range" xml:"range" form:"range"`
	Status         int    `json:"status" xml:"status" form:"status"`
	Position       string `json:"position" xml:"position" form:"position"`
	Description    string `json:"description" xml:"description" form:"description"`
	Image          string `json:"image" xml:"image" form:"image"`
	ReleasedAt     string `json:"released_at" xml:"released_at" form:"released_at"`
	CategoryID     string `json:"category_id" xml:"category_id" form:"category_id"`
	ManufacturerID string `json:"manufacturer_id" xml:"manufacturer_id" form:"manufacturer_id"`
	PlaceID        string `json:"place_id" xml:"place_id" form:"place_id"`
	VerifierID     string `json:"verifier_id" xml:"verifier_id" form:"verifier_id"`
	SubmitterID    string `json:"submitter_id" xml:"submitter_id" form:"submitter_id"`
	WorkshopID     string `json:"workshop_id" xml:"workshop_id" form:"workshop_id"`
	HolderID       string `json:"holder_id" xml:"holder_id" form:"holder_id"`
	ExpiryDate     string `json:"expiry_date" xml:"expiry_date" form:"expiry_date"`
	VerifyDate     string `json:"verify_date" xml:"verify_date" form:"verify_date"`
	VerifyPeriod   int    `json:"verify_period" xml:"verify_period" form:"verify_period"`
}

type ToolshedHTTPPostAuthLogin struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type ToolshedHTTPPostBorrow struct {
	StaffID string   `json:"staff_id" xml:"staff_id" form:"staff_id"`
	ToolIDs []string `json:"tool_ids" xml:"tool_ids" form:"tool_ids"`
}

type ToolshedHTTPPostRevert struct {
	StaffID string   `json:"staff_id" xml:"staff_id" form:"staff_id"`
	ToolIDs []string `json:"tool_ids" xml:"tool_ids" form:"tool_ids"`
}

type ToolshedHTTPPostVerify struct {
	StaffID         string `json:"staff_id" xml:"staff_id" form:"staff_id"`
	ToolID          string `json:"tool_id" xml:"tool_id" form:"tool_id"`
	VerifyDate      string `json:"verify_date" xml:"verify_date" form:"verify_date"`
	NextVerifyDate  string `json:"next_verify_date" xml:"next_verify_date" form:"next_verify_date"`
	Organization    string `json:"organization" xml:"organization" form:"organization"`
	CertificationNo string `json:"certificattion_no" xml:"certificattion_no" form:"certificattion_no"`
	Result          int    `json:"result" xml:"result" form:"result"`
	Remark          string `json:"remark" xml:"remark" form:"remark"`
}

type ToolshedHTTPPostScan struct {
	Type string `json:"type" xml:"type" form:"type"`
	Data string `json:"data" xml:"data" form:"data"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
