package main

import (
	"fmt"
	"time"
)

func time1() {
	now := time.Now() // 获取本地的时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year() // 年
	month := now.Month() // 月
	day := now.Day() // 日
	hour := now.Hour() // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timestamp() {
	now := time.Now()
	timestamp1 := now.Unix() // 时间戳
	timestamp2 := now.UnixNano()
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func timestampToTime(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)	
	fmt.Println(timeObj)
	year := timeObj.Year()	// 年
	month := timeObj.Month() // 月
	day := timeObj.Day() // 日
	hour := timeObj.Hour() // 小时
	minute := timeObj.Minute() // 分钟
	second := timeObj.Second() //  秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func tick() {
	ticker := time.Tick(time.Second) 
	for i := range ticker {
		fmt.Println(i)
	}
}

func format() {
	now := time.Now()

	// 需要使用go的诞生时间2006年1月2号15点04分
	fmt.Println(now.Format("2006/01/02"))
	// 获取当前微秒
	fmt.Println(now.UnixNano() / 1000000)
}

func main() {
	time1()
	timestamp()
	timestampToTime(time.Now().Unix())

	now := time.Now()
	/*
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
	*/

	m1, _ := time.ParseDuration("1m")
	m1s := now.Add(m1)
	fmt.Println(m1s)
	/*
	h1, _ := time.ParseDuration("1h")
	h1s := now.Sub(h1)
	fmt.Println(h1s.Hours())
	*/

	subM := now.Sub(m1s)
	fmt.Println(subM.Minutes())

	if now.Equal(m1s) {
		fmt.Println("now and m1s is equal")
	}

	//tick()

	format()
}

