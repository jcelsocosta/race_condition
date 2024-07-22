package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type itemType struct {
	itemId int
	name   string
	price  int
	unit   int
}

func (i *itemType) takeItem(customer *customerType, unit int) {
	// Simula um atraso para aumentar a chance de condição de corrida
	time.Sleep(50 * time.Millisecond)

	if i.unit > 0 && i.unit >= unit {
		i.unit = i.unit - unit
		fmt.Println("O Cliente " + customer.name + " comprou " + strconv.Itoa(unit) + " unidade(s) de " + i.name + " e a quantidade restante no estoque é " + strconv.Itoa(i.unit))
	} else {
		fmt.Println("O Cliente " + customer.name + " não conseguiu comprar " + strconv.Itoa(unit) + " de " + i.name)
	}
}

type customerType struct {
	customerId int
	name       string
}

var wg sync.WaitGroup

func main() {
	shirt := itemType{441, "camisa", 1000, 10}
	customerA := customerType{1, "A"}
	customerB := customerType{2, "B"}
	customerC := customerType{3, "C"}

	wg.Add(3)

	go routineCustomer(&shirt, &customerA, 4)

	go routineCustomer(&shirt, &customerB, 5)

	go routineCustomer(&shirt, &customerC, 3)

	wg.Wait()
}

func routineCustomer(shirt *itemType, customer *customerType, unit int) {
	shirt.takeItem(customer, unit)

	wg.Done()
}
