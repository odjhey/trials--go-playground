// package to manage worker expenses in a company
package worker

// a namespace that allows creation of different type of employees
type WorkerFactory struct{}

// create an employee, record perdiem
func (WorkerFactory) NewHire(name string, totalPerDiem int) (*perm, error) {
	return &perm{totalPerDiem: totalPerDiem, name: name}, nil
}

// create a contractor, record cashadvances
func (WorkerFactory) NewContract(name string, totalCashAdvance int) (*contract, error) {
	return &contract{totalCashAdvance: totalCashAdvance, name: name}, nil
}

// create an intern, record allowances
func (WorkerFactory) NewIntern(name string, totalAllowance int) (*intern, error) {
	return &intern{totalAllowance: totalAllowance, name: name}, nil
}
