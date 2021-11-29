package userActivity

import (
	"fmt"
	"workProject/widget/model/statistics"
	"reflect"
	"encoding/json"
	"strings"
	"strconv"
)
var date string

var originalData []interface{};

var handleData map[string]map[string]int


func Run(days string) {
	fmt.Println(days)
	// getOriginalData(days);
	// getHandleData();
	// insertData();
}

func getOriginalData(days string){
	date = days
	originalData = statistics.OptionQueryMulti(days)
}

func getHandleData(){
	handleData = make(map[string]map[string]int);

	for i := 0; i < len(originalData); i++ {
		item   := reflect.ValueOf(originalData[i])
		param  := item.FieldByName("param")
		vcname := fmt.Sprintf("%s",item.FieldByName("vcname"))

		//取字符串json里的数据  先转成map 再取出来
		var p map[string]interface{}
		json.Unmarshal([]byte(fmt.Sprintf("%s",param)), &p)

		keystr := fmt.Sprintf("%s^%s^%s^%s",p["designName"],item.FieldByName("country"),item.FieldByName("device"),item.FieldByName("version"))

		if _, ok := handleData[keystr]; !ok{
			//新增一个
			oneData := make(map[string]int)//一维数组重置
			oneData["widgetClick"]  = 0
			oneData["widgetAdd"]    = 0
			oneData["widgetCreate"] = 0
			oneData["widgetShow"]   = 0
			handleData[keystr] = oneData
		}
		// 看是要增加哪个
		handleData[keystr][vcname] += 1
	}
}

func insertData(){
	for index,item := range handleData{
		indexArr := strings.Split(index,"^")

		widgetId := getWidgetId(indexArr[0])
		if widgetId <= 0{
			continue
		}

		/*
			version,show,click,add,create 转int
			country 看是否在表里 不在就更新拿id 在的话直接拿id
			device  看是否在表里 不在就更新拿id 在的话直接拿id
		*/
		insertWidgetS := make(map[string]interface{});
		insertWidgetS["widgetId"]  = widgetId;
		insertWidgetS["add"]       = int64(item["widgetAdd"]);
		insertWidgetS["show"]      = int64(item["widgetShow"]);
		insertWidgetS["create"]    = int64(item["widgetCreate"]);
		insertWidgetS["click"]     = int64(item["widgetClick"]);
		insertWidgetS["countryId"] = getCountryId(indexArr[1]);
		insertWidgetS["deviceId"]  = getDeviceId(indexArr[2]);
		insertWidgetS["version"],_ = strconv.ParseInt(indexArr[3], 10, 64)


		insertWidgetS["created_at"] = fmt.Sprintf("%s",date)
		insertWidgetS["updated_at"] = fmt.Sprintf("%s",date)
		fmt.Println(insertWidgetS)

		/*
		 	插到表里去
		*/
		// res := statistics.WidgetStatisInsert(insertWidgetS)
		fmt.Printf("insert res :%v",res)
	}
}

func getWidgetId(design string) (id int64) {
	id = statistics.WidgetQueryOne(design)
	return id
}

func getDeviceId(title string) (id int64) {
	id = statistics.DeviceQueryOne(title)
	if id <= 0 {
	  id = statistics.DeviceInsert(title)
	}
	return id
}

func getCountryId(title string) (id int64) {
	id = statistics.CountryQueryOne(title)
	if id <= 0 {
	  id = statistics.CountryInsert(title)
	}
	return id
}