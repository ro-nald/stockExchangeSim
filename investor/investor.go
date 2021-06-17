package investor

import (
	"stockExchangeSim/portfolio"
)

type investor struct {
	name      string
	portfolio portfolio.Portfolio
}

func NewInvestor(name string) *investor {
	newInvestorInstance := new(investor)
	newInvestorInstance.name = name
	return newInvestorInstance
}

func (investor *investor) GetName() string {
	return investor.name
}

func (investor *investor) setupPortfolio(newPortfolio portfolio.Portfolio) {
	investor.portfolio = newPortfolio
}

func (investor *investor) GetPortfolio() portfolio.Portfolio {
	return investor.portfolio
}
