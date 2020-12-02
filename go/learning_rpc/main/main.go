package main

import (
	"context"
	"encoding/json"
	"fmt"
	"learning_rpc"
	"learning_rpc/codec"
	"learning_rpc/xclient"
	"log"
	"net"
	"sync"
	"time"
)

// Foo ...
type Foo int

// Args ...
type Args struct{ Num1, Num2 int }

// Sum ...
func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

// func startServer(addr chan string) {
// 	var foo Foo
// 	if err := learning_rpc.Register(&foo); err != nil {
// 		log.Fatal("register error:", err)
// 	}
// 	l, err := net.Listen("tcp", ":0")
// 	if err != nil {
// 		log.Fatal("network error:", err)
// 	}
// 	log.Println("start rpc server on", l.Addr())
// 	addr <- l.Addr().String()
// 	learning_rpc.Accept(l)
// }

func serverTest() {
	addr := make(chan string)
	go startServer(addr)

	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)

	_ = json.NewEncoder(conn).Encode(learning_rpc.DefaultOption)
	cc := codec.NewGobCodec(conn)

	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("learning_rpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}

// func clientTest() {
// 	log.SetFlags(0)
// 	addr := make(chan string)
// 	go startServer(addr)
// 	client, _ := learning_rpc.Dial("tcp", <-addr)
// 	defer func() { _ = client.Close() }()

// 	time.Sleep(time.Second)

// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			args := fmt.Sprintf("learning_rpc req %d", i)
// 			var reply string
// 			if err := client.Call("Foo.Sum", args, &reply); err != nil {
// 				log.Fatal("call Foo.Sum error:", err)
// 			}
// 			log.Println("reply:", reply)
// 		}(i)
// 	}
// 	wg.Wait()
// }

// func getMethodRecord() {
// 	var wg sync.WaitGroup
// 	typ := reflect.TypeOf(&wg)
// 	for i := 0; i < typ.NumMethod(); i++ {
// 		method := typ.Method(i)
// 		argv := make([]string, 0, method.Type.NumIn())
// 		returns := make([]string, 0, method.Type.NumOut())
// 		// start from j, the 0th input is wg itself
// 		for j := 1; j < method.Type.NumIn(); j++ {
// 			argv = append(argv, method.Type.In(j).Name())
// 		}
// 		for j := 0; j < method.Type.NumOut(); j++ {
// 			returns = append(returns, method.Type.Out(j).Name())
// 		}
// 		fmt.Printf("func (w *%s) %s(%s) %s\n",
// 			typ.Elem().Name(),
// 			method.Name,
// 			strings.Join(argv, ","),
// 			strings.Join(returns, ","))
// 	}
// }

// func sendRPCRequest() {
// 	log.SetFlags(0)
// 	addr := make(chan string)
// 	go startServer(addr)
// 	client, _ := learning_rpc.Dial("tcp", <-addr)
// 	defer func() { _ = client.Close() }()

// 	time.Sleep(time.Second)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			args := &Args{Num1: i, Num2: i * i}
// 			var reply int
// 			if err := client.Call("Foo.Sum", args, &reply); err != nil {
// 				log.Fatal("call Foo.Sum error:", err)
// 			}
// 			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
// 		}(i)
// 	}
// 	wg.Wait()
// }

// func call(addrCh chan string) {
// 	client, _ := learning_rpc.DialHTTP("tcp", <-addrCh)
// 	defer func() { _ = client.Close() }()
// 	time.Sleep(time.Second)

// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			args := &Args{Num1: i, Num2: i * i}
// 			var reply int
// 			if err := client.Call(context.Background(), "Foo.Sum", args, &reply); err != nil {
// 				log.Fatal("call Foo.Sum error:", err)
// 			}
// 			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
// 		}(i)
// 	}
// 	wg.Wait()
// }

// func debug() {
// 	log.SetFlags(0)
// 	ch := make(chan string)
// 	go call(ch)
// 	startServer(ch)
// }

func (f Foo) Sleep(args Args, reply *int) error {
	time.Sleep(time.Second * time.Duration(args.Num1))
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addrCh chan string) {
	var foo Foo
	l, _ := net.Listen("tcp", ":0")
	server := learning_rpc.NewServer()
	_ = server.Register(&foo)
	addrCh <- l.Addr().String()
	server.Accept(l)
}

func foo(xc *xclient.XClient, ctx context.Context, typ, serviceMethod string, args *Args) {
	var reply int
	var err error
	switch typ {
	case "call":
		err = xc.Call(ctx, serviceMethod, args, &reply)
	case "broadcase":
		err = xc.Broadcast(ctx, serviceMethod, args, &reply)
	}
	if err != nil {
		log.Printf("%s %s error: %v", typ, serviceMethod, err)
	} else {
		log.Printf("%s %s success: %d + %d = %d", typ, serviceMethod, args.Num1, args.Num2, reply)
	}
}

func call(addr1, addr2 string) {
	d := xclient.NewMultiServerDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2})
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
	defer func() { _ = xc.Close() }()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			foo(xc, context.Background(), "call", "Foo.Sum", &Args{Num1: i, Num2: i * i})
		}(i)
	}
	wg.Wait()
}

func broadcast(addr1, addr2 string) {
	d := xclient.NewMultiServerDiscovery([]string{"tpc@" + addr1, "tcp@" + addr2})
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
	defer func() { _ = xc.Close() }()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			foo(xc, context.Background(), "broadcast", "Foo.Sum", &Args{Num1: i, Num2: i * i})
			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
			foo(xc, ctx, "broadcast", "Foo.Sleep", &Args{Num1: i, Num2: i * i})
		}(i)
	}
	wg.Wait()
}

func main() {
	// getMethodRecord()
	// sendRPCRequest()
	log.SetFlags(0)
	ch1 := make(chan string)
	ch2 := make(chan string)

	go startServer(ch1)
	go startServer(ch2)

	addr1 := <-ch1
	addr2 := <-ch2

	time.Sleep(time.Second)
	call(addr1, addr2)
	broadcast(addr1, addr2)
}
