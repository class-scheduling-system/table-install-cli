package entity

import (
	"time"
)

// CsCourseType represents the GORM entity for the `cs_course_type` table.
type CsCourseType struct {
	CourseTypeUUID string    `gorm:"type:char(32);primaryKey;comment:课程类型主键"`                                                    // 课程类型主键
	Name           string    `gorm:"type:varchar(32);uniqueIndex:uk_course_type_name;not null;comment:课程类型名称"`                   // 课程类型名称，唯一索引
	Description    string    `gorm:"type:varchar(255);comment:课程类型描述"`                                                           // 课程类型描述
	CreatedAt      time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`                             // 创建时间
	UpdatedAt      time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"` // 更新时间
}

// TableName specifies the custom table name for the CsCourseType entity.
func (CsCourseType) TableName() string {
	return "cs_course_type"
}
