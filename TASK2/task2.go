package main
import ("fmt")

type employee struct {
name string
position string
salary int
}

type company struct {
	compName string
	employees []employee
}

func printEmployee(e employee){
		
	fmt.Println("Name: ", e.name)
	fmt.Println("Position: ", e.position)
	fmt.Println("Salary: ", e.salary)
}
func printCompany(c company){
		
	fmt.Println("Company Name: ", c.compName)

	for i :=0; i < len(c.employees); i++ {
		printEmployee(c.employees[i])
	}
}
func main() {

	var e1,e2,e3 employee
	
	e1.name = "A"
	e1.position = "Teacher"
	e1.salary = 6000

	e2.name = "B"
	e2.position = "Coordinator"
	e2.salary = 10000

	e3.name = "C"
	e3.position = "Manager"
	e3.salary = 30000

	var empArray = []employee{
		e1,e2,e3,
	}
	var c1 company
	c1.compName = "FAST"
	c1.employees = empArray


	printCompany(c1)


}