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

// Permission
const (
	PermissionValid               = 1       // 可用
	PermissionDeviceInventory     = 1 << 1  // 设备库存
	PermissionDeviceIssue         = 1 << 2  // 设备领用
	PermissionProductionTask      = 1 << 3  // 生产任务
	PermissionProductionRecord    = 1 << 4  // 生产记录
	PermissionProductionQuality   = 1 << 5  // 生产质量
	PermissionProductionCraft     = 1 << 6  // 生产工艺
	PermissionProductionTemplate  = 1 << 7  // 生产模板
	PermissionStatisticsStatus    = 1 << 8  // 统计状态
	PermissionStatisticsProduce   = 1 << 9  // 统计生产
	PermissionStatisticsDashboard = 1 << 10 // 统计看板
	PermissionStatisticsLog       = 1 << 11 // 统计日志
	PermissionStaffList           = 1 << 12 // 人员列表
	PermissionStaffPermission     = 1 << 13 // 人员权限
	PermissionStaffStation        = 1 << 14 // 人员工位
	PermissionStaffLogin          = 1 << 15 // 人员登录
	PermissionDataBackup          = 1 << 16 // 数据备份
	PermissionDataRestore         = 1 << 17 // 数据恢复
	PermissionPassWord            = 1 << 18 // 密码修改
)

// Pre-defined permission group
var (
	PermissionsAdministrator = PermissionValid |
		PermissionDeviceInventory |
		PermissionDeviceIssue |
		PermissionProductionTask |
		PermissionProductionRecord |
		PermissionProductionQuality |
		PermissionProductionCraft |
		PermissionProductionTemplate |
		PermissionStatisticsStatus |
		PermissionStatisticsProduce |
		PermissionStatisticsDashboard |
		PermissionStatisticsLog |
		PermissionStaffList |
		PermissionStaffPermission |
		PermissionStaffStation |
		PermissionStaffLogin |
		PermissionDataBackup |
		PermissionDataRestore |
		PermissionPassWord
	PermissionsDirector = PermissionValid |
		PermissionDeviceInventory |
		PermissionDeviceIssue |
		PermissionProductionTask |
		PermissionProductionRecord |
		PermissionProductionQuality |
		PermissionProductionCraft |
		PermissionProductionTemplate |
		PermissionStatisticsStatus |
		PermissionStatisticsProduce |
		PermissionStatisticsDashboard |
		PermissionStatisticsLog
	PermissionsWorker = PermissionValid
)

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
