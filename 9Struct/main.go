package main

import "fmt"

type Student struct {
	Name string
	School string
	Year int
}

func main() {
	student1 := Student{
		Name: "Alex",
		School: "Founders High School",
		Year: 2024,
	}

		student2 := Student{
			Name: "David",
			School: "City University",
			Year: 2006,
		}

	fmt.Println("Student Information")
	fmt.Println("Name:", student1.Name)
	fmt.Println("School:", student1.School)
	fmt.Println("Year:", student1.Year)

	fmt.Println("Student Information")
	fmt.Println("Name:", student2.Name)
	fmt.Println("School:", student2.School)
	fmt.Println("Year:", student2.Year)
}