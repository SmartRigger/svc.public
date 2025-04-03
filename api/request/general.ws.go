/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file general.ws.go
 * @package request
 * @author Dr.NP <conan.np@gmail.com>
 * @since 04/03/2025
 */

package request

type WSBase struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type WSBindToken struct {
	Type string `json:"type"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type WSPing struct {
	Type string `json:"type"`
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
