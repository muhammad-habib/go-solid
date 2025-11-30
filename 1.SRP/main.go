package main

import (
	"fmt"
)

//////////--------------------Bad Practice--------------------/////////////////////////

//type employee struct {
//	firstName string
//	lastName  string
//	email     string
//}
//
//type empRepository struct {
//}
//
//func (em *employee) getFullName() string {
//	return fmt.Sprintf("%s %s", em.firstName, em.lastName)
//}
//
//func (em *employee) getEmail() string {
//	return em.email
//}
//
//func (em *employee) saveEmployee() {
//	//save to DB
//}
//
//func main() {
//	em := employee{
//		firstName: "Mohamed",
//		lastName:  "Habib",
//		email:     "mohamed@gmail.com",
//	}
//	em.getFullName()
//	em.getEmail()
//	em.saveEmployee()
//}

type employee struct {
	firstName string
	lastName  string
	email     string
}

type empRepository struct {
}

func (em *employee) getFullName() string {
	return fmt.Sprintf("%s %s", em.firstName, em.lastName)
}

func (em *employee) getEmail() string {
	return em.email
}

func (emr *empRepository) saveEmployee() {
	//save to DB
}

func main() {
	user := employee{
		firstName: "Mohamed",
		lastName:  "Habib",
		email:     "mohamed@gmail.com",
	}
	user.getFullName()
	user.getEmail()
	//saving to DB is different responsibility that why we gave userRepository struct to handle database operations.
	emr := empRepository{}
	emr.saveEmployee()
}
