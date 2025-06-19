package main

import "fmt"

type Employee struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var pers1 Employee
	var pers2 Employee
	var pers5 Employee

	pers1.name = "Yasin"
	pers1.age = 20
	pers1.job = "Pastry chef"
	pers1.salary = 2000

	pers2.name = "Sat"
	pers2.age = 32
	pers2.job = "Sous chef"
	pers2.salary = 3500

	pers5.name = "Mike"
	pers5.age = 21
	pers5.job = "HRH"
	pers5.salary = 88888


	var name string
	var age int
	var job string
	var salary int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Name:", name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("Age:", age)
	fmt.Print("Enter your job: ")
	fmt.Scanln(&job)
	fmt.Print("Enter your salary: ")
	fmt.Scanln(&salary)
	fmt.Println("Salary:", salary)

	printPerson(pers1)
	
	printPerson(pers2)

	printPerson(pers5)

	var name3 string
	var age3 int
	var job3 string
	var salary3 int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name3)
	fmt.Println("Name:", name3)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age3)
	fmt.Println("Age:", age3)
	fmt.Print("Enter your job: ")
	fmt.Scanln(&job3)
	fmt.Print("Enter your salary: ")
	fmt.Scanln(&salary3)
	fmt.Println("Salary:", salary3)
}

func printPerson(pers Employee) {

	fmt.Println("Name:", pers.name)
	fmt.Println("Age:", pers.age)
	fmt.Println("Job:", pers.job)
	fmt.Println("Salary:", pers.salary)

}