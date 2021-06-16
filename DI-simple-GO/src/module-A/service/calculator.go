package moduleA

/*
Interface to allow DI

@author Arjun Dhar
*/

import model "module-A/model"

type Calculator interface {
	MonthlyAverageIncome(user *model.User) int
	HouseRentAllowance(user *model.User) int
	GetTaxPercent() int
}
