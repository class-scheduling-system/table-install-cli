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

// CsAdministrativeClass 行政班实体
type CsAdministrativeClass struct {
	AdministrativeClassUUID string    `gorm:"column:administrative_class_uuid;type:char(32);primaryKey;not null" json:"administrative_class_uuid"`               // 行政班主键
	DepartmentUUID          string    `gorm:"column:department_uuid;type:char(32);not null" json:"department_uuid"`                                              // 所属部门/院系
	MajorUUID               string    `gorm:"column:major_uuid;type:char(32);not null" json:"major_uuid"`                                                        // 所属专业
	ClassCode               string    `gorm:"column:class_code;type:varchar(32);not null;uniqueIndex:uk_class_code" json:"class_code"`                           // 班级编号
	ClassName               string    `gorm:"column:class_name;type:varchar(64);not null" json:"class_name"`                                                     // 班级名称
	GradeUUID               string    `gorm:"column:grade_uuid;type:varchar(16);not null;index:idx_administrative_class_grade" json:"grade"`                     // 年级
	StudentCount            uint      `gorm:"column:student_count;type:int unsigned;not null;default:0" json:"student_count"`                                    // 学生人数
	CounselorUUID           *string   `gorm:"column:counselor_uuid;type:char(32);index:idx_administrative_class_counselor" json:"counselor_uuid"`                // 辅导员UUID
	MonitorUUID             *string   `gorm:"column:monitor_uuid;type:char(32);index:idx_administrative_class_monitor" json:"monitor_uuid"`                      // 班长UUID
	IsEnabled               bool      `gorm:"column:is_enabled;type:tinyint(1);not null;default:1;index:idx_administrative_class_enabled" json:"is_enabled"`     // 是否启用(0:禁用,1:启用)
	Description             *string   `gorm:"column:description;type:varchar(255)" json:"description"`                                                           // 班级描述
	CreatedAt               time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt               time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间

	Department *CsDepartment `gorm:"foreignKey:DepartmentUUID" json:"-"`
	Major      *CsMajor      `gorm:"foreignKey:MajorUUID" json:"-"`
	Counselor  *CsTeacher    `gorm:"foreignKey:CounselorUUID" json:"-"`
	Monitor    *CsStudent    `gorm:"foreignKey:MonitorUUID" json:"-"`
}

// TableName 设置表名
func (CsAdministrativeClass) TableName() string {
	return "cs_administrative_class"
}
