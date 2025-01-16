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
	"fmt"
	"log"
	"os"
)

func CheckFileExist(resourcesFile embed.FS) {
	// 配置文件名
	configFileName := "configuration.toml"

	// 检查当前文件夹是否存在 configuration.toml
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		fmt.Printf("配置文件 %s 不存在，正在从默认路径拷贝...\n", configFileName)

		// 从嵌入资源中读取默认配置文件
		file, err := resourcesFile.ReadFile("resources/templates/configuration.toml")
		if err != nil {
			log.Fatal("读取嵌入文件失败: ", err)
		}

		// 创建新文件
		create, err := os.Create(configFileName)
		if err != nil {
			log.Fatal("创建配置文件失败: ", err)
		}
		defer func(create *os.File) {
			err := create.Close()
			if err != nil {
				log.Fatal("关闭文件失败: ", err)
			}
		}(create)

		// 写入内容到文件
		if _, err := create.Write(file); err != nil {
			log.Fatal("写入配置文件失败: ", err)
		}

		fmt.Printf("配置文件 %s 已成功创建。\n", configFileName)
	}
}
