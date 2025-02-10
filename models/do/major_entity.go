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

// CsMajor 专业表
type CsMajor struct {
	MajorUUID        string    `gorm:"type:char(32);primaryKey;comment:专业主键"`
	MajorName        string    `gorm:"type:varchar(32);not null;unique;comment:专业名称"`
	MajorDescription *string   `gorm:"type:varchar(255);comment:专业描述"`
	MajorCode        string    `gorm:"type:varchar(32);not null;unique;comment:专业代码"`
	MajorStatus      bool      `gorm:"type:tinyint(1);default:1;not null;comment:专业状态 0:禁用 1:启用"`
	DepartmentUUID   string    `gorm:"type:char(32);not null;comment:学院外键"`
	EducationYears   uint16    `gorm:"type:smallint unsigned;not null;comment:学制（年）"`
	TrainingLevel    string    `gorm:"type:varchar(32);not null;comment:培养层次（例如：本科，专科）"`
	CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;not null;comment:更新时间"`

	// 外键关联
	DepartmentRel *CsDepartment `gorm:"foreignKey:DepartmentUUID;references:DepartmentUUID;onUpdate:CASCADE"`
}

// TableName 指定表名
func (CsMajor) TableName() string {
	return "cs_major"
}
