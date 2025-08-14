package bank

import (
	"errors"
	"fmt"
	"math/rand"
)

type bank struct {
	money float64
}

var errNotEnough error = errors.New("недостаточно средств")

// Функция NewUser для создания нового пользователя
// Принимает: money - количество денег на счету
// Выводит: &bank - указатель на созданную структуру

func NewUser(money float64) *bank {
	return &bank{money: money}
}

// Метод ShowBalance - показывает текущий счёт
// Выводит: error - ошибку операции
func (b bank) ShowBalance() error {
	err := randomBreakdown()
	if err != nil {
		return err
	}
	fmt.Printf("Твой баланс равен: %.2fUSD\n", b.money)
	return nil
}

// Метод CashWithdrawal - вывод средств из банкомата
// Принимает: charges - сумма, которую сняли
// Выводит: error - ошибку операции
func (b *bank) CashWithdrawal(charges float64) error {
	err := randomBreakdown()
	if err != nil {
		return err
	}
	if b.money - charges < 0 {
		return errNotEnough
	}
	b.money -= charges
	fmt.Println("Операция прошла успешно!")
	fmt.Printf("Вы сняли %.2fUSD\n", charges)
	return nil
}

// Метод Payment - проводит оплату
// Принимает: cost - стоимость услуги
// Выводит: error - ошибку операции
func (b *bank) Payment(cost float64) error {
	err := randomBreakdown()
	if err != nil {
		return err
	}
	if b.money - cost < 0 {
		return errNotEnough
	}
	b.money -= cost
	fmt.Println("Операция прошла успешно!")
	fmt.Printf("Вы оплатили услугу, стоимостью %.2fUSD\n", cost)
	return nil
}

// Функция randomBreakdown - вероятность ошибки проведении операции
// Выводит: error - ошибку операции
func randomBreakdown() error {
	n := rand.Intn(10)
	if n < 4 {
		return errors.New("ошибка операции")
	}
	return nil
}
