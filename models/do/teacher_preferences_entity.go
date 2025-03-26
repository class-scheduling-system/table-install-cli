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
 * 本软件是"按原样"提供的，没有任何形式的明示或暗示的保证，包括但不限于
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

// CsTeacherPreferences 教师课程时间偏好实体
type CsTeacherPreferences struct {
	PreferenceUUID  string    `gorm:"column:preference_uuid;type:char(32);primaryKey;not null" json:"preference_uuid"`                                   // 教师喜好主键
	TeacherUUID     string    `gorm:"column:teacher_uuid;type:char(32);not null;index:idx_teacher_preferences_teacher" json:"teacher_uuid"`              // 教师主键
	SemesterUUID    string    `gorm:"column:semester_uuid;type:char(32);not null;index:idx_teacher_preferences_semester" json:"semester_uuid"`           // 学期主键
	DayOfWeek       int8      `gorm:"column:day_of_week;type:tinyint;not null" json:"day_of_week"`                                                       // 星期几（1-7）
	TimeSlot        int8      `gorm:"column:time_slot;type:tinyint;not null" json:"time_slot"`                                                           // 第几节课（1-12）
	PreferenceLevel int8      `gorm:"column:preference_level;type:tinyint;not null" json:"preference_level"`                                             // 偏好程度（1：最不期望，2：尽量避免，3：可接受，4：较期望，5：非常期望）
	Reason          string    `gorm:"column:reason;type:varchar(255)" json:"reason"`                                                                     // 偏好原因
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
}

// TableName 设置表名
func (CsTeacherPreferences) TableName() string {
	return "cs_teacher_preferences"
}
