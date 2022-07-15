package main

import (
	"fmt"

	"github.com/go-interfaces/users"
)

func main() {
	var frank users.Cashier = users.NewEmployee("Daniel")
	var robert users.Admin = users.NewEmployee("Robert")

	total := frank.CalcTotal(80, 60, 65, 93, 93.6)
	fmt.Printf("El total es: %f\n", total)

	robert.DeactivateEmployee(frank)

	fmt.Println(frank.CalcTotal(80, 60, 65, 93, 93.6))

}
