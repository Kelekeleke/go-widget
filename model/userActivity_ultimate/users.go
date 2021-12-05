package userActivity_ultimate

import (
	"workProject/widget/utils"
)

type users struct {
    id  int64
    created_at string
}

// 查询单个记录
func UsersQueryOne(design string) (int64,string) {
	db := utils.MysqlDb

    var u1 users //用来接收查询结果
    // 1. 写查询单条记录的sql语句
    sqlStr := `select id,DATE_FORMAT(created_at,'%Y-%m-%d') from users where udid=?;` //？占位 下面的id
    // 2. 执行并拿到结果
    // 必须对rowObj对象调用Scan方法,因为该方法会释放数据库链接 // 从连接池里拿一个连接出来去数据库查询单条记录
    db.QueryRow(sqlStr, design).Scan(&u1.id,&u1.created_at) //&u1.id初始化u1结构体对象（变量）
    //row一行
 
    return u1.id,u1.created_at
}