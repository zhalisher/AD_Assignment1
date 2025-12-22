package company

import (
	"fmt"
)

type Employee interface {
	GetDetails() string
}
type FullTimeEmployee struct {
	ID   uint64
	Name string
	Wage float64
}
type PartTimeEmployee struct {
	ID   uint64
	Name string
	Wage float64
}
type Company struct {
	Employees map[uint64]Employee
}

// func for full time
func (fullTime FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full time. ID: %v. Name: %v. Wage: %v\n",
		fullTime.ID, fullTime.Name, fullTime.Wage)
}

// func fot part time
func (partTime PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part time. ID: %v. Name: %v. Wage: %v\n",
		partTime.ID, partTime.Name, partTime.Wage)
}

// add
func (company *Company) AddEmployee(employee Employee, id uint64) {
	company.Employees[id] = employee
	fmt.Println("Employee added : ", id)
}
func (company *Company) ListtAllEmployee() {
	fmt.Println("Employees list:")
	for _, e := range company.Employees {
		fmt.Println(e.GetDetails())
	}
}
func CompanyCLI() {
	company := Company{make(map[uint64]Employee)}
	id := 0

	fmt.Println("Welcome to Employee Management System application!")

	for loop := true; loop; {

		var choice uint8

		fmt.Println("==============================")
		fmt.Println("Add Employee Full: 1")
		fmt.Println("Add Employee Part: 2")
		fmt.Println("List Employee: 3")
		fmt.Println("Exit: 4")
		fmt.Println("==============================")

		fmt.Println("What to do now?")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var wage float64

			fmt.Println("Enter Name of employee: ")
			fmt.Scan(&name)

			fmt.Println("Enter Wage of employee: ")
			fmt.Scan(&wage)

			id++
			employee := FullTimeEmployee{
				ID:   uint64(id),
				Name: name,
				Wage: float64(wage),
			}
			company.AddEmployee(employee, employee.ID)
		case 2:
			var name string
			var wage float64

			fmt.Println("Enter Name of employee: ")
			fmt.Scan(&name)

			fmt.Println("Enter Wage of employee: ")
			fmt.Scan(&wage)

			id++
			employee := PartTimeEmployee{
				ID:   uint64(id),
				Name: name,
				Wage: float64(wage),
			}
			company.AddEmployee(employee, employee.ID)
		case 3:
			company.ListtAllEmployee()
		case 4:
			fmt.Println("Exiting from application")
			loop = false
		}

	}
}
