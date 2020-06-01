package main

import (
	"fmt"
	"time"
)

func basic() {
	//time.Time对象的时区为CST(本地时区)
	now := time.Now()
	fmt.Printf("%#v\n", now)

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println(now.Date())

}

func timestamp() {
	now := time.Now()
	fmt.Println(now.Unix())     //返回GMT时间戳(second)
	fmt.Println(now.UnixNano()) //返回GMT时间戳(Nanosecond)

	//now.Add(nanosecond Duration) func (t Time)Add(nanosecond Duration)
	now = now.Add(20 * time.Millisecond)
	now = now.Add(99 * time.Nanosecond)
	fmt.Println(now.UnixNano())

	//将时间戳转化为time类型time.Unix(second int64, nanosecond int64)
	newNow := time.Unix(now.Unix(), 0)
	fmt.Println(newNow.UnixNano())

}

func operator() {
	now1 := time.Now()
	time.Sleep(2 * time.Second) // time.Nanosecond
	now2 := time.Now()
	fmt.Println(now2.UnixNano())
	difference := now2.Sub(now1) //type Duration int64 //Duration自定义类型实现了一些自己的方法, 例如打印
	fmt.Printf("Nanosecond:%v\n", difference)

	now3 := time.Now()
	//fmt.Println("hello, world")
	now4 := time.Now()

	// now3 == now4
	if now3.Equal(now4) {
		fmt.Println("quick")
	}

	// now3 < now4
	if now3.Before(now4) {
		fmt.Println("now4 before now5 !")
	}

	//now3 > now4
	if now3.After(now4) {
		fmt.Println("now4 after now5 ?")
	}

}

func timerTick() {
	timer := time.Tick(2 * time.Second) //time.Nanosecond
	for now := range timer {
		fmt.Println(now)
	}
}

func format() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05.000000000 PM"))
}

func parse() {

	//默认将字符串作为UTC时区的时间
	if retTime, err := time.Parse("2006/01/02 15:04:05 PM", "2020/02/09 13:43:00 PM"); err == nil {
		fmt.Println(retTime)
	}

	//获取当前时间, CST tim对象
	fmt.Println(time.Now())
	//获取当前时间,对应其他时区的time对象 //时间没有发生改变, 只是显示基准发生改变
	fmt.Println(time.Now().UTC())

	//以指定时区为基准解析字符串, return time.Time对象
	//if loc, err := time.LoadLocation("EST"); err == nil {//以EST时区为基准解析字符串
	if loc, err := time.LoadLocation("Asia/Shanghai"); err == nil { //以CST时区为基准解析字符串
		if retTime, err := time.ParseInLocation("2006/01/02 15:04:05 PM", "2020/02/09 13:43:00 PM", loc); err == nil {
			fmt.Println(retTime)
			locTime := time.Unix(retTime.Unix(), 0)
			fmt.Println(locTime)
		}
	}
}

func test() {
	//获取当前时间, 时区默认为系统时区
	nowTime := time.Now()
	fmt.Println(nowTime)
	//默认将字符串作为UTC时区的时间
	if retTime, err := time.Parse("2006/01/01 15:04:05", "2020/02/09 15:10:00"); err == nil {
		fmt.Println(retTime)
	}
	if loc, err := time.LoadLocation("Asia/Shanghai"); err == nil {
		if locTime, err := time.ParseInLocation("2006/01/02 15:04:05", "2020/03/05 15:10:00", loc); err == nil {
			fmt.Println(locTime)
			fmt.Println(locTime.Sub(nowTime))
		}
	}

}

func main() {
	fmt.Println("hello, loop")

	//timestamp()
	//operator()
	//timerTick()
	//parse()
	test()
}
