package database

import (
	"fmt"
	"frontleaves-table-install-cli/entity"
	"frontleaves-table-install-cli/utils"
)

// InitCourseNatureData 初始化课程性质数据
func (db *DbOperate) InitCourseNatureData(name, description string) {
	var courseNature = entity.CsCourseNature{
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
	var courseProperty = entity.CsCourseProperty{
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
	var courseType = entity.CsCourseType{
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
	var courseCategory = entity.CsCourseCategory{
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
