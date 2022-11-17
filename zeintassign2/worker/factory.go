package worker

type WorkerFactory struct{}

func (WorkerFactory) NewHire(name string, totalPerDiem int) (*perm, error) {
	return &perm{totalPerDiem: totalPerDiem, name: name}, nil
}
func (WorkerFactory) NewContract(name string, totalCashAdvance int) (*contract, error) {
	return &contract{totalCashAdvance: totalCashAdvance, name: name}, nil
}
func (WorkerFactory) NewIntern(name string, totalAllowance int) (*intern, error) {
	return &intern{totalAllowance: totalAllowance, name: name}, nil
}
