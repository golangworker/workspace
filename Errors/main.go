package main

import (
	"fmt"
	"math/rand"

	"app/bank"
)

func main() {
	user := bank.NewUser(1000)
	for {
		err := user.ShowBalance()
		if err != nil {
			fmt.Println(err.Error())
		}
		err = user.CashWithdrawal(float64(rand.Intn(50)))
		if err != nil {
			fmt.Println(err.Error())
		}
		err = user.Payment(float64(rand.Intn(75)))
		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "недостаточно средств" {
				break
			}
		}
	}
}
