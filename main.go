package main

import (
	"embed"
	"fmt"
	"github.com/pelletier/go-toml"
)

//go:embed resources/*
var resourcesFile embed.FS

func main() {
	CheckFileExist(resourcesFile)
	// 读取 TOML 文件
	config, err := toml.LoadFile("configuration.toml")
	if err != nil {
		fmt.Println("Error reading TOML file:", err)
		return
	}
	Cli(config, resourcesFile)
}
