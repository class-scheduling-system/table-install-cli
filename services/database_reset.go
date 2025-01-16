package services

import (
	"embed"
	"github.com/pelletier/go-toml"
)

func ResetDatabase(config *toml.Tree, resourcesFile embed.FS) {
	// 重置数据库
	DeleteDatabase(config)
	// 初始化数据库
	CreateDatabase(config, resourcesFile)
	println("数据库重置成功")
}
