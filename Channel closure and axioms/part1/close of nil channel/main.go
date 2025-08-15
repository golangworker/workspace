package main

// nil канал + закрытие = panic
func main() {
	var nilChannel chan string
	close(nilChannel)
}
