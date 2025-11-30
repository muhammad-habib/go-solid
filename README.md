# Go SOLID Principles

This repository demonstrates the five SOLID principles of object-oriented design using Go. Each principle is illustrated with both bad and good practice examples to help understand how to write clean, maintainable, and scalable code.

## 📚 Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [SOLID Principles](#solid-principles)
  - [1. Single Responsibility Principle (SRP)](#1-single-responsibility-principle-srp)
  - [2. Open/Closed Principle (OCP)](#2-openclosed-principle-ocp)
  - [3. Liskov Substitution Principle (LSP)](#3-liskov-substitution-principle-lsp)
  - [4. Interface Segregation Principle (ISP)](#4-interface-segregation-principle-isp)
  - [5. Dependency Inversion Principle (DIP)](#5-dependency-inversion-principle-dip)
- [Running the Examples](#running-the-examples)
- [Key Takeaways](#key-takeaways)

## Overview

SOLID is an acronym for five design principles intended to make software designs more understandable, flexible, and maintainable. These principles are fundamental to writing quality object-oriented code.

## Prerequisites

- Go 1.25 or higher
- Basic understanding of Go programming language
- Familiarity with structs, interfaces, and methods in Go

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
```
**Solution**: Break down fat interfaces into smaller, more specific interfaces. Types only implement what they need.

---

### 5. Dependency Inversion Principle (DIP)

**"High-level modules should not depend on low-level modules; both should depend on abstractions."**

#### Location: `5.DIP/main.go`

#### Implementation ✅
```go
type baseEmployee interface {
    getName() string
    getSalary() int
}

// High-level function depends on abstraction, not concrete types
func printEmployeeInfo(em baseEmployee) {
    fmt.Printf("Name: %s, Salary: %d\n", em.getName(), em.getSalary())
}
```
**Key Point**: The `printEmployeeInfo()` function depends on the `baseEmployee` interface (abstraction), not on concrete implementations like `fullTimeEmployee` or `contractorEmployee`. This makes the code more flexible and testable.

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

## 🎯 Best Practices

1. **Favor composition over inheritance** - Go doesn't have classical inheritance; use struct embedding and interfaces
2. **Keep interfaces small** - Go encourages small, focused interfaces (often 1-3 methods)
3. **Accept interfaces, return structs** - Make your functions flexible by accepting interfaces
4. **Design from the caller's perspective** - Think about how your code will be used

## 📖 Further Reading

- [Effective Go](https://golang.org/doc/effective_go)
- [Go Proverbs](https://go-proverbs.github.io/)
- [SOLID Principles in Object-Oriented Design](https://en.wikipedia.org/wiki/SOLID)

## 🤝 Contributing

Feel free to contribute improvements or additional examples! Each example should clearly demonstrate both the problem (bad practice) and the solution (good practice).

---

**Happy Learning! 🚀**

