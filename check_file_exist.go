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
