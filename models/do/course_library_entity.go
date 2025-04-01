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

// CsCourseLibrary 课程库表
type CsCourseLibrary struct {
	CourseLibraryUUID       string    `gorm:"type:char(32);primaryKey;comment:课程库主键"`
	ID                      string    `gorm:"type:varchar(32);unique;not null;comment:课程编号"`
	Name                    string    `gorm:"type:varchar(32);unique;not null;comment:课程库名称"`
	EnglishName             *string   `gorm:"type:varchar(128);not null;comment:课程英文名称"`
	Category                *string   `gorm:"type:char(32);not null;comment:课程类别"`
	Property                *string   `gorm:"type:char(32);not null;comment:课程属性"`
	Type                    string    `gorm:"type:char(32);not null;comment:课程类型"`
	Nature                  *string   `gorm:"type:char(32);not null;comment:课程性质"`
	Department              string    `gorm:"type:char(32);not null;comment:开课学院"`
	IsEnabled               bool      `gorm:"type:tinyint(1);default:1;not null;comment:是否启用"`
	TotalHours              float64   `gorm:"type:decimal(10,2);default:0.00;not null;comment:总学时"`
	WeekHours               float64   `gorm:"type:decimal(10,2);default:0.00;not null;comment:周学时"`
	TheoryHours             float64   `gorm:"type:decimal(10,2);default:0.00;not null;comment:理论学时"`
	ExperimentHours         float64   `gorm:"type:decimal(10,2);default:0.00;not null;comment:实验学时"`
	PracticeHours           float64   `gorm:"type:decimal(10,2);default:0.00;not null;comment:实践学时"`
	ComputerHours           float64   `gorm:"type:decimal(10,2);default:0.00;not null;comment:上机学时"`
	OtherHours              float64   `gorm:"type:decimal(10,2);default:0.00;not null;comment:其他学时"`
	Credit                  float64   `gorm:"type:decimal(10,2);default:0.00;not null;comment:学分"`
	TheoryClassroomType     *string   `gorm:"type:char(32);comment:理论课教室类型"`
	ExperimentClassroomType *string   `gorm:"type:char(32);comment:实验课教室类型"`
	PracticeClassroomType   *string   `gorm:"type:char(32);comment:实践课教室类型"`
	ComputerClassroomType   *string   `gorm:"type:char(32);comment:上机课教室类型"`
	Description             string    `gorm:"type:text;comment:课程库描述"`
	EditUser                *string   `gorm:"type:char(32);comment:编辑人"`
	CreatedAt               time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdatedAt               time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;not null;comment:更新时间"`

	// 外键关联
	CategoryRel                CsCourseCategory `gorm:"foreignKey:Category;references:CourseCategoryUUID"`
	PropertyRel                CsCourseProperty `gorm:"foreignKey:Property;references:CoursePropertyUUID"`
	TypeRel                    CsCourseType     `gorm:"foreignKey:Type;references:CourseTypeUUID"`
	NatureRel                  CsCourseNature   `gorm:"foreignKey:Nature;references:CourseNatureUUID"`
	DepartmentRel              CsDepartment     `gorm:"foreignKey:Department;references:DepartmentUUID"`
	TheoryClassroomTypeRel     CsClassroomType  `gorm:"foreignKey:TheoryClassroomType;references:ClassTypeUUID"`
	ExperimentClassroomTypeRel CsClassroomType  `gorm:"foreignKey:ExperimentClassroomType;references:ClassTypeUUID"`
	PracticeClassroomTypeRel   CsClassroomType  `gorm:"foreignKey:PracticeClassroomType;references:ClassTypeUUID"`
	ComputerClassroomTypeRel   CsClassroomType  `gorm:"foreignKey:ComputerClassroomType;references:ClassTypeUUID"`
}

// TableName 指定表名
func (CsCourseLibrary) TableName() string {
	return "cs_course_library"
}
