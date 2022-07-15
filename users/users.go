package users

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}

type Employee struct {
	User
	Active bool
	status string
}

type Cashier interface {
	CalcTotal(item ...float32) float32
	deactivate()
	reactive()
}

type Admin interface {
	DeactivateEmployee(e Cashier)
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
		status: "receiver",
	}
}

func (e *Employee) CalcTotal(item ...float32) float32 {

	if e.status == "receiver" {
		fmt.Println("Diferente receiver")
	}
	if !e.Active {
		fmt.Println("Cashier deactivated")
		return 0
	}

	var sum float32

	for _, i := range item {
		sum += i
	}

	return sum * 1.15
}

func (e *Employee) deactivate() {
	e.Active = false
}

func (e *Employee) reactive() {
	e.Active = true
}

func (e *Employee) DeactivateEmployee(c Cashier) {
	c.deactivate()
}
