package main

import (
	"database/sql"
	"fmt"

	"github.com/rhadamez/go1/internal/order/entity"
	"github.com/rhadamez/go1/internal/order/infra/database"
)

func main() {
	order, err := entity.NewOrder("123", 10, 2)
	if err != nil {
		panic(err)
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The final price is: %f", order.FinalPrice)

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}

	repository := database.NewOrderRepository(db)
	err = repository.Save(order)
	if err != nil {
		panic(err)
	}
}
