package customer

type (
	Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}
)

func GetCustomers() (customers []Customer) {
	customers = []Customer{}

	customers = append(customers,
		Customer{FirstName: "Niko", LastName: "Feng"},
		Customer{FirstName: "Dora", LastName: "Deng"},
		Customer{FirstName: "Michelle", LastName: "Feng"})

	return customers
}
