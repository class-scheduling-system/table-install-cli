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

package setup

import (
	"fmt"
	"frontleaves-table-install-cli/models/do"
	"frontleaves-table-install-cli/utils"
	"time"
)

// OperateSetupUser 创建一些测试用户
func (is *InitStruct) OperateSetupUser() {

	// 创建一个管理员测试用户
	var getRoleAdmin = do.CsRole{}
	is.operate.GetDatabase().Where("role_name = '管理员'").First(&getRoleAdmin)
	var newAdminUser = do.CsUser{
		UserUUID:   utils.GenerateUUIDNoDash(),
		Name:       "test_admin",
		Password:   utils.Encrypt("admin"),
		Email:      "admin@x-lf.cn",
		Phone:      "18888888880",
		RoleUUID:   getRoleAdmin.RoleUUID,
		Permission: utils.Ptr("[]"),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	tx := is.operate.GetDatabase().Create(&newAdminUser)
	if tx.Error != nil {
		panic(tx.Error)
	}
	fmt.Println("初始化 测试用户[管理员] 成功")

	// 创建一个管理用户
	var getRoleManager = do.CsRole{}
	is.operate.GetDatabase().Where("role_name = '管理'").First(&getRoleManager)
	var newManagerUser = do.CsUser{
		UserUUID:   utils.GenerateUUIDNoDash(),
		Name:       "test_manager",
		Password:   utils.Encrypt("manager"),
		Email:      "manager@x-lf.cn",
		Phone:      "18888888881",
		RoleUUID:   getRoleAdmin.RoleUUID,
		Permission: utils.Ptr("[]"),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	tx = is.operate.GetDatabase().Create(&newManagerUser)
	if tx.Error != nil {
		panic(tx.Error)
	}
	fmt.Println("初始化 测试用户[管理] 成功")

	// 创建一个教务用户
	var getRoleAcademic = do.CsRole{}
	is.operate.GetDatabase().Where("role_name = '教务'").First(&getRoleAcademic)
	var getNewAcademicUserUUID = utils.GenerateUUIDNoDash()
	var newAcademicUser = do.CsUser{
		UserUUID:   getNewAcademicUserUUID,
		Name:       "test_academic",
		Password:   utils.Encrypt("academic"),
		Email:      "academic@x-lf.cn",
		Phone:      "18888888882",
		RoleUUID:   getRoleAcademic.RoleUUID,
		Permission: utils.Ptr("[]"),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	is.operate.GetDatabase().Create(&newAcademicUser)
	var getDepartment = do.CsDepartment{}
	is.operate.GetDatabase().Where("department_name = '现代农业学院'").First(&getDepartment)
	var newAcademicAffairsPermission = do.CsAcademicAffairsPermission{
		AcademicAffairsPermissionUUID: utils.GenerateUUIDNoDash(),
		AuthorizedUser:                getNewAcademicUserUUID,
		Department:                    getDepartment.DepartmentUUID,
		Type:                          0,
		CreatedAt:                     time.Now(),
		UpdatedAt:                     time.Now(),
	}
	tx = is.operate.GetDatabase().Create(&newAcademicAffairsPermission)
	if tx.Error != nil {
		panic(tx.Error)
	}
	fmt.Println("初始化 测试用户[教务] 成功")

	// 创建一个教师用户
	var getRoleTeacher = do.CsRole{}
	is.operate.GetDatabase().Where("role_name = '教师'").First(&getRoleTeacher)
	var getTeacher = do.CsTeacher{}
	is.operate.GetDatabase().Where("name = '测试教师'").First(&getTeacher)
	getTeacher.UserUUID = utils.Ptr(utils.GenerateUUIDNoDash())
	var newTeacherUser = do.CsUser{
		UserUUID:   *getTeacher.UserUUID,
		Name:       "test_teacher",
		Password:   utils.Encrypt("teacher"),
		Email:      "teacher@x-lf.cn",
		Phone:      "18888888883",
		RoleUUID:   getRoleTeacher.RoleUUID,
		Permission: utils.Ptr("[]"),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	tx = is.operate.GetDatabase().Create(&newTeacherUser)
	if tx.Error != nil {
		panic(tx.Error)
	}
	is.operate.GetDatabase().Save(&getTeacher)
	fmt.Println("初始化 测试用户[教师] 成功")

	// 创建一个学生用户
	var getRoleStudent = do.CsRole{}
	is.operate.GetDatabase().Where("role_name = '学生'").First(&getRoleStudent)
	var getStudent = do.CsStudent{}
	is.operate.GetDatabase().Where("name = '筱锋'").First(&getStudent)
	getStudent.UserUUID = utils.Ptr(utils.GenerateUUIDNoDash())
	var newStudentUser = do.CsUser{
		UserUUID:   *getStudent.UserUUID,
		Name:       "test_student",
		Password:   utils.Encrypt("student"),
		Email:      "student@x-lf.cn",
		Phone:      "18888888884",
		Permission: utils.Ptr("[]"),
		RoleUUID:   getRoleStudent.RoleUUID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	tx = is.operate.GetDatabase().Create(&newStudentUser)
	if tx.Error != nil {
		panic(tx.Error)
	}
	is.operate.GetDatabase().Save(&getStudent)
	fmt.Println("初始化 测试用户[学生] 成功")
}
