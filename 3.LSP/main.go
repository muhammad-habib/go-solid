package main

import "fmt"

type baseEmployee interface {
	getName() string
	getSalary() int
}
type fullTimeEmployee struct {
	name   string
	salary int
}

func (em fullTimeEmployee) getName() string { return em.name }

func (em fullTimeEmployee) getSalary() int { return em.salary }

type contractorEmployee struct {
	name        string
	hourlyRate  int
	hoursWorked int
}

func (cem contractorEmployee) getName() string { return cem.name }

func (cem contractorEmployee) getSalary() int { return cem.hourlyRate * cem.hoursWorked }

func printEmployeeInfo(em baseEmployee) {
	fmt.Printf("Name: %s, Salary: %d\n", em.getName(), em.getSalary())
}

func main() {
	em1 := fullTimeEmployee{
		name:   "Mohamed",
		salary: 5000,
	}

	em2 := contractorEmployee{
		name:        "Ahmed",
		hourlyRate:  120,
		hoursWorked: 10,
	}

	//Subtypes (concrete types) must be substitutable for their base type (interface or parent type) without breaking behavior.
	printEmployeeInfo(em1)
	printEmployeeInfo(em2)
}
