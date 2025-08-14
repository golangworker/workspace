package messengers

import (
	"math/rand"
)

type email struct {
	recipient string
	message string
	id int
}

func NewEmail(recipient string) *email {
	return &email{
		recipient: recipient,
		message: "Сообщение в email",
		id: rand.Int(),
	}
}

func (e email) Send() string {
	return e.message
}
