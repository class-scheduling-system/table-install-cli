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
	// 马克思主义学院教师课程资格
	// 张三老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"张三",
		"中国特色社会主义",
		3, // 高级资格
		true,
		8,
		"马克思主义理论研究专家，教学经验丰富",
	)
	is.operate.InitTeacherCourseQualificationData(
		"张三",
		"思想道德与法治",
		3, // 高级资格
		true,
		6,
		"思政课教学名师",
	)

	// 李四老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"李四",
		"中国近现代史纲要",
		3, // 高级资格
		true,
		7,
		"历史研究专家，教学经验丰富",
	)
	is.operate.InitTeacherCourseQualificationData(
		"李四",
		"哲学与人生",
		2, // 中级资格
		false,
		4,
		"理论研究扎实",
	)

	// 王五老师
	is.operate.InitTeacherCourseQualificationData(
		"王五",
		"形势与政策",
		2, // 中级资格
		true,
		5,
		"关注时事政治，善于分析国际形势",
	)

	// 黄晓明老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"黄晓明",
		"哲学与人生",
		2, // 中级资格
		true,
		4,
		"理论功底扎实，教学方法新颖",
	)
	is.operate.InitTeacherCourseQualificationData(
		"黄晓明",
		"思想道德与法治",
		2, // 中级资格
		false,
		3,
		"法学背景，善于案例教学",
	)

	// 朱丽华老师
	is.operate.InitTeacherCourseQualificationData(
		"朱丽华",
		"形势与政策",
		1, // 初级资格
		true,
		2,
		"年轻教师，教学热情高",
	)

	// 徐国强老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"徐国强",
		"中国近现代史纲要",
		3, // 高级资格
		true,
		9,
		"历史学专家，研究成果丰富",
	)
	is.operate.InitTeacherCourseQualificationData(
		"徐国强",
		"哲学与人生",
		3, // 高级资格
		true,
		8,
		"理论研究深入，教学经验丰富",
	)

	// 教育艺术学院教师课程资格
	// 赵六老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"赵六",
		"教师口语-1",
		3, // 高级资格
		true,
		10,
		"语言表达专家，教学经验丰富",
	)
	is.operate.InitTeacherCourseQualificationData(
		"赵六",
		"教师口语-2",
		3, // 高级资格
		true,
		10,
		"语言表达专家，教学经验丰富",
	)

	// 孙七老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"孙七",
		"学前儿童社会教育",
		3, // 高级资格
		true,
		8,
		"幼儿教育专家，实践经验丰富",
	)
	is.operate.InitTeacherCourseQualificationData(
		"孙七",
		"教师口语-1",
		2, // 中级资格
		false,
		5,
		"语言表达能力强",
	)

	// 周八老师
	is.operate.InitTeacherCourseQualificationData(
		"周八",
		"教师口语-2",
		2, // 中级资格
		true,
		4,
		"普通话水平一级甲等",
	)

	// 马云峰老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"马云峰",
		"教师口语-1",
		2, // 中级资格
		true,
		6,
		"语言教育专家",
	)
	is.operate.InitTeacherCourseQualificationData(
		"马云峰",
		"学前儿童社会教育",
		2, // 中级资格
		false,
		4,
		"儿童教育研究经验丰富",
	)

	// 韩雪梅老师
	is.operate.InitTeacherCourseQualificationData(
		"韩雪梅",
		"教师口语-2",
		1, // 初级资格
		true,
		3,
		"教学新秀",
	)

	// 董志远老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"董志远",
		"学前儿童社会教育",
		3, // 高级资格
		true,
		12,
		"儿童教育专家，理论与实践经验丰富",
	)
	is.operate.InitTeacherCourseQualificationData(
		"董志远",
		"教师口语-1",
		2, // 中级资格
		false,
		6,
		"教学经验丰富",
	)

	// 现代农业学院教师课程资格
	// 陈明华老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"陈明华",
		"植物生理学",
		3, // 高级资格
		true,
		15,
		"植物学专家，研究成果显著",
	)
	is.operate.InitTeacherCourseQualificationData(
		"陈明华",
		"作物栽培学",
		3, // 高级资格
		true,
		12,
		"农作物栽培专家",
	)

	// 林志强老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"林志强",
		"土壤肥料学",
		3, // 高级资格
		true,
		8,
		"土壤学专家",
	)
	is.operate.InitTeacherCourseQualificationData(
		"林志强",
		"植物生理学",
		2, // 中级资格
		false,
		5,
		"植物学研究经验丰富",
	)

	// 杨丽娟老师
	is.operate.InitTeacherCourseQualificationData(
		"杨丽娟",
		"作物栽培学",
		2, // 中级资格
		true,
		4,
		"农作物栽培研究经验丰富",
	)

	// 刘建国老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"刘建国",
		"农业经营与管理",
		3, // 高级资格
		true,
		10,
		"农业经济专家",
	)
	is.operate.InitTeacherCourseQualificationData(
		"刘建国",
		"棉花加工",
		3, // 高级资格
		true,
		8,
		"棉花加工专家",
	)
	is.operate.InitTeacherCourseQualificationData(
		"刘建国",
		"棉花检验",
		2, // 中级资格
		false,
		5,
		"具有丰富的棉花质量检验经验",
	)

	// 吴秀英老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"吴秀英",
		"棉花加工",
		2, // 中级资格
		true,
		6,
		"棉花加工研究经验丰富",
	)
	is.operate.InitTeacherCourseQualificationData(
		"吴秀英",
		"农业经营与管理",
		2, // 中级资格
		false,
		4,
		"农业经济管理研究经验",
	)

	// 郑伟老师
	is.operate.InitTeacherCourseQualificationData(
		"郑伟",
		"棉花检验",
		2, // 中级资格
		true,
		5,
		"棉花质量检验经验丰富",
	)

	// 测试教师
	is.operate.InitTeacherCourseQualificationData(
		"测试教师",
		"植物生理学",
		1, // 初级资格
		true,
		1,
		"测试用教师账号",
	)

	// 畜牧专业教师课程资格
	// 王建军老师 - 多门课程资格
	is.operate.InitTeacherCourseQualificationData(
		"王建军",
		"畜牧学概论",
		3, // 高级资格
		true,
		12,
		"畜牧学专家，理论与实践经验丰富",
	)
	is.operate.InitTeacherCourseQualificationData(
		"王建军",
		"畜牧场规划与设计",
		3, // 高级资格
		true,
		10,
		"畜牧场规划设计专家",
	)
	is.operate.InitTeacherCourseQualificationData(
		"王建军",
		"畜牧业机械化",
		2, // 中级资格
		false,
		5,
		"具有丰富的畜牧业机械化实践经验",
	)
}
