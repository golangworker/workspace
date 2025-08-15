package main

// закрытый канал + закрытие = panic
func main() {
	ch := make(chan int)
	close(ch)
	close(ch)
}
