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
	"time"
)

// MajorData 插入一条新的专业记录到数据库。
// 参数:
//   - name: 专业的名称，字符串类型，必填。
//   - description: 专业的描述，可选指针类型，允许为nil表示没有描述信息。
//   - code: 专业的代码，字符串类型，用以唯一标识专业。
//   - department: 所属院系的名称，用于关联到院系表中的记录。
//
// 此函数首先根据提供的院系名称查询对应的院系记录，然后创建一个新的CsMajor结构体实例，
// 其中包含生成的UUID、专业名称、描述、代码、状态（默认为true即启用）、所属院系的UUID以及创建和更新时间戳。
// 随后，通过GORM的Create方法将新专业数据存入数据库。
// 如果在创建过程中发生错误，则函数会恐慌并打印错误信息；如果成功，则输出初始化成功的提示信息。
func (db *DbOperate) MajorData(name string, description *string, code, department string, year uint16, trainingLevel string) {
	var departmentDO = do.CsDepartment{}
	db.database.Where("department_name = ?", department).First(&departmentDO)
	major := do.CsMajor{
		MajorUUID:        utils.GenerateUUIDNoDash(),  // 生成无短划线的UUID作为主键
		MajorName:        name,                        // 设置专业名称
		MajorDescription: description,                 // 设置专业描述，可能为nil
		MajorCode:        code,                        // 设置专业代码
		MajorStatus:      true,                        // 设置专业状态为启用，默认为true表示该专业当前是启用状态
		EducationYears:   year,                        // 设置学制年限，以整数形式表示
		TrainingLevel:    trainingLevel,               // 设置培养层次，例如："本科"、"硕士"、"博士"
		DepartmentUUID:   departmentDO.DepartmentUUID, // 关联院系UUID
		CreatedAt:        time.Now(),                  // 记录创建时间
		UpdatedAt:        time.Now(),                  // 记录更新时间（此处与创建时间为同一时刻）
		DepartmentRel:    nil,                         // 这里假设不需要直接存储部门关系对象，故设为nil
	}
	tx := db.database.Create(&major)
	if tx.Error != nil {
		panic("初始化专业表失败：" + tx.Error.Error())
	} else {
		fmt.Printf("初始化 专业表[%s] (%s) 成功\n", name, code)
	}
}
