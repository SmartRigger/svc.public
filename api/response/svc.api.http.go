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
}

type ApiHTTPGetAccount struct {
	ID        string `json:"id" yaml:"id" xml:"id"`
	Username  string `json:"username" yaml:"username" xml:"username"`
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

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
