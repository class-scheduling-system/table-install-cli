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
	"frontleaves-table-install-cli/utils"
	"time"
)

func (is *InitStruct) OperateSetupDepartment() {
	// 初始化 cs_department 数据表
	is.operate.InitDepartmentData(
		"LB01", "院系", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		false,
	)
	is.operate.InitDepartmentData(
		"LB03", "行政类", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		false,
	)
	is.operate.InitDepartmentData(
		"LB02", "科研机构", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		false,
	)
	is.operate.InitDepartmentData(
		"LB04", "设计类", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		false,
	)
	is.operate.InitDepartmentData(
		"LB99", "其它", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		false,
	)
	is.operate.InitDepartmentData(
		"9001", "客座教师组", "其它", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		false,
	)
	is.operate.InitDepartmentData(
		"1004", "组织人事处", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2011", "新媒体产业学院", "院系", "校企合办", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		true,
		true,
	)
	is.operate.InitDepartmentData(
		"1011", "信息管理中心", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"101101", "超星技术人员", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("信息管理中心"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1010", "后勤管理服务中心", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"101001", "项目办", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("后勤管理服务中心"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2006", "现代农业学院", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2021, 8, 20, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1002", "纪检监察办公室", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1008", "学生工作处", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"100801", "宿管", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("学生工作处"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2002", "教育艺术学院", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"9000", "未分配人员", "其它", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("其它"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1000", "学院领导", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1003", "党群办公室", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2010", "继续教育与培训部", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1009", "计划财务处", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2009", "马克思主义学院", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2021, 8, 20, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1006", "教务处", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2008", "心理教学部", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		true,
		true,
	)
	is.operate.InitDepartmentData(
		"101002", "校医室", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("后勤管理服务中心"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2005", "化学与材料工程学院", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2007", "汽车与智能交通学院", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1012", "教科研处", "科研机构", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("科研机构"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1001", "党委行政办公室", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2004", "信息智能工程学院", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2003", "机械电气工程学院", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1005", "交流合作办公室", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		false,
		false,
		false,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"2001", "公共基础学院", "院系", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("院系"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		true,
		true,
		false,
		true,
	)
	is.operate.InitDepartmentData(
		"1007", "招生就业办", "行政类", "直属（校办）", nil,
		nil, nil, nil, nil, utils.Ptr("行政类"), nil, nil, nil,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // 建立年月
		nil,
		true,
		false,
		false,
		false,
		true,
	)
}
