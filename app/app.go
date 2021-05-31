package main

import (
	"fmt"
	// "sync"
	customer "github.com/nikosheng/golearning/customer"
)

func main() {
	// customers := GetCustomers()

	// for index := 0; index < len(customers); index++ {
	// 	fmt.Println(customers[index])
	// }

	// for _, customer := range customers {
	// 	fmt.Println(customer)
	// }

	// intSet := map[int]bool{}
	// vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	// for _, v := range vals {
	// 	intSet[v] = true
	// }
	// fmt.Println(len(vals), len(intSet))
	// fmt.Println(intSet[5])
	// fmt.Println(intSet[500])
	// if intSet[10] {
	// 	fmt.Println("10 is in the set")
	// }

	// if n := rand.Intn(10); n == 0 {
	// 	fmt.Println("That's too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big:", n)
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }

	// words := []string{
	// 	"a",
	// 	"cow",
	// 	"is",
	// 	"walking",
	// }

	// for _, word := range words {
	// 	switch size := len(word); size {
	// 	case 1, 2:
	// 		fmt.Println(word, "is a short word")
	// 	default:
	// 		fmt.Println(word, "is a long word")
	// 	}
	// }

	// customers := make(map[int]string, 10)
	// customers[1] = "Niko"
	// customers[2] = "Dora"

	// for index, _ := range customers {
	// 	fmt.Println(customers[index])
	// }

	// fmt.Println(addTo(3))
	// fmt.Println(addTo(3, 2))
	// fmt.Println(addTo(3, 2, 4, 6, 8))

	// twoBase := makeMult(2)
	// threeBase := makeMult(3)

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(twoBase(i), threeBase(i))
	// }

	// var scene sync.Map

	// // 将键值对保存到sync.Map
	// scene.Store("greece", 97)
	// scene.Store("london", 100)
	// scene.Store("egypt", 200)

	// // 从sync.Map中根据键取值
	// fmt.Println(scene.Load("london"))

	// // 根据键删除对应的键值对
	// scene.Delete("london")

	// // 遍历所有sync.Map中的键值对
	// scene.Range(func(k, v interface{}) bool {

	// 	fmt.Println("iterate:", k, v)
	// 	return true
	// })

	// b := 255
	// a := &b
	// fmt.Println("address of b is", a)
	// fmt.Println("value of b is", *a)

	customers := customer.GetCustomers()

	for _, cus := range customers {
		fmt.Printf("Customer: %s, %s\n", cus.FirstName, cus.LastName)
	}

}

func modMap(m map[int]string) {
	m[2] = "yes"
	m[3] = "no"
	delete(m, 1)
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// func getCustomers() (customers []string) {
// 	customers = []string{}

// 	customers = append(customers, "Niko Feng")
// 	customers = append(customers, "Dora Deng")
// 	return customers
// }
