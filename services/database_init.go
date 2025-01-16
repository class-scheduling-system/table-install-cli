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
	"embed"
	"frontleaves-table-install-cli/database"
	"github.com/pelletier/go-toml"
)

func InitDatabase(config *toml.Tree, resourcesFile embed.FS) {
	// 初始化数据库
	CreateDatabase(config, resourcesFile)

	// 初始化数据
	operate := database.NewDatabaseOperate(config)
	// 初始化 cs_system 数据表
	operate.InitSystemData("author", "锋楪技术团队")

	// 初始化 cs_course_nature 数据表
	operate.InitCourseNatureData("必修课", "所有学生必须修读的课程")
	operate.InitCourseNatureData("实习实训", "通过实践提高学生综合能力的课程")
	operate.InitCourseNatureData("专业选修课", "学生根据专业方向选择的课程")
	operate.InitCourseNatureData("公共必修课", "全校范围内必须修读的基础课程")
	operate.InitCourseNatureData("公共基础课", "全校通用的基础性课程")
	operate.InitCourseNatureData("公共选修课", "全校范围内学生自主选择的课程")
	operate.InitCourseNatureData("专业拓展课", "拓展学生专业知识领域的课程")
	operate.InitCourseNatureData("其他", "其他类型的课程")

	// 初始化 cs_course_property 数据表
	operate.InitCoursePropertyData("必修", "学生必须完成的课程，通常是核心课程")
	operate.InitCoursePropertyData("专业基础课", "为专业课程学习打下基础的课程")
	operate.InitCoursePropertyData("专业核心课", "专业领域内的核心课程，通常不可替代")
	operate.InitCoursePropertyData("实习实训", "注重实践技能训练的课程")
	operate.InitCoursePropertyData("限选", "学生需要在有限选项中选择的课程")
	operate.InitCoursePropertyData("任选", "学生可根据兴趣自由选择的课程")
	operate.InitCoursePropertyData("专业选修课", "根据专业方向开设的选修课程")
	operate.InitCoursePropertyData("公共基础课", "全校范围内通用的基础课程")
	operate.InitCoursePropertyData("公共选修课", "全校范围内的选修课程，学生可自由选择")
	operate.InitCoursePropertyData("其他", "未分类或特殊类型的课程")

	// 初始化 cs_course_type 数据表
	operate.InitCourseTypeData("A类(纯理论课)", "以理论知识为主的课程")
	operate.InitCourseTypeData("B类((理论+实践)课)", "理论与实践相结合的课程")
	operate.InitCourseTypeData("C类(纯实践课)", "以实践操作为主的课程")

	// 初始化 cs_course_category 数据表
	operate.InitCourseCategoryData("实习实训(必修)", "以实践技能训练为主的必修课程")
	operate.InitCourseCategoryData("专业基础课(必修)", "为专业学习打下基础的必修课程")
	operate.InitCourseCategoryData("专业核心课(必修)", "专业领域内的核心必修课程")
	operate.InitCourseCategoryData("专业课(选修)", "根据学生兴趣选择的专业方向课程")
	operate.InitCourseCategoryData("公共基础课(必修)", "全校通用的基础必修课程")
	operate.InitCourseCategoryData("公共课(选修)", "全校范围内的选修课程")

	// 初始化 cs_campus 数据表
	operate.InitCampusData("铁门关校区", "学校在铁门关的校区", "新疆生产建设兵团铁门关市", "TMGXQ")
	operate.InitCampusData("库尔勒校区", "学校在库尔勒的校区", "新疆巴音郭楞蒙古自治州库尔勒市", "KELXQ")

	// 初始化 cs_building 数据表
	operate.InitBuildingData("现代农业实训中心", "铁门关校区")
	operate.InitBuildingData("汽车工程实训中心", "铁门关校区")
	operate.InitBuildingData("信息技术中心", "铁门关校区")
	operate.InitBuildingData("教育艺术中心", "铁门关校区")
	operate.InitBuildingData("教学楼", "铁门关校区")
	operate.InitBuildingData("机电工程实训中心", "铁门关校区")
	operate.InitBuildingData("化学工程实训中心", "铁门关校区")
	operate.InitBuildingData("实训楼", "库尔勒校区")
	operate.InitBuildingData("技能训练中心", "库尔勒校区")
	operate.InitBuildingData("综合教学楼", "库尔勒校区")
	operate.InitBuildingData("车间", "库尔勒校区")
}
