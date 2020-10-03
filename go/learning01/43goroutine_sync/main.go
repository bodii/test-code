package main

import (
	"fmt"
	"image"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	wg.Done()
	fmt.Println("hello goroutine")
}

// sync包中提供了一个针对只执行一次场景的解决方案sync.Once

var icons map[string]image.Image
var loadIconOnce sync.Once

func loadIcon(icon string) image.Image {
	return nil
}

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon icon
func Icon(name string) image.Image {
	loadIconOnce.Do(loadIcons)
	return icons[name]
}

/*
	并发安全的单例模式
	sync.Once 其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，
	而布尔值用来记录初始化是否完成。
*/
type singleton struct{}

var instance *singleton
var once sync.Once

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	wg.Add(1)
	go hello()

	fmt.Println("main goroutine done!")
	wg.Wait()

	fmt.Printf("%#v %p\n", getInstance(), getInstance())
	fmt.Printf("%#v %p\n", getInstance(), getInstance())
}
