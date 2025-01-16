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

package do

import (
	"time"
)

// CsClassroom represents the classroom table in the database.
type CsClassroom struct {
	ClassroomUUID           string    `gorm:"column:classroom_uuid;primaryKey;type:char(32);not null;comment:教室主键"`
	Number                  string    `gorm:"column:number;type:varchar(32);unique;not null;comment:教室编号"`
	Name                    string    `gorm:"column:name;type:varchar(32);not null;comment:教室名称"`
	CampusUUID              string    `gorm:"column:campus_uuid;type:char(32);not null;comment:校区主键"`
	BuildingUUID            string    `gorm:"column:building_uuid;type:char(32);not null;comment:楼栋主键"`
	Floor                   uint      `gorm:"column:floor;type:int;not null;comment:楼层"`
	Type                    string    `gorm:"column:type;type:char(32);not null;comment:教室类型"`
	Tag                     *string   `gorm:"column:tag;type:json;comment:教室标签"`
	Capacity                uint      `gorm:"column:capacity;type:int;not null;comment:教室容量"`
	ExaminationRoom         bool      `gorm:"column:examination_room;type:tinyint(1);default:0;not null;comment:是否是考场"`
	ExaminationRoomCapacity *uint     `gorm:"column:examination_room_capacity;type:int;comment:考场容量"`
	IsMultimedia            bool      `gorm:"column:is_multimedia;type:tinyint(1);default:0;not null;comment:是否是多媒体教室"`
	IsAirConditioned        bool      `gorm:"column:is_air_conditioned;type:tinyint(1);default:0;not null;comment:是否有空调"`
	Status                  bool      `gorm:"column:status;type:tinyint(1);default:1;not null;comment:教室状态 0:禁用 1:启用"`
	Description             *string   `gorm:"column:description;type:varchar(255);comment:教室描述"`
	ManagementDepartment    string    `gorm:"column:management_department;type:char(32);not null;comment:管理部门"`
	Area                    float64   `gorm:"column:area;type:decimal(10,2);not null;comment:教室面积"`
	TablesChairsType        string    `gorm:"column:tables_chairs_type;type:char(32);not null;comment:桌椅类型"`
	CreatedAt               time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdatedAt               time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;not null;comment:更新时间"`
}

// TableName specifies the table name for the CsClassroom model.
func (CsClassroom) TableName() string {
	return "cs_classroom"
}
