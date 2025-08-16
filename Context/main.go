package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx, cancelParentCtx := context.WithCancel(context.Background())
	middleCtx, cancelMiddleCtx := context.WithCancel(parentCtx)
	childCtx, cancelChildCtx := context.WithCancel(middleCtx)
	go Foo(1, parentCtx)
	go Boo(2, middleCtx)
	go Soo(3, childCtx)

	time.Sleep(time.Second)
	cancelChildCtx()
	time.Sleep(time.Second)
	cancelMiddleCtx()
	time.Sleep(time.Second)
	cancelParentCtx()
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Все горутины завершились")
}

func Foo(n int, ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("Горутина Foo №%d завершилась\n", n)
			return
		default:
			fmt.Println("Foo действует")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func Boo(n int, ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("Горутина Boo №%d завершилась\n", n)
			return
		default:
			fmt.Println("Boo действует")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func Soo(n int, ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("Горутина Soo №%d завершилась\n", n)
			return
		default:
			fmt.Println("Soo действует")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
