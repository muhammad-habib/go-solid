package main

import "fmt"

//////////--------------------Bad Practice--------------------/////////////////////////

//type employee struct {
//	name string
//	role string
//}
//
//func (em employee) getSalary() int {
//	if em.role == "SWE" {
//		return 3000
//	} else if em.role == "SSWE" {
//		return 5000
//	}
//	return 0
//}
//
//func main() {
//	em1 := employee{
//		name: "Mohamed",
//		role: "SWE",
//	}
//
//	em2 := employee{
//		name: "Ahmed",
//		role: "SSWE",
//	}
//
//	fmt.Println("Salary", em1.getSalary())
//	fmt.Println("Salary", em2.getSalary())
//}

//////////////-----------------------------Good Practice-------------------/////////////////////////////////////////////////////////

type role interface {
	getSalary() int
}
type employee struct {
	name string
	role role
}
type swe struct{}

func (s swe) getSalary() int { return 3000 }

type sswe struct{}

func (s sswe) getSalary() int { return 5000 }

func (em employee) getSalary() int {
	return em.role.getSalary()
}

func main() {
	em1 := employee{
		name: "Mohamed",
		role: swe{},
	}

	em2 := employee{
		name: "Ahmed",
		role: sswe{},
	}

	// using interface for the role giving the flexibility to extend
	//the code by adding more roles without modifying existing getSalary func
	fmt.Println("Salary", em1.getSalary())
	fmt.Println("Salary", em2.getSalary())
}
