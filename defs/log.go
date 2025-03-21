/*
 * Copyright (C) HereweTech, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file permission.go
 * @package defs
 * @author Dr.NP <conan.np@gmail.com>
 * @since 01/04/2025
 */

package defs

type LogOperation int

const (
	OpEmpty LogOperation = iota
	OpLogin
	OpLogout
	OpAccountCreate
	OpAccountEdit
	OpAccountDelete
	OpPasswordChange
	OpPasswordChangeByAdmin
	OpToolCreate
	OpToolEdit
	OpToolDelete
	OpCraftCreate
	OpCraftEdit
	OpCraftDelete
	OpMeasurementCreate
	OpMeasurementEdit
	OpMeasurementDelete
	OpTemplateCreate
	OpTemplateEdit
	OpTemplateDelete
	OpIssueCreate
	OpIssueEdit
	OpIssueDelete
	OpGroupCreate
	OpGroupEdit
	OpGroupDelete
)

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
