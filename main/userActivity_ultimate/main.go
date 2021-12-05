package main

import (
	"fmt"
	"time"
	"workProject/widget/utils"
	"workProject/widget/logic/userActivity_ultimate"
)
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
	// endTime   := "2021-10-15"
	// userActivity_ultimate.Run(endTime);
	all();
	// yesterday();
}

/*
	从8.5开始
	截止到昨天
*/
func all(){
	var cstSh, _ = time.LoadLocation("Asia/Shanghai");
	startTime := "2021-08-05";
	// startTime := "2021-09-13";

	endTime   := time.Now().In(cstSh).AddDate(0, 0, -1).Format("2006-01-02");
	// endTime   := "2021-09-16"
	days      := utils.GetDaysDiffer(startTime, endTime)//看差了几天
	var i int64
	iTime := startTime
	oneDay,_ := time.ParseDuration("24h")
	for i = 1; i <= days; i++ {
		day, _ := time.ParseInLocation("2006-01-02", iTime,time.Local)
	
		iTime := day.Add(oneDay).Format("2006-01-02")
		fmt.Println(iTime)	
		userActivity_ultimate.Run(iTime)
		// fmt.Printf("i type:%t,oneDay:%t",i,oneDay)
		oo,_ := time.ParseDuration("24h")
		oneDay += oo

	}
}

func yesterday(){
	var cstSh, _ = time.LoadLocation("Asia/Shanghai");
	Time   := time.Now().In(cstSh).AddDate(0, 0, -1).Format("2006-01-02");
		// endTime   := "2021-10-15"
	userActivity_ultimate.Run(Time)
}