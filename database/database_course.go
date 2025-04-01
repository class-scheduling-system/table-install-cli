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

package database

import (
	"fmt"
	"frontleaves-table-install-cli/models/do"
	"frontleaves-table-install-cli/utils"
	"time"
)

// InitCourseNatureData 初始化课程性质数据
func (db *DbOperate) InitCourseNatureData(name, description string) {
	var courseNature = do.CsCourseNature{
		CourseNatureUUID: utils.GenerateUUIDNoDash(),
		Name:             name,
		Description:      description,
	}
	tx := db.database.Create(&courseNature)
	if tx.Error != nil {
		panic("初始化课程性质数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 课程属性表 [%s-%s] 成功\n", name, description)
	}
}

// InitCoursePropertyData 初始化课程属性数据
func (db *DbOperate) InitCoursePropertyData(name, description string) {
	var courseProperty = do.CsCourseProperty{
		CoursePropertyUUID: utils.GenerateUUIDNoDash(),
		Name:               name,
		Description:        description,
	}
	tx := db.database.Create(&courseProperty)
	if tx.Error != nil {
		panic("初始化课程属性数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 课程属性表 [%s-%s] 成功\n", name, description)
	}
}

// InitCourseTypeData 初始化课程类型数据
func (db *DbOperate) InitCourseTypeData(name, description string) {
	var courseType = do.CsCourseType{
		CourseTypeUUID: utils.GenerateUUIDNoDash(),
		Name:           name,
		Description:    description,
	}
	tx := db.database.Create(&courseType)
	if tx.Error != nil {
		panic("初始化课程类型数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 课程类型表 [%s-%s] 成功\n", name, description)
	}
}

// InitCourseCategoryData 初始化课程类别数据
func (db *DbOperate) InitCourseCategoryData(name, description string) {
	var courseCategory = do.CsCourseCategory{
		CourseCategoryUUID: utils.GenerateUUIDNoDash(),
		Name:               name,
		Description:        description,
	}
	tx := db.database.Create(&courseCategory)
	if tx.Error != nil {
		panic("初始化课程类别数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 课程类别表 [%s-%s] 成功\n", name, description)
	}
}

// CourseLibraryData 初始化课程库数据，生成UUID并设置创建及更新时间，使用数据库事务保存数据。
// 参数course为要存入数据库的课程库结构体。若保存过程中发生错误，则程序将 panic。
func (db *DbOperate) CourseLibraryData(course do.CsCourseLibrary) {
	course.CourseLibraryUUID = utils.GenerateUUIDNoDash()
	// 检查 category 数据
	if course.Category != nil {
		var category = do.CsCourseCategory{}
		db.database.Where("name = ?", course.Category).First(&category)
		course.Category = utils.Ptr(category.CourseCategoryUUID)
	}
	// 检查 property 数据
	if course.Property != nil {
		var property = do.CsCourseProperty{}
		db.database.Where("name = ?", course.Property).First(&property)
		course.Property = utils.Ptr(property.CoursePropertyUUID)
	}
	// 检查 type 数据
	var typeInfo = do.CsCourseType{}
	db.database.Where("name = ?", course.Type).First(&typeInfo)
	course.Type = typeInfo.CourseTypeUUID
	// 检查 nature 数据
	if course.Nature != nil {
		var nature = do.CsCourseNature{}
		db.database.Where("name = ?", course.Nature).First(&nature)
		course.Nature = utils.Ptr(nature.CourseNatureUUID)
	}
	// 检查 department
	var department = do.CsDepartment{}
	db.database.Where("department_name = ?", course.Department).First(&department)
	course.Department = department.DepartmentUUID

	// 检查理论课教室类型
	if course.TheoryClassroomType != nil {
		var theoryClassroomType = do.CsClassroomType{}
		db.database.Where("name = ?", course.TheoryClassroomType).First(&theoryClassroomType)
		course.TheoryClassroomType = utils.Ptr(theoryClassroomType.ClassTypeUUID)
	}

	// 检查实验课教室类型
	if course.ExperimentClassroomType != nil {
		var experimentClassroomType = do.CsClassroomType{}
		db.database.Where("name = ?", course.ExperimentClassroomType).First(&experimentClassroomType)
		course.ExperimentClassroomType = utils.Ptr(experimentClassroomType.ClassTypeUUID)
	}

	// 检查实践课教室类型
	if course.PracticeClassroomType != nil {
		var practiceClassroomType = do.CsClassroomType{}
		db.database.Where("name = ?", course.PracticeClassroomType).First(&practiceClassroomType)
		course.PracticeClassroomType = utils.Ptr(practiceClassroomType.ClassTypeUUID)
	}

	// 检查上机课教室类型
	if course.ComputerClassroomType != nil {
		var computerClassroomType = do.CsClassroomType{}
		db.database.Where("name = ?", course.ComputerClassroomType).First(&computerClassroomType)
		course.ComputerClassroomType = utils.Ptr(computerClassroomType.ClassTypeUUID)
	}

	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()
	tx := db.database.Create(&course)
	if tx.Error != nil {
		panic("初始化课程库数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 课程库[%s] 成功\n", course.Name)
	}
}
