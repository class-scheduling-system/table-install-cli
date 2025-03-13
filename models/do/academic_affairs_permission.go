/*
 * --------------------------------------------------------------------------------
 * Copyright (c) 2022-NOW(至今) 锋楪技术团队
 * Author: 锋楪技术团队 (https://www.frontleaves.com)
 *
 * 本文件包含锋楪技术团队项目的源代码，项目的所有源代码均遵循 MIT 开源许可证协议。
 * --------------------------------------------------------------------------------
 * 许可证声明：
 *
 * 版权所有 (c) 2022-2025 锋楪技术团队。保留所有权利。
 *
 * 本软件是“按原样”提供的，没有任何形式的明示或暗示的保证，包括但不限于
 * 对适销性、特定用途的适用性和非侵权性的暗示保证。在任何情况下，
 * 作者或版权持有人均不承担因软件或软件的使用或其他交易而产生的、
 * 由此引起的或以任何方式与此软件有关的任何索赔、损害或其他责任。
 *
 * 使用本软件即表示您了解此声明并同意其条款。
 *
 * 有关 MIT 许可证的更多信息，请查看项目根目录下的 LICENSE 文件或访问：
 * https://opensource.org/licenses/MIT
 * --------------------------------------------------------------------------------
 * 免责声明：
 *
 * 使用本软件的风险由用户自担。作者或版权持有人在法律允许的最大范围内，
 * 对因使用本软件内容而导致的任何直接或间接的损失不承担任何责任。
 * --------------------------------------------------------------------------------
 */

package do

import (
	"time"
)

// CsAcademicAffairsPermission 教务权限表
type CsAcademicAffairsPermission struct {
	AcademicAffairsPermissionUUID string    `gorm:"column:academic_affairs_permission_uuid;type:char(32);primaryKey;comment:教务权限主键" json:"academic_affairs_permission_uuid"`
	AuthorizedUser                string    `gorm:"column:authorized_user;type:char(32);not null;comment:授权用户" json:"authorized_user"`
	Department                    string    `gorm:"column:department;type:char(32);not null;comment:部门（要求该部门为院系）" json:"department"`
	Type                          int8      `gorm:"column:type;type:tinyint;not null;comment:类型 0:所有权限, 1:教务权限..." json:"type"`
	CreatedAt                     time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt                     time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`

	// 外键关联
	User         *CsUser       `gorm:"foreignKey:AuthorizedUser;references:UserUUID" json:"user,omitempty"`
	CsDepartment *CsDepartment `gorm:"foreignKey:Department;references:DepartmentUUID" json:"department_info,omitempty"`
}

// TableName 指定表名
func (CsAcademicAffairsPermission) TableName() string {
	return "cs_academic_affairs_permission"
}
