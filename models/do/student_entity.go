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

import "time"

// CsStudent 学生表
type CsStudent struct {
	StudentUUID string    `gorm:"column:student_uuid;type:char(32);primaryKey;not null" json:"student_uuid"`
	ID          string    `gorm:"column:id;type:varchar(32);not null;uniqueIndex:uk_student_id" json:"id"`
	Name        string    `gorm:"column:name;type:varchar(32);not null;index:idx_student_name" json:"name"`
	Gender      bool      `gorm:"column:gender;type:tinyint(1);not null" json:"gender"`
	GradeUuid   string    `gorm:"column:grade_uuid;type:varchar(32);not null;index:idx_student_grade" json:"grade"`
	IsGraduated *bool     `gorm:"column:is_graduated;type:tinyint(1);not null;default:0" json:"is_graduate"`
	Department  string    `gorm:"column:department;type:char(32);not null;index:idx_student_department" json:"department"`
	Major       string    `gorm:"column:major;type:char(32);not null;index:idx_student_major" json:"major"`
	Class       *string   `gorm:"column:class;type:char(32);index:idx_student_class" json:"class"`
	UserUUID    *string   `gorm:"column:user_uuid;type:char(32)" json:"user_uuid"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"`

	// 可选：添加关联关系
	AdministrativeClass *CsAdministrativeClass `gorm:"foreignKey:Class;references:AdministrativeClassUUID" json:"-"`
	DepartmentRef       *CsDepartment          `gorm:"foreignKey:Department;references:DepartmentUUID" json:"-"`
	MajorRef            *CsMajor               `gorm:"foreignKey:Major;references:MajorUUID" json:"-"`
	UserRef             *CsUser                `gorm:"foreignKey:UserUUID;references:UserUUID" json:"-"`
}

func (CsStudent) TableName() string {
	return "cs_student"
}
