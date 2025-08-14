package main

import (
	"app/module"
	"app/module/messengers"
)

func main() {
	m1 := messengers.NewEmail("gondon")
	m2 := messengers.NewPush("siska")
	m3 := messengers.NewSms("popka")
	allM := []module.Messenger{m1, m2, m3}
	p := module.Postman{}
	p.SelectiveMailing(allM)
}
