package entity

import (
	"time"
)

// CsCourseCategory represents the GORM entity for the `cs_course_category` table.
type CsCourseCategory struct {
	CourseCategoryUUID string    `gorm:"type:char(32);primaryKey;comment:课程分类主键"`                                                    // 课程分类主键
	Name               string    `gorm:"type:varchar(32);uniqueIndex:uk_course_category_name;not null;comment:课程分类名称"`               // 课程分类名称，唯一索引
	Description        string    `gorm:"type:varchar(255);comment:课程分类描述"`                                                           // 课程分类描述
	CreatedAt          time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`                             // 创建时间
	UpdatedAt          time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"` // 更新时间
}

// TableName specifies the custom table name for the CsCourseCategory entity.
func (CsCourseCategory) TableName() string {
	return "cs_course_category"
}
