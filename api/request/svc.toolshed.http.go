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
	Name      string `json:"name" xml:"name" form:"name"`
	DisplayNo string `json:"display_no" xml:"display_no" form:"display_no"`
	CardNo    string `json:"card_no" xml:"card_no" form:"card_no"`
	TeamID    string `json:"team_id" xml:"team_id" form:"team_id"`
	JobTitle  string `json:"job_title" xml:"job_title" form:"job_title"`
	Telephone string `json:"telephone" xml:"telephone" form:"telephone"`
	Status    int    `json:"status" xml:"status" form:"status"`
}

type ToolshedHTTPAccount struct {
	Username string `json:"username" xml:"username" form:"username"`
	Name     string `json:"name" xml:"name" form:"name"`
	Password string `json:"password" xml:"password" form:"password"`
}

type ToolshedHTTPPostAuthLogin struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
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
