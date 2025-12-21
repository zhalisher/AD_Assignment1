package Company

type Employee interface {
	GetDetails() string
}
type FullTimeEmployee struct {
}
type PartTimeEmployee struct {
}
type Company struct {
	Employee map[string]Employee
}

func CompanyCLI() {

}
