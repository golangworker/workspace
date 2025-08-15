package main

// nil канал + отправка = deadlock
func main() {
	var nilChannel chan string
	nilChannel <- "привет"
}
