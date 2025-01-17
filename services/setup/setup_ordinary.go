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

import "frontleaves-table-install-cli/utils"

func (is *InitStruct) OperateSetupSystem() {
	// 初始化 cs_system 数据表
	is.operate.InitSystemData("author", "锋楪技术团队")
}

func (is *InitStruct) OperateSetupOrdinary() {
	// 初始化 cs_course_nature 数据表
	is.operate.InitCourseNatureData("必修课", "所有学生必须修读的课程")
	is.operate.InitCourseNatureData("实习实训", "通过实践提高学生综合能力的课程")
	is.operate.InitCourseNatureData("专业选修课", "学生根据专业方向选择的课程")
	is.operate.InitCourseNatureData("公共必修课", "全校范围内必须修读的基础课程")
	is.operate.InitCourseNatureData("公共基础课", "全校通用的基础性课程")
	is.operate.InitCourseNatureData("公共选修课", "全校范围内学生自主选择的课程")
	is.operate.InitCourseNatureData("专业拓展课", "拓展学生专业知识领域的课程")
	is.operate.InitCourseNatureData("其他", "其他类型的课程")

	// 初始化 cs_course_property 数据表
	is.operate.InitCoursePropertyData("必修", "学生必须完成的课程，通常是核心课程")
	is.operate.InitCoursePropertyData("专业基础课", "为专业课程学习打下基础的课程")
	is.operate.InitCoursePropertyData("专业核心课", "专业领域内的核心课程，通常不可替代")
	is.operate.InitCoursePropertyData("实习实训", "注重实践技能训练的课程")
	is.operate.InitCoursePropertyData("限选", "学生需要在有限选项中选择的课程")
	is.operate.InitCoursePropertyData("任选", "学生可根据兴趣自由选择的课程")
	is.operate.InitCoursePropertyData("专业选修课", "根据专业方向开设的选修课程")
	is.operate.InitCoursePropertyData("公共基础课", "全校范围内通用的基础课程")
	is.operate.InitCoursePropertyData("公共选修课", "全校范围内的选修课程，学生可自由选择")
	is.operate.InitCoursePropertyData("其他", "未分类或特殊类型的课程")

	// 初始化 cs_course_type 数据表
	is.operate.InitCourseTypeData("A类(纯理论课)", "以理论知识为主的课程")
	is.operate.InitCourseTypeData("B类((理论+实践)课)", "理论与实践相结合的课程")
	is.operate.InitCourseTypeData("C类(纯实践课)", "以实践操作为主的课程")

	// 初始化 cs_course_category 数据表
	is.operate.InitCourseCategoryData("实习实训(必修)", "以实践技能训练为主的必修课程")
	is.operate.InitCourseCategoryData("专业基础课(必修)", "为专业学习打下基础的必修课程")
	is.operate.InitCourseCategoryData("专业核心课(必修)", "专业领域内的核心必修课程")
	is.operate.InitCourseCategoryData("专业课(选修)", "根据学生兴趣选择的专业方向课程")
	is.operate.InitCourseCategoryData("公共基础课(必修)", "全校通用的基础必修课程")
	is.operate.InitCourseCategoryData("公共课(选修)", "全校范围内的选修课程")

	// 初始化 cs_campus 数据表
	is.operate.InitCampusData("铁门关校区", "学校在铁门关的校区", "新疆生产建设兵团铁门关市", "TMGXQ")
	is.operate.InitCampusData("库尔勒校区", "学校在库尔勒的校区", "新疆巴音郭楞蒙古自治州库尔勒市", "KELXQ")

	// 初始化 cs_building 数据表
	is.operate.InitBuildingData("现代农业实训中心", "铁门关校区")
	is.operate.InitBuildingData("汽车工程实训中心", "铁门关校区")
	is.operate.InitBuildingData("信息技术中心", "铁门关校区")
	is.operate.InitBuildingData("教育艺术中心", "铁门关校区")
	is.operate.InitBuildingData("教学楼", "铁门关校区")
	is.operate.InitBuildingData("机电工程实训中心", "铁门关校区")
	is.operate.InitBuildingData("化学工程实训中心", "铁门关校区")
	is.operate.InitBuildingData("实训楼", "库尔勒校区")
	is.operate.InitBuildingData("技能训练中心", "库尔勒校区")
	is.operate.InitBuildingData("综合教学楼", "库尔勒校区")
	is.operate.InitBuildingData("车间", "库尔勒校区")

	// 初始化 cs_classroom_type 数据表
	is.operate.InitClassroomTypeData("多媒体教室", "配备多媒体设备的教室")
	is.operate.InitClassroomTypeData("实训车间", "用于实践技能训练的车间")
	is.operate.InitClassroomTypeData("机房", "配备计算机的教室")
	is.operate.InitClassroomTypeData("汽修实训室", "用于汽车维修技能训练的教室")
	is.operate.InitClassroomTypeData("琴房", "用于音乐教学的教室")
	is.operate.InitClassroomTypeData("美术教室", "用于美术课程教学的教室")
	is.operate.InitClassroomTypeData("舞蹈教室", "用于舞蹈课程教学的教室")
	is.operate.InitClassroomTypeData("声乐教室", "用于声乐课程教学的教室")
	is.operate.InitClassroomTypeData("化工仿真实训室", "用于化工专业仿真训练的教室")
	is.operate.InitClassroomTypeData("虚拟仿真实训室", "用于虚拟仿真教学的教室")
	is.operate.InitClassroomTypeData("自动化生产线实训室", "用于自动化生产线技能训练的教室")
	is.operate.InitClassroomTypeData("液压气动实训室", "用于液压气动专业技能训练的教室")
	is.operate.InitClassroomTypeData("电工实训室", "用于电工技能训练的教室")
	is.operate.InitClassroomTypeData("网络实训室", "用于网络技能训练的教室")
	is.operate.InitClassroomTypeData("幼儿保健室", "用于幼儿保健教学的教室")
	is.operate.InitClassroomTypeData("机械装调实训室", "用于机械装配与调试技能训练的教室")
	is.operate.InitClassroomTypeData("网络综合布线室", "用于网络布线技能训练的教室")
	is.operate.InitClassroomTypeData("公共教室", "全校范围通用的教室")
	is.operate.InitClassroomTypeData("蒙台梭立实训室", "用于蒙台梭利教学法实训的教室")
	is.operate.InitClassroomTypeData("活动室", "用于课外活动和兴趣小组的教室")
	is.operate.InitClassroomTypeData("移动一体机", "配备移动设备的灵活教学空间")
	is.operate.InitClassroomTypeData("普通教室", "以常规教学为主的教室")
	is.operate.InitClassroomTypeData("现代电气控制实训室", "用于现代电气控制技能训练的教室")

	// 初始化 cs_classroom_tag 数据表
	is.operate.InitClassroomTagData("多媒体", "配备多媒体设备的教室")
	is.operate.InitClassroomTagData("计算机", "配备计算机设备的教室或机房")
	is.operate.InitClassroomTagData("实验室", "用于科学实验的专用教室")
	is.operate.InitClassroomTagData("音乐", "用于音乐教学及练习的教室")
	is.operate.InitClassroomTagData("美术", "用于美术课程教学的教室")
	is.operate.InitClassroomTagData("舞蹈", "用于舞蹈课程教学及排练的教室")
	is.operate.InitClassroomTagData("理论", "以理论教学为主的普通教室")
	is.operate.InitClassroomTagData("实践", "以实践技能训练为主的教室")
	is.operate.InitClassroomTagData("汽车维修", "用于汽车维修技能训练的教室")
	is.operate.InitClassroomTagData("化工", "用于化工类课程实验与训练的教室")
	is.operate.InitClassroomTagData("电工", "用于电工技能教学的专用教室")
	is.operate.InitClassroomTagData("网络", "用于网络布线及调试训练的教室")
	is.operate.InitClassroomTagData("通用", "全校范围通用的普通教室")
	is.operate.InitClassroomTagData("幼教", "用于幼儿教育相关课程教学的教室")
	is.operate.InitClassroomTagData("机械装调", "用于机械装配与调试技能训练的教室")
	is.operate.InitClassroomTagData("虚拟仿真", "用于虚拟仿真教学的专用教室")
	is.operate.InitClassroomTagData("自动化", "用于自动化生产线技能训练的教室")
	is.operate.InitClassroomTagData("液压气动", "用于液压气动设备训练的教室")
	is.operate.InitClassroomTagData("活动室", "用于课外活动和兴趣小组的教室")
	is.operate.InitClassroomTagData("灵活教学", "配备灵活移动设备的多功能教学空间")

	// 初始化 cs_tables_chairs_type 数据表
	is.operate.InitTableChairTypeData("实验椅", utils.Ptr("适用于实验室环境的高脚椅或专用椅子"), nil)
	is.operate.InitTableChairTypeData("绘图椅", utils.Ptr("适用于工科学生绘图课的椅子，通常带有高度调节功能"), nil)
	is.operate.InitTableChairTypeData("工程椅", utils.Ptr("用于工程类实验室或实训室，耐用且符合人体工程学设计"), nil)
	is.operate.InitTableChairTypeData("机房椅", utils.Ptr("适用于计算机机房，带有滑轮和可调节高度的功能"), nil)
	is.operate.InitTableChairTypeData("电工实训椅", utils.Ptr("用于电工实训室的专用椅子，耐用且易于清洁"), nil)
	is.operate.InitTableChairTypeData("仿真实训椅", utils.Ptr("用于虚拟仿真实训室的专业座椅，通常与工作台配套"), nil)
	is.operate.InitTableChairTypeData("多功能座椅", utils.Ptr("适用于多媒体教室或报告厅，通常带有折叠写字板"), nil)
	is.operate.InitTableChairTypeData("排椅", utils.Ptr("适用于大型教室或阶梯教室的连排座椅"), nil)
	is.operate.InitTableChairTypeData("焊接实训椅", utils.Ptr("用于焊接技能训练，耐高温且稳固"), nil)
	is.operate.InitTableChairTypeData("车间工作椅", utils.Ptr("用于车间或工厂环境的椅子，结实耐用"), nil)
	is.operate.InitTableChairTypeData("实训车间椅", utils.Ptr("专为机械或电气实训车间设计的椅子"), nil)
	is.operate.InitTableChairTypeData("机器人实验椅", utils.Ptr("适用于机器人实验室，通常配有高度调节功能"), nil)
	is.operate.InitTableChairTypeData("休息椅", utils.Ptr("校园公共区域或实验室休息区的普通座椅"), nil)
	is.operate.InitTableChairTypeData("图书馆座椅", utils.Ptr("适用于图书馆安静学习区的专用椅子"), nil)
	is.operate.InitTableChairTypeData("会议椅", utils.Ptr("适用于学术会议室或研讨室的舒适椅子"), nil)

	// 初始化 cs_unit_category 数据表
	is.operate.InitUnitCategoryData("其它", nil, utils.Ptr("不属于特定分类的单位"))
	is.operate.InitUnitCategoryData("行政类", nil, utils.Ptr("学校的行政管理部门"))
	is.operate.InitUnitCategoryData("院系", nil, utils.Ptr("负责教学与学生管理的学院或系"))
	is.operate.InitUnitCategoryData("科研机构", nil, utils.Ptr("从事科学研究的机构"))

	// 初始化 cs_unit_type 数据表
	is.operate.InitUnitTypeData("直属（校办）", utils.Ptr("DirectlyUnder(SchoolOffice)"), utils.Ptr("直属"))
	is.operate.InitUnitTypeData("校企合办", utils.Ptr("SchoolEnterpriseJointly"), utils.Ptr("校企"))
}
