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

package database

import (
	"fmt"
	"frontleaves-table-install-cli/models/do"
	"frontleaves-table-install-cli/utils"
)

// InitTeacherTypeData 初始化教师类型数据
func (db *DbOperate) InitTeacherTypeData(name, englishName, description string) {
	var teacherType = do.CsTeacherType{
		TeacherTypeUUID: utils.GenerateUUIDNoDash(),
		TypeName:        name,
		TypeEnglishName: englishName,
		TypeDesc:        description,
	}
	tx := db.database.Create(&teacherType)
	if tx.Error != nil {
		panic("初始化教师类型数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 教师类型表 [%s-%s] 成功\n", name, description)
	}
}

// InitTeacherData 初始化教师数据
func (db *DbOperate) InitTeacherData(name, englishName, id, ethnic string, sex bool, typeName, unitName string) {
	var teacherType = do.CsTeacherType{}
	db.database.Where("type_name = ?", typeName).First(&teacherType)
	var unit = do.CsDepartment{}
	db.database.Where("department_name = ?", unitName).First(&unit)

	var teacher = do.CsTeacher{
		TeacherUUID: utils.GenerateUUIDNoDash(),
		UnitUUID:    unit.DepartmentUUID,
		ID:          id,
		Name:        name,
		EnglishName: englishName,
		Ethnic:      ethnic,
		Sex:         sex,
		Type:        teacherType.TeacherTypeUUID,
	}

	tx := db.database.Create(&teacher)
	if tx.Error != nil {
		panic("初始化教师数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 教师表 [%s-%s] 成功\n", name, id)
	}
}
