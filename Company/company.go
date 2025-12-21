package Company

import (
	"fmt"
)

type Employee interface {
	GetDetails() string
}
type FullTimeEmployee struct {
}
type PartTimeEmployee struct {
}
type Company struct {
	Employee map[uint64]Employee
}

// func for full time
func (fullTime FullTimeEmployee) GetDetails() string {
	return "Full Time Employee"
}

// func fot part time
func (partTime PartTimeEmployee) GetDetails() string {
	return "Part Time Employee"
}

func CompanyCLI() {
	fmt.Println("Welcome to Employee Management System application!")

	for loop := true; loop; {

		var choice uint8

		fmt.Println("==============================")
		fmt.Println("Add Employee: 1")
		fmt.Println("List Employee: 2")
		fmt.Println("Exit: 3")
		fmt.Println("==============================")

		fmt.Println("What to do now?")
		fmt.Scan(&choice)

		switch choice {
		case 1:
		case 2:
		case 3:
			fmt.Println("Exiting from application")
			loop = false
		}

	}
}
