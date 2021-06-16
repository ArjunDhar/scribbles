package moduleA

/*
The Helper module code that is generic and will not change from application to application.

@author Arjun Dhar
*/

import model "module-A/model"

var appCalculator Calculator

func SetAppCalculator(c Calculator) {
	appCalculator = c
	// Note : As an alternative one can derive the Calculator singleton, from the container here.
}

func CalcTax(user *model.User) float32 {
	taxableIncomePerMonth := appCalculator.MonthlyAverageIncome(user) - appCalculator.HouseRentAllowance(user)
	tax := float32(taxableIncomePerMonth) * 12 * float32(appCalculator.GetTaxPercent())/100
	return tax
}