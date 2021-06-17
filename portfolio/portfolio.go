package portfolio

type Investment struct {
	name      string
	unitPrice float64
	quantity  int
}

type Portfolio struct {
	owner       string
	investments map[string]Investment
}

func NewInvestment(investmentName string, unitPrice float64, quantity int) *Investment {
	newInvestment := new(Investment)
	newInvestment.name = investmentName
	newInvestment.unitPrice = unitPrice
	newInvestment.quantity = quantity
	return newInvestment
}

func NewPortfolio(ownerName string, investmentPortfolio []Investment) *Portfolio {
	newPortfolio := new(Portfolio)
	newPortfolio.owner = ownerName
	newPortfolio.investments = make(map[string]Investment)
	for _, investment := range investmentPortfolio {
		newPortfolio.investments[investment.name] = investment
	}
	return newPortfolio
}

func (portfolio *Portfolio) AddInvestment(newInvestment Investment) {
	portfolio.investments[newInvestment.name] = newInvestment
}

func (portfolio *Portfolio) RemoveInvestment(investment Investment) {
	delete(portfolio.investments, investment.name)
}

func (portfolio *Portfolio) GetPortfolio() map[string]Investment {
	return portfolio.investments
}
