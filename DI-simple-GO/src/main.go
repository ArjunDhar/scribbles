package main

import "fmt"
import application "application/module-A"
import modAServices "module-A/service"
import modAModel "module-A/model" // Model Def

func main() {
	application.Init()
	tax := modAServices.CalcTax(&modAModel.User{
		Name: "Arjun",
		ID:   "1",
	})
	fmt.Println(tax) // 3240
}
