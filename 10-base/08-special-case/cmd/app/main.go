package main

import (
	"fmt"

	"github.com/jibaru/special-case/internal/domain/specialcase"
	"github.com/jibaru/special-case/internal/repository"
)

func main() {
	repo := repository.NewOrderRepository()

	for _, id := range []int{1, 0} {
		order := repo.FindOrderById(id)

		// Check if the order is a special case (NoOrder)
		if _, isNoOrder := order.(specialcase.NoOrder); isNoOrder {
			fmt.Println("No order found, handling special case.")
		} else {
			fmt.Printf("Order found: %+v\n", order)
		}
	}
}
