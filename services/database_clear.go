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

package services

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func ClearDatabase(config *toml.Tree) {
	// 数据表倒序清空数据的顺序
	var tables = []string{
		"cs_class_assignment",            // 排课表
		"cs_academic_affairs_permission", // 教务权限表
		"cs_course_library",              // 课程库表
		"cs_teacher",                     // 教师表
		"cs_student",                     // 学生表
		"cs_classroom",                   // 教室表
		"cs_classroom_type",              // 教室类型表
		"cs_classroom_tag",               // 教室标签表
		"cs_department",                  // 部门表
		"cs_building",                    // 教学楼表
		"cs_campus",                      // 校区表
		"cs_course_category",             // 课程类别表
		"cs_course_nature",               // 课程性质表
		"cs_course_property",             // 课程属性表
		"cs_course_type",                 // 课程类型表
		"cs_credit_hour_type",            // 学时类型表
		"cs_major",                       // 专业表
		"cs_semester",                    // 学期表
		"cs_tables_chairs_type",          // 桌椅类型表
		"cs_unit_category",               // 单位类别表
		"cs_unit_type",                   // 单位类型表
		"cs_user",                        // 用户表
		"cs_role",                        // 角色表
		"cs_permission",                  // 权限管理表
		"cs_system",                      // 系统配置相关表
	}

	// 从配置文件中获取数据库连接信息
	host := config.Get("database.host").(string)
	port := config.Get("database.port").(int64)
	name := config.Get("database.name").(string)
	user := config.Get("database.user").(string)
	pass := config.Get("database.password").(string)

	// 数据库连接字符串
	dsn := user + ":" + pass + "@tcp(" + host + ":" + strconv.FormatInt(port, 10) + ")/" + name + "?charset=utf8mb4"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 倒序清空表数据
	for _, table := range tables {
		sql := fmt.Sprintf("DELETE FROM `%s`;", table)
		if err := db.Exec(sql).Error; err != nil {
			fmt.Printf("清空表 %s 数据失败: %v\n", table, err)
		} else {
			fmt.Printf("表 %s 数据已成功清空。\n", table)
		}
	}
}
