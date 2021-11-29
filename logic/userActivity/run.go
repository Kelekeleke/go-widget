package userActivity

import (
	"fmt"
	"workProject/widget/model/userActivity"
	"reflect"
	"encoding/json"
	"strings"
	"strconv"
)
var date string

var originalData []interface{};

var handleData map[string]map[string]int
// var handleData map[string]map[int]string

func Run(days string) {
	getOriginalData(days);
	getHandleData();
	insertData();
}

func getOriginalData(days string){
	date = days
	originalData = userActivity.OptionQueryMulti(days)
}

func getHandleData(){
	handleData = make(map[string]map[string]int);
	// handleData = make(map[string]map[int]string);


	for i := 0; i < len(originalData); i++ {
		item   := reflect.ValueOf(originalData[i])
		param  := item.FieldByName("param")
		udid   := fmt.Sprintf("%s",item.FieldByName("udid"))

		//取字符串json里的数据  先转成map 再取出来
		var p map[string]interface{}
		json.Unmarshal([]byte(fmt.Sprintf("%s",param)), &p)

		keystr := fmt.Sprintf("%s^%s^%s^%s",p["designName"],item.FieldByName("country"),item.FieldByName("device"),item.FieldByName("version"))

		if _, ok := handleData[keystr]; !ok{
			//在这个条件下 这个人出现过没有 如果没有把这个人加上
			if _,okk := handleData[keystr][udid]; !okk {
				oneData := make(map[string]int);
				oneData[udid]  = 0
				handleData[keystr] = oneData
			}
	
		}else{
			appendData := make(map[string]int);
			appendData = handleData[keystr]
			appendData[udid]  = 0
			handleData[keystr] = appendData
		}
		// 看这个人出现了几次
		handleData[keystr][udid] += 1
	}
}

func insertData(){
	for index,item := range handleData{
		indexArr := strings.Split(index,"^")

		widgetId := getWidgetId(indexArr[0])
		if widgetId <= 0{
			continue
		}
		// if indexArr[0] == "Colorful_22_S"{
		// 	fmt.Printf("colorful_22_S len:%b",len(item))
		// }
		// if len(item) > 1{
		// 	fmt.Println(len(item))
		// }

		/*
			version,show,click,add,create 转int
			country 看是否在表里 不在就更新拿id 在的话直接拿id
			device  看是否在表里 不在就更新拿id 在的话直接拿id
		*/
		insertWidgetS := make(map[string]interface{});
		insertWidgetS["widgetId"]  = widgetId;
		insertWidgetS["userCount"] = int64(len(item));
		insertWidgetS["countryId"] = getCountryId(indexArr[1]);
		insertWidgetS["deviceId"]  = getDeviceId(indexArr[2]);
		insertWidgetS["version"],_ = strconv.ParseInt(indexArr[3], 10, 64)
		insertWidgetS["created_at"] = fmt.Sprintf("%s",date)
		insertWidgetS["updated_at"] = fmt.Sprintf("%s",date)
		fmt.Println(insertWidgetS)

		/*
		 	插到表里去
		*/
		res := userActivity.UserActivityInsert(insertWidgetS)
		fmt.Printf("insert res :%v",res)
	}
}

func getWidgetId(design string) (id int64) {
	id = userActivity.WidgetQueryOne(design)
	return id
}

func getDeviceId(title string) (id int64) {
	id = userActivity.DeviceQueryOne(title)
	if id <= 0 {
	  id = userActivity.DeviceInsert(title)
	}
	return id
}

func getCountryId(title string) (id int64) {
	id = userActivity.CountryQueryOne(title)
	if id <= 0 {
	  id = userActivity.CountryInsert(title)
	}
	return id
}