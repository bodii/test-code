package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cinema struct {
	rw *sync.RWMutex
}

type Person struct {
	Name string
}

type Admin struct {
	Name string
}

func (p *Person) SitDown(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s 观众坐下\n", p.Name)
}

func (p *Person) StartWatch() {
	fmt.Printf("%s 开始观看\n", p.Name)
	i := time.Duration(rand.Intn(5) + 1)
	time.Sleep(i * time.Second)
	fmt.Printf("%s 观看结束\n", p.Name)
}

func (p *Person) WatchMovie(c *Cinema, wg *sync.WaitGroup, b chan int) {
	c.rw.RLock()
	p.SitDown(wg)
	<-b
	p.StartWatch()
	c.rw.RUnlock()
}

func (a *Admin) ChangeMovie(c *Cinema, wg *sync.WaitGroup) {
	defer wg.Done()

	c.rw.Lock()

	// for !c.rw.TryLock() {
	// 	fmt.Println("还有人看电影")
	// 	time.Sleep(1 * time.Second)
	// }

	fmt.Printf("所有观众看完电影，%s管理员切换影片\n", a.Name)
	c.rw.Unlock()
}

func example1() {
	c := &Cinema{
		rw: &sync.RWMutex{},
	}

	a := &Admin{
		Name: "影院",
	}

	b := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int, c *Cinema, wg *sync.WaitGroup, b chan int) {
			p := Person{
				Name: fmt.Sprintf("%d 号观众", i),
			}
			p.WatchMovie(c, wg, b)
		}(i, c, &wg, b)
	}

	wg.Wait()
	fmt.Println("所有观众入座，开始播放电影")
	close(b)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		a.ChangeMovie(c, wg)
	}(&wg)
	wg.Wait()
}

type Desk struct {
	rw     *sync.RWMutex
	Snacks int
}

func (d *Desk) PlaceSnacks(wg *sync.WaitGroup) {
	defer wg.Done()

	d.rw.Lock()
	d.Snacks = 10
	time.Sleep(3 * time.Second)
	fmt.Println("零食准备完毕...")
	d.rw.Unlock()
}

func (p *Person) GetSnacks(d *Desk, wg *sync.WaitGroup) {
	for !d.rw.TryRLock() {
		fmt.Printf("桌子上没有零食， %s望眼欲穿\n", p.Name)
		time.Sleep(2 * time.Second)
	}

	d.rw.RUnlock()

	defer wg.Done()
	defer d.rw.Unlock()

	d.rw.Lock()
	if d.Snacks > 0 {
		fmt.Printf("%s抢到零售，开心！\n", p.Name)
		d.Snacks--
		return
	}
	fmt.Printf("%s没有抢到零食，难受！\n", p.Name)
}

func example2() {
	d := Desk{
		rw: &sync.RWMutex{},
	}
	wg := sync.WaitGroup{}
	wg.Add(12)

	go d.PlaceSnacks(&wg)
	for i := 0; i < 11; i++ {
		p := Person{
			Name: fmt.Sprintf("%d 号猿猴", i),
		}

		go p.GetSnacks(&d, &wg)
	}

	wg.Wait()
}

func main() {
	// example1()
	example2()
}
