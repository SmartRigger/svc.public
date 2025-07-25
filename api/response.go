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
	CodeInvalidSession     = 10401002
	MsgInvalidSession      = "Invalid session"
	CodeNullSession        = 10401003
	MsgNullSession         = "Session not found"
	CodeUnauthorized       = 10401999
	MsgUnauthorized        = "Unauthorized"
	CodePermissionDenied   = 10403001
	MsgPermissionDenied    = "Permission denied"
	CodeNotFound           = 10404001
	MsgNotFound            = "Not found"
	CodeConflict           = 10409001
	MsgConflict            = "Conflict"
)

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
