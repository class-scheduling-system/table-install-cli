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

// SetupTeacherCourseQualificationData 初始化教师课程资格数据
func (is *InitStruct) OperateSetupTeacherCourseQualification() {

	// 为张三老师添加课程资格（农业相关课程）
	is.operate.InitTeacherCourseQualificationData(
		"张三",
		"植物生理学",
		3, // 高级资格
		true,
		5,
		"多年教学经验，专业知识扎实",
	)
	is.operate.InitTeacherCourseQualificationData(
		"张三",
		"农业经营与管理",
		2, // 中级资格
		false,
		3,
		"具备农业管理基础知识",
	)

	// 为李四老师添加课程资格（作物相关课程）
	is.operate.InitTeacherCourseQualificationData(
		"李四",
		"作物栽培学",
		3, // 高级资格
		true,
		6,
		"资深作物栽培专家",
	)
	is.operate.InitTeacherCourseQualificationData(
		"李四",
		"土壤肥料学",
		2, // 中级资格
		true,
		4,
		"具有丰富的土壤研究经验",
	)

	// 为王五老师添加课程资格（棉花相关课程）
	is.operate.InitTeacherCourseQualificationData(
		"王五",
		"棉花加工",
		3, // 高级资格
		true,
		7,
		"棉花加工领域专家",
	)
	is.operate.InitTeacherCourseQualificationData(
		"王五",
		"棉花检验",
		3, // 高级资格
		true,
		5,
		"具有丰富的棉花质量检验经验",
	)
}
