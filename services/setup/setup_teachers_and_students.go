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

import "fmt"

func (is *InitStruct) OperateSetupTeacherAndStudent() {
	// 初始化教师数据
	is.operate.InitTeacherData("张三", "Zhang San", "T001", "汉族", true, "教授", "马克思主义学院")
	is.operate.InitTeacherData("李四", "Li Si", "T002", "汉族", true, "教授", "马克思主义学院")
	is.operate.InitTeacherData("王五", "Wang Wu", "T003", "汉族", true, "辅导员", "马克思主义学院")
	is.operate.InitTeacherData("赵六", "Zhao Liu", "T004", "汉族", true, "教授", "教育艺术学院")
	is.operate.InitTeacherData("孙七", "Sun Qi", "T005", "汉族", true, "教授", "教育艺术学院")
	is.operate.InitTeacherData("周八", "Zhou Ba", "T006", "汉族", true, "讲师", "教育艺术学院")

	// 初始化学生数据
	// 畜智1班 39名学生
	for i := 1; i <= 39; i++ {
		is.operate.InitStudentData(
			fmt.Sprintf("李%d", i),
			fmt.Sprintf("CC100001S%03d", i),
			"2022",
			"现代农业学院",
			"畜禽智能化养殖",
			"畜智1班",
			i%2 == 1,
		)
	}

	// 畜智2班 20名学生
	for i := 1; i <= 20; i++ {
		is.operate.InitStudentData(
			fmt.Sprintf("王%d", i),
			fmt.Sprintf("CC100002S%03d", i),
			"2022",
			"现代农业学院",
			"畜禽智能化养殖",
			"畜智2班",
			i%2 == 1,
		)
	}

	// 棉经1班 39名学生
	for i := 1; i <= 39; i++ {
		is.operate.InitStudentData(
			fmt.Sprintf("张%d", i),
			fmt.Sprintf("CC100003S%03d", i),
			"2022",
			"现代农业学院",
			"棉花加工与经营管理",
			"棉经1班",
			i%2 == 1,
		)
	}

	// 棉经2班 16名学生
	for i := 1; i <= 16; i++ {
		is.operate.InitStudentData(
			fmt.Sprintf("刘%d", i),
			fmt.Sprintf("CC100004S%03d", i),
			"2022",
			"现代农业学院",
			"棉花加工与经营管理",
			"棉经2班",
			i%2 == 1,
		)
	}

	// 棉经3班 42名学生
	for i := 1; i <= 42; i++ {
		is.operate.InitStudentData(
			fmt.Sprintf("陈%d", i),
			fmt.Sprintf("CC100005S%03d", i),
			"2022",
			"现代农业学院",
			"棉花加工与经营管理",
			"棉经3班",
			i%2 == 1,
		)
	}

	// 农经1班 15名学生
	for i := 1; i <= 15; i++ {
		is.operate.InitStudentData(
			fmt.Sprintf("赵%d", i),
			fmt.Sprintf("CC100006S%03d", i),
			"2022",
			"现代农业学院",
			"现代农业经济管理",
			"农经1班",
			i%2 == 1,
		)
	}

	// 农经2班 18名学生
	for i := 1; i <= 18; i++ {
		is.operate.InitStudentData(
			fmt.Sprintf("黄%d", i),
			fmt.Sprintf("CC100007S%03d", i),
			"2022",
			"现代农业学院",
			"现代农业经济管理",
			"农经2班",
			i%2 == 1,
		)
	}

	// 畜兽1班 42名学生
	for i := 1; i <= 42; i++ {
		is.operate.InitStudentData(
			fmt.Sprintf("周%d", i),
			fmt.Sprintf("CC100008S%03d", i),
			"2022",
			"现代农业学院",
			"畜牧兽医",
			"畜兽1班",
			i%2 == 1,
		)
	}
}
