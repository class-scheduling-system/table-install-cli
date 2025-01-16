package database

import (
	"fmt"
	"frontleaves-table-install-cli/entity"
	"frontleaves-table-install-cli/utils"
)

// InitSystemData 初始化系统数据
func (db *DbOperate) InitSystemData(key string, value any) {
	var system = entity.CsSystem{
		SystemUUID: utils.GenerateUUIDNoDash(),
		SystemKey:  key,
		SystemVal:  value.(string),
	}
	tx := db.database.Create(&system)
	if tx.Error != nil {
		panic("初始化系统数据失败: " + tx.Error.Error())
	} else {
		fmt.Printf("初始化 系统表[%s] 数据 (%s) 成功\n", key, value)
	}
}
