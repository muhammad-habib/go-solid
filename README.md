# Go SOLID Principles

This repository demonstrates the five SOLID principles of object-oriented design using Go. Each principle is illustrated with both bad and good practice examples to help understand how to write clean, maintainable, and scalable code.

## 📚 Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [SOLID Principles](#solid-principles)
  - [1. Single Responsibility Principle (SRP)](#1-single-responsibility-principle-srp)
  - [2. Open/Closed Principle (OCP)](#2-openclosed-principle-ocp)
  - [3. Liskov Substitution Principle (LSP)](#3-liskov-substitution-principle-lsp)
  - [4. Interface Segregation Principle (ISP)](#4-interface-segregation-principle-isp)
  - [5. Dependency Inversion Principle (DIP)](#5-dependency-inversion-principle-dip)
- [Running the Examples](#running-the-examples)
- [Key Takeaways](#key-takeaways)
- [Best Practices in Go](#-best-practices-in-go)
- [Common Pitfalls to Avoid](#-common-pitfalls-to-avoid)
- [Contributing](#-contributing)

## Overview

SOLID is an acronym for five design principles intended to make software designs more understandable, flexible, and maintainable. These principles are fundamental to writing quality object-oriented code.

## Prerequisites

- Go 1.16 or higher
- Basic understanding of Go programming language
- Familiarity with structs, interfaces, and methods in Go

## Project Structure

```
go-solid/
├── 1.SRP/
│   └── main.go          # Single Responsibility Principle
├── 2.OCP/
│   └── main.go          # Open/Closed Principle
├── 3.LSP/
│   └── main.go          # Liskov Substitution Principle
├── 4.ISP/
│   └── main.go          # Interface Segregation Principle
├── 5.DIP/
│   └── main.go          # Dependency Inversion Principle
├── go.mod
├── LICENSE
└── README.md
```

Each directory contains a complete, runnable example demonstrating both bad and good practices for that specific SOLID principle.

## SOLID Principles

### 1. Single Responsibility Principle (SRP)

**"A class should have one, and only one, reason to change."**

#### Location: `1.SRP/main.go`

#### Bad Practice ❌
```go
type employee struct {
    firstName string
    lastName  string
    email     string
}

func (em *employee) saveEmployee() {
    //save to DB - This violates SRP!
}
```
**Problem**: The `employee` struct is responsible for both employee data management AND database operations.

#### Good Practice ✅
```go
type employee struct {
    firstName string
    lastName  string
    email     string
}

type empRepository struct {}

func (emr *empRepository) saveEmployee() {
    //save to DB
}
```
**Solution**: Separate concerns - `employee` handles employee data, while `empRepository` handles database operations.

---

### 2. Open/Closed Principle (OCP)

**"Software entities should be open for extension but closed for modification."**

#### Location: `2.OCP/main.go`

#### Bad Practice ❌
```go
func (em employee) getSalary() int {
    if em.role == "SWE" {
        return 3000
    } else if em.role == "SSWE" {
        return 5000
    }
    return 0
}
```
**Problem**: Adding new roles requires modifying the `getSalary()` function, violating OCP.

#### Good Practice ✅
```go
type role interface {
    getSalary() int
}

type swe struct{}
func (s swe) getSalary() int { return 3000 }

type sswe struct{}
func (s sswe) getSalary() int { return 5000 }

type employee struct {
    name string
    role role
}
```
**Solution**: Use interfaces to allow extension without modification. New roles can be added without changing existing code.

---

### 3. Liskov Substitution Principle (LSP)

**"Subtypes must be substitutable for their base types without breaking behavior."**

#### Location: `3.LSP/main.go`

#### Implementation ✅
```go
type baseEmployee interface {
    getName() string
    getSalary() int
}

type fullTimeEmployee struct {
    name   string
    salary int
}

type contractorEmployee struct {
    name        string
    hourlyRate  int
    hoursWorked int
}

func printEmployeeInfo(em baseEmployee) {
    fmt.Printf("Name: %s, Salary: %d\n", em.getName(), em.getSalary())
}
```
**Key Point**: Both `fullTimeEmployee` and `contractorEmployee` can be used interchangeably wherever `baseEmployee` is expected, without breaking the program's behavior.

---

### 4. Interface Segregation Principle (ISP)

**"Clients should not be forced to depend on interfaces they do not use."**

#### Location: `4.ISP/main.go`

#### Bad Practice ❌
```go
type Employee interface {
    GetName() string
    CalculateMonthlyPay() float64
    ApproveLeave(days int) error    // Not all employees can do this!
    AssignTask(task string) error    // Not all employees can do this!
    GenerateReport() string
}
```
**Problem**: A `Developer` or `Intern` is forced to implement methods they shouldn't have (like `ApproveLeave`).

#### Good Practice ✅
```go
type Employee interface {
    GetName() string
}

type PaidEmployee interface {
    Employee
    CalculateMonthlyPay() float64
}

type TaskAssigner interface {
    AssignTask(task string, assignee Employee) error
}

// Usage example
func main() {
    dev := Developer{Name: "Alice", Salary: 3000}
    mgr := Manager{Name: "Bob", Salary: 5000}
    intern := Intern{Name: "Charlie"}

    ProcessPayroll(dev) // ✅ Developer implements PaidEmployee
    ProcessPayroll(mgr) // ✅ Manager implements PaidEmployee
    // ProcessPayroll(intern) // ❌ Compile error - Intern doesn't implement PaidEmployee
    
    AssignWork(mgr, dev, "Implement feature") // ✅ Manager implements TaskAssigner
    // AssignWork(dev, intern, "Task") // ❌ Compile error - Developer doesn't implement TaskAssigner
}
```
**Solution**: Break down fat interfaces into smaller, more specific interfaces. Types only implement what they need. This prevents forcing implementations that don't make sense (like making an Intern handle payroll or a Developer assign tasks).

---

### 5. Dependency Inversion Principle (DIP)

**"High-level modules should not depend on low-level modules; both should depend on abstractions."**

#### Location: `5.DIP/main.go`

#### Bad Practice ❌
```go
type MySQLDatabase struct{}

func (db MySQLDatabase) SaveToMySQL(name string) {
    fmt.Println("Saving employee to MySQL:", name)
}

type EmployeeManager struct {
    database MySQLDatabase // ❌ Direct dependency on concrete type
}

func (em EmployeeManager) SaveEmployee(name string) {
    em.database.SaveToMySQL(name) // ❌ Tightly coupled to MySQL
}
```
**Problem**: The high-level `EmployeeManager` is tightly coupled to the low-level `MySQLDatabase`. If you want to switch to PostgreSQL, you need to modify `EmployeeManager`.

#### Good Practice ✅
```go
// Abstraction - both high and low level modules depend on this
type EmployeeRepository interface {
    Save(emp Employee) error
    GetByName(name string) (Employee, error)
}

// Low-level modules - implement the abstraction
type MySQLRepository struct{}
func (db MySQLRepository) Save(emp Employee) error { /* ... */ }

type PostgresRepository struct{}
func (db PostgresRepository) Save(emp Employee) error { /* ... */ }

// High-level module - depends on abstraction, not concrete types
type EmployeeManager struct {
    repository EmployeeRepository // ✅ Depends on abstraction
}

func (em EmployeeManager) AddEmployee(emp Employee) {
    em.repository.Save(emp) // ✅ Works with any repository implementation
}
```
**Solution**: Both high-level (`EmployeeManager`) and low-level modules (`MySQLRepository`, `PostgresRepository`) depend on the `EmployeeRepository` abstraction. You can easily swap database implementations without changing `EmployeeManager`.

---

## Running the Examples

Each principle has its own directory with a standalone `main.go` file. You can run any example using:

```bash
# Run SRP example
go run 1.SRP/main.go

# Run OCP example
go run 2.OCP/main.go

# Run LSP example
go run 3.LSP/main.go

# Run ISP example
go run 4.ISP/main.go

# Run DIP example
go run 5.DIP/main.go
```

## Key Takeaways

| Principle | Key Benefit | Go Implementation |
|-----------|-------------|-------------------|
| **SRP** | Easier maintenance and testing | Separate structs for different responsibilities |
| **OCP** | Extensible without modification | Use interfaces for polymorphism |
| **LSP** | Predictable substitution | Consistent interface contracts |
| **ISP** | No forced dependencies | Small, focused interfaces |
| **DIP** | Loose coupling, better testing | Depend on interfaces, not concrete types |

## 🎯 Best Practices in Go

1. **Favor composition over inheritance** - Go doesn't have classical inheritance; use struct embedding and interfaces
2. **Keep interfaces small** - Go encourages small, focused interfaces (often 1-3 methods). "The bigger the interface, the weaker the abstraction." - Rob Pike
3. **Accept interfaces, return structs** - Make your functions flexible by accepting interfaces as parameters
4. **Design from the caller's perspective** - Think about how your code will be used, not just how it's implemented
5. **Use dependency injection** - Pass dependencies (especially interfaces) as parameters rather than creating them internally
6. **Avoid interface pollution** - Don't create interfaces "just in case". Create them when you actually need abstraction

## ⚠️ Common Pitfalls to Avoid

1. **Over-engineering** - Don't apply all SOLID principles everywhere. Use them when they add value
2. **Too many small interfaces** - While ISP is important, don't create an interface for every single method
3. **Premature abstraction** - Start concrete, refactor to interfaces when you see patterns
4. **Ignoring Go idioms** - SOLID comes from OOP, but Go has its own way of doing things. Adapt accordingly
5. **Interface in the producer** - In Go, define interfaces where they're used (consumer), not where they're implemented (producer)

## 🤝 Contributing

Feel free to contribute improvements or additional examples! Each example should clearly demonstrate both the problem (bad practice) and the solution (good practice).

---

**Happy Learning! 🚀**

