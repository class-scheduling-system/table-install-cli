package main

import (
	"embed"
	"frontleaves-table-install-cli/services"
	"github.com/pelletier/go-toml"
	"os"
)

func Cli(config *toml.Tree, resourcesFile embed.FS) {
	print("\n")
	args := os.Args

	if args == nil || len(args) == 1 || args[1] == "help" {
		println("智课方舟 | 基于AI的大学智慧排课系统")
		println("数据库快捷操作 CLI，使用帮助：")
		println("\thelp\t\t显示帮助信息")
		println("\tconn\t\t检查数据库链接是否正常")
		println("\tcreate\t\t创建数据表")
		println("\tinit\t\t初始化数据库(含创建数据表和初始化数据)")
		println("\tclear\t\t清空数据表数据(保留数据表结构)")
		println("\tdelete\t\t删除数据表")
		println("\treset\t\t重置数据库(删除当前数据表再重新创建后初始化数据)")
		print("\n")
		println("Copyright © 2022-2025 锋楪技术团队. All rights reserved.")
		return
	}

	switch args[1] {
	case "conn":
		services.ConnCheck(config)
	case "create":
		services.CreateDatabase(config, resourcesFile)
	case "init":
		services.InitDatabase(config, resourcesFile)
	case "clear":
		services.ClearDatabase(config)
	case "delete":
		services.DeleteDatabase(config)
	case "reset":
		services.ResetDatabase(config, resourcesFile)
	default:
		println("未知命令，请使用 help 查看帮助信息")
	}
}
