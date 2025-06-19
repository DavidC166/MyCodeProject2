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

	pers1.name = "Yasin"
	pers1.age = 20
	pers1.job = "Pastry chef"
	pers1.salary = 2000

	pers2.name = "Sat"
	pers2.age = 32
	pers2.job = "Sous chef"
	pers2.salary = 3500

	printPerson(pers1)
	
	printPerson(pers2)
}

func printPerson(pers Employee) {

	fmt.Println("Name:", pers.name)
	fmt.Println("Age:", pers.age)
	fmt.Println("Job:", pers.job)
	fmt.Println("Salary:", pers.salary)
  fmt.Println("-- -- -- -- --")
}