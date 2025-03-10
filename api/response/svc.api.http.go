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

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
