package utils

import (
	"embed"
	"github.com/google/uuid"
	"strings"
)

// GenerateCreateTableSQL 生成创建数据表的 SQL 语句
func GenerateCreateTableSQL(table string, resourcesFile embed.FS) []string {
	file, err := resourcesFile.ReadFile("resources/sql/" + table + ".sql")
	if err != nil {
		panic("读取嵌入文件失败: " + err.Error())
	}
	stringFile := string(file)
	strings.ReplaceAll(stringFile, "USE `class-scheduling-system`;", "")
	split := strings.Split(stringFile, ";")
	// 去除空白行
	for i := 0; i < len(split); i++ {
		split[i] = strings.TrimSpace(split[i])
		if split[i] == "" {
			split = append(split[:i], split[i+1:]...)
			i--
		}
	}
	return split
}

func GenerateUUIDNoDash() string {
	getUUID := uuid.New()
	return strings.ReplaceAll(getUUID.String(), "-", "")
}
