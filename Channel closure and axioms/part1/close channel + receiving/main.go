package main

import "fmt"

// закрытый канал + получение = zero value мгновенно
func main() {
	ch := make(chan int)
	close(ch)
	v, ok := <- ch
	fmt.Println("Is there a value in the channel?", ok)
	fmt.Println("Default value:", v)
}
