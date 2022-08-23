package main

import (
	"context"
	"fmt"
	"time"
)

var stop chan bool

func reqTask(ctx context.Context, name string) {
	for {
		select {
		//case <-stop:
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	//stop = make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask(ctx, "worker 1")
	go reqTask(ctx, "worker 2")
	time.Sleep(3 * time.Second)
	//stop <- true
	cancel()
	time.Sleep(3 * time.Second)
}
