package main
import ("fmt")

type Person struct {
name string
age int
job string
salary int
}


func printPerson(p Person){
	// Access and print Pers1 info
fmt.Println("Name: ", p.name)
fmt.Println("Age: ", p.age)
fmt.Println("Job: ", p.job)
fmt.Println("Salary: ", p.salary)
}

func main() {
var pers1 Person
// Pers1 specification
pers1.name = "Hege"
pers1.age = 45
pers1.job = "Teacher"
pers1.salary = 6000

printPerson(pers1)


}