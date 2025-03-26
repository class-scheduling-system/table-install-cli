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

package setup

import (
	"fmt"
	"time"
)

// SetupSemesterData 初始化学期数据
func (is *InitStruct) OperateSetupSemester() {

	// 2022-2023学年第一学期
	is.operate.InitSemesterData(
		"2022-2023-1",
		"2022-2023学年第一学期",
		time.Date(2022, 9, 1, 0, 0, 0, 0, time.Local),
		time.Date(2023, 1, 15, 0, 0, 0, 0, time.Local),
		false,
	)

	// 2022-2023学年第二学期
	is.operate.InitSemesterData(
		"2022-2023-2",
		"2022-2023学年第二学期",
		time.Date(2023, 2, 20, 0, 0, 0, 0, time.Local),
		time.Date(2023, 7, 10, 0, 0, 0, 0, time.Local),
		false,
	)

	// 2023-2024学年第一学期
	is.operate.InitSemesterData(
		"2023-2024-1",
		"2023-2024学年第一学期",
		time.Date(2023, 9, 1, 0, 0, 0, 0, time.Local),
		time.Date(2024, 1, 15, 0, 0, 0, 0, time.Local),
		false,
	)

	// 2023-2024学年第二学期
	semester2024Spring := is.operate.InitSemesterData(
		"2023-2024-2",
		"2023-2024学年第二学期",
		time.Date(2024, 2, 20, 0, 0, 0, 0, time.Local),
		time.Date(2024, 7, 10, 0, 0, 0, 0, time.Local),
		true, // 当前学期
	)

	fmt.Println("开始初始化教师喜好数据...")

	// 为张三老师添加喜好数据
	is.operate.InitTeacherPreferencesData("张三", semester2024Spring, 1, 1, 5, "早上第一节精力充沛") // 周一第一节课
	is.operate.InitTeacherPreferencesData("张三", semester2024Spring, 1, 2, 4, "上午状态良好")    // 周一第二节课
	is.operate.InitTeacherPreferencesData("张三", semester2024Spring, 5, 3, 1, "周五中午要开会")   // 周五第三节课

	// 为李四老师添加喜好数据
	is.operate.InitTeacherPreferencesData("李四", semester2024Spring, 2, 1, 4, "早上状态好")     // 周二第一节课
	is.operate.InitTeacherPreferencesData("李四", semester2024Spring, 2, 4, 2, "下午容易疲劳")    // 周二第四节课
	is.operate.InitTeacherPreferencesData("李四", semester2024Spring, 4, 3, 5, "这个时间最适合上课") // 周四第三节课

	// 为王五老师添加喜好数据
	is.operate.InitTeacherPreferencesData("王五", semester2024Spring, 3, 2, 5, "最佳授课时间")  // 周三第二节课
	is.operate.InitTeacherPreferencesData("王五", semester2024Spring, 3, 3, 4, "状态良好")    // 周三第三节课
	is.operate.InitTeacherPreferencesData("王五", semester2024Spring, 5, 4, 1, "周五下午要开会") // 周五第四节课
}
