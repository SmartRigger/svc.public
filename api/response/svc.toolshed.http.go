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
	ID   string `json:"id" xml:"id" yaml:"id"`
	Name string `json:"name" xml:"name" yaml:"name"`
}

type ToolshedHTTPAccount struct{}

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

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
