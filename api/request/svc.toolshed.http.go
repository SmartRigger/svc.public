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
	Name string `json:"name" xml:"name" form:"name"`
}

type ToolshedHTTPAccount struct{}

type ToolshedHTTPPostAuthLogin struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
