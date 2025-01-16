package entity

import (
	"time"
)

// CsSystem represents the GORM entity for the `cs_system` table.
type CsSystem struct {
	SystemUUID string    `gorm:"type:char(32);primaryKey;comment:系统主键"`                                                      // 系统主键
	SystemKey  string    `gorm:"type:varchar(32);uniqueIndex:uk_system_key;not null;comment:系统键"`                            // 系统键，唯一索引
	SystemVal  string    `gorm:"type:text;comment:系统值"`                                                                      // 系统值
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`                             // 创建时间
	UpdatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"` // 更新时间
}

// TableName specifies the custom table name for the CsSystem entity.
func (CsSystem) TableName() string {
	return "cs_system"
}
