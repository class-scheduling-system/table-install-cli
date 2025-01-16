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
		println("\n\tUsage: ./<program> [command]\n")
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
