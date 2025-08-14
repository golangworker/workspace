package messengers

import "math/rand"

type push struct {
	recipient string
	message string
	id int
}

func NewPush(recipient string) *push {
	return &push{
		recipient: recipient,
		message: "Новое Push-уведомление",
		id: rand.Int(),
	}
}

func (p push) Send() string {
	return p.message
}
