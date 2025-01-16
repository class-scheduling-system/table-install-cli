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
	dsn := user + ":" + pass + "@tcp(" + host + ":" + strconv.FormatInt(port, 10) + ")/" + name + "?charset=utf8mb4"

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
