package statistics

import (
	"fmt"
	"workProject/widget/utils"
)

type device struct {
    id  int64
}

// 查询单个记录
func DeviceQueryOne(device_title string) (int64) {
	db := utils.MysqlDb

    var u1 device //用来接收查询结果
    // 1. 写查询单条记录的sql语句
    sqlStr := `select id from device where title=?;` //？占位 下面的id
    // 2. 执行并拿到结果
    // 必须对rowObj对象调用Scan方法,因为该方法会释放数据库链接 // 从连接池里拿一个连接出来去数据库查询单条记录
    db.QueryRow(sqlStr, device_title).Scan(&u1.id) //&u1.id初始化u1结构体对象（变量）
    //row一行
 
    return u1.id
}


// 插入数据
func DeviceInsert(device_title string)  (id int64) {
	db := utils.MysqlDb
    // 1. 写SQL语句
    sqlStr := `insert into device(title) values("` + device_title + `")`
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