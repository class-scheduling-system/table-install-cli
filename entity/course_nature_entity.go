package entity

import (
	"time"
)

// CsCourseNature represents the GORM entity for the `cs_course_nature` table.
type CsCourseNature struct {
	CourseNatureUUID string    `gorm:"type:char(32);primaryKey;comment:课程性质主键"`                                                    // 课程性质主键
	Name             string    `gorm:"type:varchar(32);uniqueIndex:uk_course_nature_name;not null;comment:课程性质名称"`                 // 课程性质名称，唯一索引
	Description      string    `gorm:"type:varchar(255);comment:课程性质描述"`                                                           // 课程性质描述
	CreatedAt        time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`                             // 创建时间
	UpdatedAt        time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"` // 更新时间
}

// TableName specifies the custom table name for the CsCourseNature entity.
func (CsCourseNature) TableName() string {
	return "cs_course_nature"
}
