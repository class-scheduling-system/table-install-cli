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

// CsTeacher 教师表
type CsTeacher struct {
	TeacherUUID string    `gorm:"column:teacher_uuid;type:char(32);primaryKey;not null" json:"teacher_uuid"`
	UnitUUID    string    `gorm:"column:unit_uuid;type:char(32);not null;index:idx_teacher_unit" json:"unit_uuid"`
	UserUUID    *string   `gorm:"column:user_uuid;type:char(32)" json:"user_uuid"`
	ID          string    `gorm:"column:id;type:varchar(64);not null;uniqueIndex:uk_teacher_id" json:"id"`
	Name        string    `gorm:"column:name;type:varchar(32);not null;index:idx_teacher_name" json:"name"`
	EnglishName string    `gorm:"column:english_name;type:varchar(256);not null" json:"english_name"`
	Ethnic      string    `gorm:"column:ethnic;type:char(32);not null" json:"ethnic"`
	Sex         bool      `gorm:"column:sex;type:tinyint(1);not null;index:idx_teacher_sex" json:"sex"`
	Type        string    `gorm:"column:type;type:char(32);not null" json:"type"`
	Phone       *string   `gorm:"column:phone;type:varchar(16)" json:"phone"`
	Email       *string   `gorm:"column:email;type:varchar(255)" json:"email"`
	JobTitle    *string   `gorm:"column:job_title;type:varchar(32);index:idx_teacher_job_title" json:"job_title"`
	Description *string   `gorm:"column:desc;type:varchar(255)" json:"desc"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"`

	// 可选：添加关联关系
	Unit        *CsDepartment  `gorm:"foreignKey:UnitUUID;references:DepartmentUUID" json:"-"`
	TeacherType *CsTeacherType `gorm:"foreignKey:Type;references:TeacherTypeUUID" json:"-"`
	User        *CsUser        `gorm:"foreignKey:UserUUID;references:UserUUID" json:"-"`
}

func (CsTeacher) TableName() string {
	return "cs_teacher"
}
