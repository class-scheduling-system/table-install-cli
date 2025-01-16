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

package database

import (
	"github.com/pelletier/go-toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

// NewDatabaseOperate 创建数据库操作对象
func NewDatabaseOperate(config *toml.Tree) *DbOperate {
	// 读取数据库配置
	host := config.Get("database.host").(string)
	port := config.Get("database.port").(int64)
	name := config.Get("database.name").(string)
	user := config.Get("database.user").(string)
	pass := config.Get("database.password").(string)

	// 数据库连接字符串
	dsn := user + ":" + pass + "@tcp(" + host + ":" + strconv.FormatInt(port, 10) + ")/" + name + "?charset=utf8mb4&parseTime=true&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	return &DbOperate{
		config:   config,
		database: db,
	}
}

// DbOperate 数据库操作对象
type DbOperate struct {
	database *gorm.DB
	config   *toml.Tree
}
