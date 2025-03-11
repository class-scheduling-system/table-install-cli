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

// CsSchedulingConflict 排课冲突实体
type CsSchedulingConflict struct {
	ConflictUUID         string     `gorm:"column:conflict_uuid;type:char(32);primaryKey;not null" json:"conflict_uuid"`                                                                              // 冲突主键
	SemesterUUID         string     `gorm:"column:semester_uuid;type:char(32);not null;index:idx_conflict_semester" json:"semester_uuid"`                                                             // 学期主键
	FirstAssignmentUUID  string     `gorm:"column:first_assignment_uuid;type:char(32);not null;index:idx_conflict_first_assignment,idx_conflict_assignments" json:"first_assignment_uuid"`            // 第一个排课主键
	SecondAssignmentUUID string     `gorm:"column:second_assignment_uuid;type:char(32);not null;index:idx_conflict_second_assignment,idx_conflict_assignments" json:"second_assignment_uuid"`         // 第二个排课主键
	ConflictType         uint8      `gorm:"column:conflict_type;type:tinyint unsigned;not null;index:idx_conflict_type,idx_conflict_status_type" json:"conflict_type"`                                // 冲突类型: 1-教师冲突 2-教室冲突 3-班级冲突 4-其他冲突
	ConflictTime         string     `gorm:"column:conflict_time;type:json;not null" json:"conflict_time"`                                                                                             // 冲突时间
	Description          string     `gorm:"column:description;type:varchar(255)" json:"description"`                                                                                                  // 冲突描述
	ResolutionStatus     uint8      `gorm:"column:resolution_status;type:tinyint unsigned;not null;default:0;index:idx_conflict_resolution_status,idx_conflict_status_type" json:"resolution_status"` // 解决状态: 0-未解决 1-已解决 2-忽略
	ResolutionMethod     *uint8     `gorm:"column:resolution_method;type:tinyint unsigned" json:"resolution_method"`                                                                                  // 解决方法: 1-调整第一个课程 2-调整第二个课程 3-同时调整 4-其他
	ResolutionNotes      string     `gorm:"column:resolution_notes;type:varchar(255)" json:"resolution_notes"`                                                                                        // 解决备注
	ResolvedBy           string     `gorm:"column:resolved_by;type:char(32);index:idx_conflict_resolved_by" json:"resolved_by"`                                                                       // 解决人
	ResolvedAt           *time.Time `gorm:"column:resolved_at;type:timestamp" json:"resolved_at"`                                                                                                     // 解决时间
	CreatedAt            time.Time  `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                                                                    // 创建时间
	UpdatedAt            time.Time  `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"`                                        // 更新时间

	// 外键关联
	Semester         CsSemester            `gorm:"foreignKey:SemesterUUID;references:SemesterUUID;constraint:OnUpdate:CASCADE" json:"semester"`                                          // 关联学期
	FirstAssignment  CsAdministrativeClass `gorm:"foreignKey:FirstAssignmentUUID;references:ClassAssignmentUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"first_assignment"`   // 关联第一个排课
	SecondAssignment CsAdministrativeClass `gorm:"foreignKey:SecondAssignmentUUID;references:ClassAssignmentUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"second_assignment"` // 关联第二个排课
	Resolver         CsUser                `gorm:"foreignKey:ResolvedBy;references:UserUUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"resolver"`                              // 关联解决人
}

// TableName 设置表名
func (CsSchedulingConflict) TableName() string {
	return "cs_scheduling_conflict"
}
