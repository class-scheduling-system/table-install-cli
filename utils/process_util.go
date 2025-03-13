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

package utils

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// GenerateCreateTableSQL
//
// 该函数用于生成创建数据表的 SQL 语句，通过读取嵌入文件系统中的 SQL 文件来生成 SQL 语句。
//
// 用法示例：
//
//	GenerateCreateTableSQL("cs_system", resourcesFile)
//
// 参数说明：
//
//	table         string        数据表名称
//	resourcesFile embed.FS      嵌入文件系统
//
// 返回值：
//
//	[]string 生成的 SQL 语句
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

// GenerateUUIDNoDash
//
// 该函数用于生成不带破折号的 UUID, 通过调用第三方库生成 UUID 并去除其中的破折号。
//
// 用法示例：
//
//	GenerateUUIDNoDash()
//
// 返回值：
//
//	string 生成的 UUID
func GenerateUUIDNoDash() string {
	getUUID := uuid.New()
	return strings.ReplaceAll(getUUID.String(), "-", "")
}

// Ptr
//
// 该函数用于将字符串转换为指针类型。
//
// 用法示例：
//
//	Ptr("hello")
//
// 参数说明：
//
//	value string 需要转换的字符串
//
// 返回值：
//
//	*string 转换后的指针
func Ptr(value string) *string {
	if value == "" || value == "null" || value == "NULL" || value == "Null" || value == "nil" {
		return nil
	} else {
		return &value
	}
}

// Encrypt 对密码进行加密处理
// 1. 先使用 SHA-256 进行哈希
// 2. 再使用 BCrypt 加盐哈希
func Encrypt(password string) string {
	// 先进行 SHA-256 哈希
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte(password))
	sha256Hash := hex.EncodeToString(sha256Hasher.Sum(nil))

	// 使用 BCrypt 加盐哈希
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(sha256Hash), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(bcryptHash)
}
