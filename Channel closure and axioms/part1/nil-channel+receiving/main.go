package main

import "fmt"

// nil канал + получение = deadlock
func main() {
	var nilChannel chan string
	fmt.Println(<-nilChannel)
}
