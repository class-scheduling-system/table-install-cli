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
