package statistics

import (
	"fmt"
	"workProject/widget/utils"
)

type Options struct {
	// id  int `db:"id"`
	// is_pass int `db:"is_pass"`
	version string `db:"version"`
	device string `db:"device"`
	country string `db:"country"`
	param string `db:"param"`
	vcname string `db:"vcname"`
}

//查询多行
func OptionQueryMulti(days string) (res []interface{}) {

	options := new(Options)
	db := utils.MysqlDb
	table := "operation_flow_" + days + "_Widgett";	

	sql := "select version,device,country,param,vcname from `" + table + "` where version != 90909 and vcname in ('widgetShow','widgetAdd','widgetClick','widgetCreate') and param like " + "'%" + "designName%'"
	rows,err:= db.Query(sql)
	
	defer func() {
	
		if rows != nil {
			rows.Close()
		}
	
	}()
	
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return 
	}
	
	// rows.Next(),用于循环获取所有数据
	for rows.Next() {
		// err = rows.Scan(&options.id,&options.is_pass)
		err = rows.Scan(&options.version,&options.device,&options.country,&options.param,&options.vcname)

		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return 
		}

		res = append(res, *options)
	}

	rows.Close()

	return res
}
