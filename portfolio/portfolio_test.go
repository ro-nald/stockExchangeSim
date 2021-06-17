package portfolio

import "testing"

func TestInvestment(t *testing.T) {

	testInvestmentName := "Rosta Coffee"
	testInvestmentUnitPrice := 67.324
	testInvestmentQuantity := 500

	investment := NewInvestment(testInvestmentName, testInvestmentUnitPrice, testInvestmentQuantity)

	if investment.name != testInvestmentName {
		t.Errorf("Expected \"%v\", got \"%v\"", testInvestmentName, investment.name)
	}
	if investment.unitPrice != testInvestmentUnitPrice {
		t.Errorf("Expected \"%v\", got \"%v\"", testInvestmentUnitPrice, investment.unitPrice)
	}
	if investment.quantity != testInvestmentQuantity {
		t.Errorf("Expected \"%v\", got \"%v\"", testInvestmentQuantity, investment.quantity)
	}
}

func TestPortfolio(t *testing.T) {
	testInvestmentName := "Rosta Coffee"
	testInvestmentUnitPrice := 67.324
	testInvestmentQuantity := 500

	investment := NewInvestment(testInvestmentName, testInvestmentUnitPrice, testInvestmentQuantity)

	portfolioName := "Ronald's portfolio"
	portfolioContents := []Investment{*investment}

	portfolio := NewPortfolio(portfolioName, portfolioContents)

	portfolioInvestments := portfolio.GetPortfolio()
	value, _ := portfolioInvestments["Rosta Coffee"]

	if value.name != testInvestmentName {
		t.Errorf("Expected \"%v\", got \"%v\"", testInvestmentName, value.name)
	}
	if value.unitPrice != testInvestmentUnitPrice {
		t.Errorf("Expected \"%v\", got \"%v\"", testInvestmentUnitPrice, value.unitPrice)
	}
	if value.quantity != testInvestmentQuantity {
		t.Errorf("Expected \"%v\", got \"%v\"", testInvestmentQuantity, value.quantity)
	}
}
