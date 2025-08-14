package module

import "fmt"

type Messenger interface {
	Send() string
}

type Postman struct {}

func (p Postman) SelectiveMailing(messages []Messenger) {
	for _, v := range messages {
		fmt.Println(v.Send())
	}
}
