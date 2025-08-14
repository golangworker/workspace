package messengers

import "math/rand"

type sms struct {
	recipient string
	message string
	id int
}

func NewSms(recipient string) *sms {
	return &sms{
		recipient: recipient,
		message: "Новое sms сообщение",
		id: rand.Int(),
	}
}

func (s sms) Send() string {
	return s.message
}
