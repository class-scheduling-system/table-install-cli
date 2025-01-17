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
	"encoding/json"
	"fmt"
	"frontleaves-table-install-cli/models/do"
	"frontleaves-table-install-cli/utils"
)

// InitClassroomTypeData 初始化教室类型数据
func (db *DbOperate) InitClassroomTypeData(name, description string) {
	var classroomType = do.CsClassroomType{
		ClassTypeUUID: utils.GenerateUUIDNoDash(),
		Name:          name,
		Description:   description,
	}
	tx := db.database.Create(&classroomType)
	if tx.Error != nil {
		panic("初始化教室类型数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 教室类型表 [%s-%s] 成功\n", classroomType.Name, classroomType.Description)
	}
}

// InitClassroomTagData 初始化教室标签数据
func (db *DbOperate) InitClassroomTagData(name, description string) {
	var data *string
	if description == "" {
		data = nil
	} else {
		data = &description
	}
	var classroomTag = do.CsClassroomTag{
		ClassTagUUID: utils.GenerateUUIDNoDash(),
		Name:         name,
		Description:  data,
	}
	tx := db.database.Create(&classroomTag)
	if tx.Error != nil {
		panic("初始化教室标签数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 教室标签表 [%s-%s] 成功\n", classroomTag.Name, *classroomTag.Description)
	}
}

// InitClassroom 初始化教室数据
func (db *DbOperate) InitClassroom(
	number, name, campus, building string, floor uint, classroomType string, classroomTag *[]string, capacity uint, examinationRoomCapacity *uint, examinationRoom, isMultimedia, isAirConditioned, status bool, description, managementDepartment *string, area float64, tablesChairsType *string,
) {
	var campusEntity = do.CsCampus{}
	db.database.Where("campus_name = ?", campus).First(&campusEntity)
	var buildingEntity = do.CsBuilding{}
	db.database.Where("building_name = ?", building).First(&buildingEntity)
	var classroomTypeEntity = do.CsClassroomType{}
	db.database.Where("name = ?", classroomType).First(&classroomTypeEntity)
	var tablesChairsTypeUUID *string
	if tablesChairsType != nil {
		var tablesChairsTypeEntity = do.CsTablesChairsType{}
		db.database.Where("name = ?", tablesChairsType).First(&tablesChairsTypeEntity)
		tablesChairsTypeUUID = &tablesChairsTypeEntity.TablesChairsTypeUUID
	} else {
		tablesChairsTypeUUID = nil
	}
	var departmentUUID *string
	if managementDepartment != nil {
		var managementDepartmentEntity = do.CsDepartment{}
		db.database.Where("department_name = ?", managementDepartment).First(&managementDepartmentEntity)
		departmentUUID = &managementDepartmentEntity.DepartmentUUID
	} else {
		departmentUUID = nil
	}
	var classroomTagEntity = make([]string, 0)
	if classroomTag != nil {
		for _, tag := range *classroomTag {
			var tagEntity = do.CsClassroomTag{}
			db.database.Where("name = ?", tag).First(&tagEntity)
			classroomTagEntity = append(classroomTagEntity, tagEntity.ClassTagUUID)
		}
	}
	var tagJsonString *string
	if len(classroomTagEntity) > 0 {
		marshal, err := json.Marshal(classroomTagEntity)
		if err != nil {
			panic("初始化教室数据失败: " + err.Error())
		}
		tagJsonString = utils.Ptr(string(marshal))
	} else {
		tagJsonString = nil
	}

	var classroom = do.CsClassroom{
		ClassroomUUID:           utils.GenerateUUIDNoDash(),
		Number:                  number,
		Name:                    name,
		CampusUUID:              campusEntity.CampusUUID,
		BuildingUUID:            buildingEntity.BuildingUUID,
		Floor:                   floor,
		Type:                    classroomTypeEntity.ClassTypeUUID,
		Tag:                     tagJsonString,
		Capacity:                capacity,
		ExaminationRoom:         examinationRoom,
		ExaminationRoomCapacity: examinationRoomCapacity,
		IsMultimedia:            isMultimedia,
		IsAirConditioned:        isAirConditioned,
		Status:                  status,
		Description:             description,
		ManagementDepartment:    departmentUUID,
		Area:                    area,
		TablesChairsType:        tablesChairsTypeUUID,
	}
	tx := db.database.Create(&classroom)
	if tx.Error != nil {
		panic("初始化教室数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 教室表 [%s] 成功\n", classroom.Name)
	}
}
