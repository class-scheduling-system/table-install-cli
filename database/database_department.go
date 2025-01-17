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
	"fmt"
	"frontleaves-table-install-cli/models/do"
	"frontleaves-table-install-cli/utils"
	"math/rand"
	"time"
)

// InitUnitCategoryData 初始化单位分类数据
func (db *DbOperate) InitUnitCategoryData(name string, englishName, shortName *string) {
	var unitCategory = do.CsUnitCategory{
		UnitCategoryUUID: utils.GenerateUUIDNoDash(),
		Name:             name,
		Order:            rand.Intn(100),
		EnglishName:      englishName,
		ShortName:        shortName,
		IsEntity:         false,
	}
	tx := db.database.Create(&unitCategory)
	if tx.Error != nil {
		panic("初始化单位分类数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 单位分类表 [%s] 成功\n", unitCategory.Name)
	}
}

// InitUnitTypeData 初始化单位类型数据
func (db *DbOperate) InitUnitTypeData(name string, englishName, shortName *string) {
	var unitType = do.CsUnitType{
		UnitTypeUUID: utils.GenerateUUIDNoDash(),
		Name:         name,
		EnglishName:  englishName,
		ShortName:    shortName,
		Order:        rand.Intn(100),
	}
	tx := db.database.Create(&unitType)
	if tx.Error != nil {
		panic("初始化单位类型数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 单位类型表 [%s] 成功\n", unitType.Name)
	}
}

// InitDepartmentData 初始化部门数据
func (db *DbOperate) InitDepartmentData(code, name, unitCategory, unitType string, englishName, shortName, address, administrativeHead, partyCommitteeHead, parentDepartment, assignedTeachingBuilding, fixedPhone, remark *string, establishmentDate time.Time, expirationDate *time.Time, isEntity, isTeachingCollege, isAttendingCollege, isTeachingOffice, isEnabled bool) {
	var unitCategoryEntity = do.CsUnitCategory{}
	db.database.Where("name = ?", unitCategory).First(&unitCategoryEntity)
	var unitTypeEntity = do.CsUnitType{}
	db.database.Where("name = ?", unitType).First(&unitTypeEntity)
	var parentDepartmentEntity = do.CsDepartment{}
	if parentDepartment != nil {
		db.database.Where("department_name = ?", *parentDepartment).First(&parentDepartmentEntity)
	}
	var buildingEntity = do.CsBuilding{}
	if assignedTeachingBuilding != nil {
		db.database.Where("building_name = ?", *assignedTeachingBuilding).First(&buildingEntity)
	}

	var parentDepartmentUUID *string = nil
	if parentDepartment != nil {
		parentDepartmentUUID = &parentDepartmentEntity.DepartmentUUID
	}

	var buildingUUID *string = nil
	if assignedTeachingBuilding != nil {
		buildingUUID = &buildingEntity.BuildingUUID
	}

	var department = do.CsDepartment{
		DepartmentUUID:           utils.GenerateUUIDNoDash(),
		DepartmentCode:           code,
		DepartmentName:           name,
		DepartmentOrder:          rand.Intn(100),
		DepartmentEnglishName:    englishName,
		DepartmentShortName:      shortName,
		DepartmentAddress:        address,
		IsEntity:                 isEntity,
		AdministrativeHead:       administrativeHead,
		PartyCommitteeHead:       partyCommitteeHead,
		EstablishmentDate:        establishmentDate,
		ExpirationDate:           expirationDate,
		UnitCategory:             unitCategoryEntity.UnitCategoryUUID,
		UnitType:                 unitTypeEntity.UnitTypeUUID,
		ParentDepartment:         parentDepartmentUUID,
		AssignedTeachingBuilding: buildingUUID,
		IsTeachingCollege:        isTeachingCollege,
		IsAttendingCollege:       isAttendingCollege,
		FixedPhone:               fixedPhone,
		Remark:                   remark,
		IsTeachingOffice:         isTeachingOffice,
		IsEnabled:                isEnabled,
	}
	tx := db.database.Create(&department)
	if tx.Error != nil {
		panic("初始化部门数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 部门表 [%s] 成功\n", department.DepartmentName)
	}
}
