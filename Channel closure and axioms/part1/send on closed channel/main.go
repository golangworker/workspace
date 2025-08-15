package main


// закрытый канал + отправка = panic
func main() {
	ch := make(chan string)
	close(ch)
	ch <- "привет"
}
