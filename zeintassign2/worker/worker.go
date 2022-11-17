package worker

// someone who spends $$$
type MoneySpender interface {
	// monthly expenses of company members
	GetMonthlyExpenses() int
	// self explanatory
	GetName() string
}

// a permanent employee
type perm struct {
	name         string
	totalPerDiem int
}

// a contractor
type contract struct {
	name             string
	totalCashAdvance int
}

// an intern
type intern struct {
	name           string
	totalAllowance int
}

func (p perm) GetMonthlyExpenses() int {
	return p.totalPerDiem
}
func (c contract) GetMonthlyExpenses() int {
	return c.totalCashAdvance
}
func (i intern) GetMonthlyExpenses() int {
	return i.totalAllowance
}

func (p perm) GetName() string {
	return p.name
}
func (c contract) GetName() string {
	return c.name
}
func (i intern) GetName() string {
	return i.name
}
