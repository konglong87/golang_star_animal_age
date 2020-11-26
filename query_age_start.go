package main

import (
	"fmt"
	"log"
	"time"
)

var TraditionAnimals = []string{"鼠","牛","虎","兔","龙","蛇","马","羊","猴", "鸡","狗","猪"}
func GetYearMonthDayFromStrDate(date string) (year, month, day int) {
	var shortForm = "2006-01-02"
	d, err := time.Parse(shortForm, date)
	if err != nil {
		log.Println("解析错误！")
		return 0, 0, 0
	}
	year = d.Year()
	month = int(d.Month())
	day = d.Day()
	return
}

func GetMyAnimal(year int)(animal string){
	if year <= 0 {
		animal = "-1"
	}

	flag := 1900
	circle := 12
	mod := (year-flag)%circle
	animal = TraditionAnimals[mod]
	return
}

func GetAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}


var startStr     = []string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"} //因为摩羯座，跨年，可以想象成一个圆圈，这个节点特殊处理，就是12-22之后也可以，或者1-19之前也可
var everyStarDay = []int{21, 20, 21, 21, 22, 22, 23, 24, 24, 24, 23, 22}
func GetMyStart2(month, day int) (star string){

	//看 这天 是在这个星座之前还是之后
	if day < everyStarDay[month] {
		return startStr[month-1]
	}
	return startStr[month]

}

func main() {
	y, m, d := GetYearMonthDayFromStrDate("1999-10-27")
	fmt.Println("[年龄]",GetAge(y))
	fmt.Println("[星座]",GetMyStart2(m, d))
	fmt.Println("[属相]",GetMyAnimal(y))
}




