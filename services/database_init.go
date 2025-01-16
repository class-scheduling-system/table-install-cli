package services

import (
	"embed"
	"github.com/pelletier/go-toml"
)

func InitDatabase(config *toml.Tree, resourcesFile embed.FS) {
	// 初始化数据库
	CreateDatabase(config, resourcesFile)
	// 初始化数据

}
