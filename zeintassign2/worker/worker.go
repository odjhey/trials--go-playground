package worker

type MoneyReceiver interface {
	GetMonthlyExpenses() int
	GetName() string
}

type perm struct {
	name         string
	totalPerDiem int
}
type contract struct {
	name             string
	totalCashAdvance int
}
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
