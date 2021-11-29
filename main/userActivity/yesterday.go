package main

import (
	"fmt"
	"time"
	"workProject/widget/utils"
	"workProject/widget/logic/userActivity"
)
/*
	userActivity (除去created_at,updated_at以外，其他的都是int类型)
		id,userCount,widgetId,countryId,version,deviceId,created_at,updated_at
*/
/*
   执行哪一天的数据
   根据 version != 90909 从表里拿数据
   遍历  分组 
		 key = designName^country^version^device^udid
		 val = userCount 统计数
	插入
		统计表
*/
func main(){
	yesterday()
}

func yesterday(){
	var cstSh, _ = time.LoadLocation("Asia/Shanghai");
	Time   := time.Now().In(cstSh).AddDate(0, 0, -1).Format("2006-01-02");
	userActivity.Run(Time)
}