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

// CsDepartment 部门表
type CsDepartment struct {
	DepartmentUUID           string     `gorm:"type:char(32);primaryKey;comment:部门主键"`
	DepartmentCode           string     `gorm:"type:varchar(32);not null;unique;comment:部门编码"`
	DepartmentName           string     `gorm:"type:varchar(64);not null;comment:部门名称"`
	DepartmentOrder          int        `gorm:"type:int;default:100;not null;comment:部门排序"`
	DepartmentEnglishName    *string    `gorm:"type:varchar(128);comment:部门英文名称"`
	DepartmentShortName      *string    `gorm:"type:varchar(32);comment:部门简称"`
	DepartmentAddress        *string    `gorm:"type:varchar(255);comment:部门地址"`
	IsEntity                 bool       `gorm:"type:tinyint(1);default:1;not null;comment:是否实体部门"`
	AdministrativeHead       *string    `gorm:"type:varchar(32);comment:行政负责人"`
	PartyCommitteeHead       *string    `gorm:"type:varchar(32);comment:党委负责人"`
	EstablishmentDate        time.Time  `gorm:"type:date;default:CURRENT_DATE;not null;comment:成立日期"`
	ExpirationDate           *time.Time `gorm:"type:date;comment:失效日期"`
	UnitCategory             string     `gorm:"type:char(32);not null;comment:单位类别"`
	UnitType                 string     `gorm:"type:char(32);not null;comment:单位办别"`
	ParentDepartment         *string    `gorm:"type:char(32);comment:上级部门"`
	AssignedTeachingBuilding *string    `gorm:"type:char(32);comment:分配教学楼"`
	IsTeachingCollege        bool       `gorm:"type:tinyint(1);default:1;not null;comment:开课院系"`
	IsAttendingCollege       bool       `gorm:"type:tinyint(1);default:1;not null;comment:上课院系"`
	FixedPhone               *string    `gorm:"type:varchar(32);comment:固定电话"`
	Remark                   *string    `gorm:"type:text;comment:备注"`
	IsTeachingOffice         bool       `gorm:"type:tinyint(1);default:0;not null;comment:开课教研室"`
	IsEnabled                bool       `gorm:"type:tinyint(1);default:1;not null;comment:是否启用"`
	CreatedAt                time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdatedAt                time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;not null;comment:更新时间"`

	// 外键关联
	AssignedTeachingBuildingRel *CsBuilding     `gorm:"foreignKey:AssignedTeachingBuilding;references:BuildingUUID"`
	ParentDepartmentRel         *CsDepartment   `gorm:"foreignKey:ParentDepartment;references:DepartmentUUID"`
	UnitCategoryRel             *CsUnitCategory `gorm:"foreignKey:UnitCategory;references:UnitCategoryUUID"`
	UnitTypeRel                 *CsUnitType     `gorm:"foreignKey:UnitType;references:UnitTypeUUID"`
}

// TableName 指定表名
func (CsDepartment) TableName() string {
	return "cs_department"
}
