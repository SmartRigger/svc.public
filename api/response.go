/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file response.go
 * @package api
 * @author Dr.NP <conan.np@gmail.com>
 * @since 01/15/2025
 */

package api

const (
	CodeOK = 0
	MsgOK  = "OK"

	CodeGeneralError   = -1
	MsgGeneralError    = "General error"
	CodeStorageFailure = 99500001
	MsgStorageFailure  = "Storage failure"
	CodeNetworkFailure = 99500002
	MsgNetworkFailure  = "Network failure"

	CodeInvalidRequest     = 10400001
	MsgInvalidRequest      = "Invalid request"
	CodeAuthenticateFailed = 10401001
	MsgAuthenticateFailed  = "Authenticate failed"
	CodeUnauthorized       = 10401999
	MsgUnauthorized        = "Unauthorized"
)

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
