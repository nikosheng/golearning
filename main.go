package main

import (
	"fmt"

	customer "github.com/nikosheng/golearning/customer"
)

func main() {
	customers := customer.GetCustomers()

	for _, cus := range customers {
		fmt.Printf("Customer: %s, %s\n", cus.FirstName, cus.LastName)
	}

}
