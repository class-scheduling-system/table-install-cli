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

// CsTeacherType 教师类型表
type CsTeacherType struct {
	TeacherTypeUUID string    `gorm:"column:teacher_type_uuid;type:char(32);primaryKey;not null" json:"teacher_type_uuid"`                               // 教师类型主键
	TypeName        string    `gorm:"column:type_name;type:varchar(64);not null;index:idx_type_name" json:"type_name"`                                   // 类型名称
	TypeEnglishName string    `gorm:"column:type_english_name;type:varchar(256);not null" json:"type_english_name"`                                      // 类型英文名称
	TypeDesc        string    `gorm:"column:type_desc;type:varchar(256)" json:"type_desc"`                                                               // 类型描述
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
}

// TableName 设置表名
func (CsTeacherType) TableName() string {
	return "cs_teacher_type"
}
