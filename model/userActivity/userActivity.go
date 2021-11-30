package userActivity

import (
	"fmt"
	"strings"
	"workProject/widget/utils"
)


// 插入数据
func UserActivityInsert(insertArr map[string]interface{})  (id int64) {
	db := utils.MysqlDb
    // 1. 写SQL语句

	var keysArr []string
	var valsArr []string
	for k, v := range insertArr {
		keysArr = append(keysArr, "`" + k + "`")
		if k != "created_at" && k != "updated_at" && k != "udid"{
			valsArr = append(valsArr, fmt.Sprintf("'%d'",v))
		}else if k == "udid"{
			valsArr = append(valsArr, fmt.Sprintf("'%s'",v))
		}else{
			valsArr = append(valsArr, fmt.Sprintf("'%v'",v))
		}
	}
    sqlStr := `insert into userActivityUdid1(` + strings.Join(keysArr, ", ") + `) values(` + strings.Join(valsArr, ",") + `)`
	fmt.Println(sqlStr)
    // 2. exec
    ret, err := db.Exec(sqlStr) //exec执行（Python中的exec就是执行字符串代码的，返回值是None，eval有返回值）
    if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
    }
    // 如果是插入数据的操作，能够拿到插入数据的id
    id, err = ret.LastInsertId()
    if err != nil {
        fmt.Printf("get id failed,err:%v\n", err)
        return
    }
    return id
}
