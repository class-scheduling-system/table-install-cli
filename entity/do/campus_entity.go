/*
 * --------------------------------------------------------------------------------
 * Copyright (c) 2016-NOW(至今) 筱锋
 * Author: 筱锋(https://www.x-lf.com)
 *
 * 本文件包含 XiaoMain 的源代码，该项目的所有源代码均遵循MIT开源许可证协议。
 * --------------------------------------------------------------------------------
 * 许可证声明：
 *
 * 版权所有 (c) 2016-2024 筱锋。保留所有权利。
 *
 * 本软件是“按原样”提供的，没有任何形式的明示或暗示的保证，包括但不限于
 * 对适销性、特定用途的适用性和非侵权性的暗示保证。在任何情况下，
 * 作者或版权持有人均不承担因软件或软件的使用或其他交易而产生的、
 * 由此引起的或以任何方式与此软件有关的任何索赔、损害或其他责任。
 *
 * 使用本软件即表示您了解此声明并同意其条款。
 *
 * 有关MIT许可证的更多信息，请查看项目根目录下的LICENSE文件或访问：
 * https://opensource.org/licenses/MIT
 * --------------------------------------------------------------------------------
 * 免责声明：
 *
 * 使用本软件的风险由用户自担。作者或版权持有人在法律允许的最大范围内，
 * 对因使用本软件内容而导致的任何直接或间接的损失不承担任何责任。
 * --------------------------------------------------------------------------------
 */

package do

import "time"

// CsCampus 表示校区信息的实体结构
type CsCampus struct {
	CampusUUID    string     `gorm:"column:campus_uuid;primaryKey;type:char(32);not null" json:"campus_uuid"`                                           // 校区主键
	CampusName    string     `gorm:"column:campus_name;unique;type:varchar(32);not null" json:"campus_name"`                                            // 校区名称
	CampusCode    string     `gorm:"column:campus_code;unique;type:varchar(32);not null" json:"campus_code"`                                            // 校区编码
	CampusDesc    string     `gorm:"column:campus_desc;type:varchar(255)" json:"campus_desc"`                                                           // 校区描述
	CampusStatus  bool       `gorm:"column:campus_status;type:boolean;default:true;not null" json:"campus_status"`                                      // 校区状态
	CampusAddress string     `gorm:"column:campus_address;type:varchar(255);not null" json:"campus_address"`                                            // 校区地址
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
}

// TableName 返回 CsCampus 对应的数据库表名
func (CsCampus) TableName() string {
	return "cs_campus"
}
