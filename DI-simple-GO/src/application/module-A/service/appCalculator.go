package moduleA

/*
App specific implementation of Calculator

@author Arjun Dhar
*/

//import serviceDef "module-A/service" //Service Def
import model "module-A/model" // Model Def

type AppCalculator struct {
}

func (c *AppCalculator) MonthlyAverageIncome(user *model.User) int {
	return 1000
}

func (c *AppCalculator) HouseRentAllowance(user *model.User) int {
	return 100
}

func (c *AppCalculator) GetTaxPercent() int {
	return 30
}
