package employeemodule

import (
	"fmt"
)

type Employee struct {
	name         string
	employeeCode int
}

// contsructor of Employee struct
func NewEmployee(name string, employeecode int) *Employee {
	return &Employee{
		name:         name,
		employeeCode: employeecode,
	}
}

// we can validate , without user implimemting this in his code
func (e *Employee) validate() error {
	if e.name == "" {

		fmt.Println("printing directly-->>> name is empty")

		return fmt.Errorf("from error return-->>> Employee details cant be empty")
	}
	return nil
}

func (e *Employee) AddEmployee() error {
	// add employee details to database opertaion
	if err := e.validate(); err != nil {
		return err
	}
	// now we do some opertaion herie , adding data to database
	//sql qurries
	e.createDBOjects()

	fmt.Println("sucessfull")
	return nil
}

func (e *Employee) createDBOjects() {
	// db opertation
}
