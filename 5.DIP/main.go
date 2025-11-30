package main

import "fmt"

//////////--------------------Bad Practice--------------------/////////////////////////

// ❌ High-level module (EmployeeManager) depends directly on low-level modules (MySQLDatabase, PostgresDatabase)
// This violates DIP because the high-level logic is tightly coupled to concrete database implementations

//type MySQLDatabase struct{}
//
//func (db MySQLDatabase) SaveToMySQL(name string) {
//	fmt.Println("Saving employee to MySQL:", name)
//}
//
//type EmployeeManager struct {
//	database MySQLDatabase // ❌ Direct dependency on concrete type
//}
//
//func (em EmployeeManager) SaveEmployee(name string) {
//	em.database.SaveToMySQL(name) // ❌ Tightly coupled to MySQL
//}
//
//func main() {
//	db := MySQLDatabase{}
//	manager := EmployeeManager{database: db}
//	manager.SaveEmployee("Mohamed")
//	// If we want to switch to Postgres, we need to modify EmployeeManager
//}

//////////////-----------------------------Good Practice-------------------/////////////////////////////////////////////////////////

type Employee struct {
	Name   string
	Salary int
}

// EmployeeRepository Abstraction (interface) - both high and low level modules depend on this
type EmployeeRepository interface {
	Save(emp Employee) error
	GetByName(name string) (Employee, error)
}

// MySQLRepository Low-level module - implements the abstraction
type MySQLRepository struct{}

func (db MySQLRepository) Save(emp Employee) error {
	fmt.Printf("💾 Saving employee '%s' to MySQL database\n", emp.Name)
	return nil
}

func (db MySQLRepository) GetByName(name string) (Employee, error) {
	fmt.Printf("🔍 Fetching employee '%s' from MySQL database\n", name)
	return Employee{Name: name, Salary: 5000}, nil
}

// PostgresRepository Low-level module - implements the abstraction
type PostgresRepository struct{}

func (db PostgresRepository) Save(emp Employee) error {
	fmt.Printf("💾 Saving employee '%s' to PostgreSQL database\n", emp.Name)
	return nil
}

func (db PostgresRepository) GetByName(name string) (Employee, error) {
	fmt.Printf("🔍 Fetching employee '%s' from PostgreSQL database\n", name)
	return Employee{Name: name, Salary: 5000}, nil
}

// MongoRepository Low-level module - implements the abstraction
type MongoRepository struct{}

func (db MongoRepository) Save(emp Employee) error {
	fmt.Printf("💾 Saving employee '%s' to MongoDB database\n", emp.Name)
	return nil
}

func (db MongoRepository) GetByName(name string) (Employee, error) {
	fmt.Printf("🔍 Fetching employee '%s' from MongoDB database\n", name)
	return Employee{Name: name, Salary: 5000}, nil
}

// EmployeeManager High-level module - depends on abstraction (EmployeeRepository), not concrete types
type EmployeeManager struct {
	repository EmployeeRepository // ✅ Depends on abstraction, not concrete implementation
}

func (em EmployeeManager) AddEmployee(emp Employee) {
	err := em.repository.Save(emp)
	if err != nil {
		fmt.Println("Error saving employee:", err)
	}
}

func (em EmployeeManager) FindEmployee(name string) {
	emp, err := em.repository.GetByName(name)
	if err != nil {
		fmt.Println("Error fetching employee:", err)
		return
	}
	fmt.Printf("✅ Found employee: %s, Salary: %d\n", emp.Name, emp.Salary)
}

func main() {
	// ✅ High-level module (EmployeeManager) doesn't know about concrete database implementations
	// ✅ Both high-level and low-level modules depend on the EmployeeRepository abstraction
	// ✅ We can easily swap database implementations without changing EmployeeManager

	mohamed := Employee{Name: "Mohamed", Salary: 5000}
	ahmed := Employee{Name: "Ahmed", Salary: 6000}
	ali := Employee{Name: "Ali", Salary: 4500}

	// Using MySQL
	mysqlRepo := MySQLRepository{}
	manager1 := EmployeeManager{repository: mysqlRepo}
	manager1.AddEmployee(mohamed)
	manager1.FindEmployee("Mohamed")

	fmt.Println()

	// Using PostgreSQL
	postgresRepo := PostgresRepository{}
	manager2 := EmployeeManager{repository: postgresRepo}
	manager2.AddEmployee(ahmed)
	manager2.FindEmployee("Ahmed")

	fmt.Println()

	// Using MongoDB
	mongoRepo := MongoRepository{}
	manager3 := EmployeeManager{repository: mongoRepo}
	manager3.AddEmployee(ali)
	manager3.FindEmployee("Ali")

	// High-level modules (EmployeeManager) should not depend on low-level modules (MySQLRepository, PostgresRepository)
	// Both should depend on abstractions (EmployeeRepository interface)
}
