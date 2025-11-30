package main

import "fmt"

//---------------------------------------------//Bad Practice//--------------------------------------------------------///

//// Employee Fat interface – mixes many responsibilities
//type Employee interface {
//	GetName() string
//
//	// CalculateMonthlyPay payment related
//	CalculateMonthlyPay() float64
//
//	// ApproveLeave management related
//	ApproveLeave(days int) error
//	AssignTask(task string) error
//
//	// GenerateReport reporting
//	GenerateReport() string
//}
//
//type Developer struct {
//	Name   string
//	Salary float64
//}
//
//func (d Developer) GetName() string { return d.Name }
//func (d Developer) CalculateMonthlyPay() float64 {
//	return d.Salary
//}
//
//// ApproveLeave ❌ Developer is forced to implement things it shouldn’t care about
//func (d Developer) ApproveLeave(days int) error {
//	return fmt.Errorf("developer cannot approve leave")
//}
//
//func (d Developer) AssignTask(task string) error {
//	return fmt.Errorf("developer cannot assign tasks")
//}
//
//func (d Developer) GenerateReport() string {
//	return "dev report"
//}
//
//type Intern struct {
//	Name string
//}
//
//func (i Intern) GetName() string { return i.Name }
//func (i Intern) CalculateMonthlyPay() float64 {
//	return 0 // maybe unpaid
//}
//
//// ApproveLeave ❌ Even worse here
//func (i Intern) ApproveLeave(days int) error {
//	return fmt.Errorf("intern cannot approve leave")
//}
//
//func (i Intern) AssignTask(task string) error {
//	return fmt.Errorf("intern cannot assign tasks")
//}
//
//func (i Intern) GenerateReport() string {
//	return ""
//}

//---------------------------------------------//Good Practice//--------------------------------------------------------///

// PersonInCompany Minimal base identity
type Employee interface {
	GetName() string
}

// PaidEmployee Only for people who get paid
type PaidEmployee interface {
	Employee
	CalculateMonthlyPay() float64
}

// TaskAssigner Only for people who can assign work
type TaskAssigner interface {
	AssignTask(task string, assignee Employee) error
}

type Developer struct {
	Name   string
	Salary float64
}

func (d Developer) GetName() string { return d.Name }

func (d Developer) CalculateMonthlyPay() float64 {
	return d.Salary
}

// ✅ Developer is *not* forced to approve leave or assign tasks

type Manager struct {
	Name   string
	Salary float64
}

func (m Manager) GetName() string { return m.Name }

func (m Manager) CalculateMonthlyPay() float64 {
	return m.Salary
}

// AssignTask Manager has more responsibilities
func (m Manager) AssignTask(task string, assignee Employee) error {
	fmt.Printf("Manager %s assigned '%s' to %s\n", m.Name, task, assignee.GetName())
	return nil
}

type Intern struct {
	Name string
}

func (i Intern) GetName() string { return i.Name }

// Maybe unpaid, maybe small stipend – but does *not* implement PaidEmployee,
// if we decide they’re out of payroll flow.

// ProcessPayroll Payroll only cares about PaidEmployee
func ProcessPayroll(e PaidEmployee) {
	fmt.Printf("Paying %s: %.2f EUR\n", e.GetName(), e.CalculateMonthlyPay())
}

// AssignWork Task assignment only needs TaskAssigner
func AssignWork(assigner TaskAssigner, dev Employee, task string) {
	_ = assigner.AssignTask(task, dev)
}

func main() {
	dev := Developer{Name: "Alice", Salary: 3000}
	mgr := Manager{Name: "Bob", Salary: 5000}
	intern := Intern{Name: "Charlie"}

	ProcessPayroll(dev) // ok: Developer is PaidEmployee
	ProcessPayroll(mgr) // ok: Manager is PaidEmployee
	//ProcessPayroll(intern) // ❌ compile error – Intern is not PaidEmployee

	AssignWork(mgr, dev, "Implement new feature") // ok
	AssignWork(dev, intern, "Review code")        // ❌ compile error – Developer is not TaskAssigner
}
