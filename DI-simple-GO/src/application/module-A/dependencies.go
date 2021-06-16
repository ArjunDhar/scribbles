package app

import modAServices "module-A/service"
import appServices "application/module-A/service"

/*
This file injects the application specific dependencies into the Helpers / Services

@author Arjun Dhar
*/

func Init() {
	modAServices.SetAppCalculator(&appServices.AppCalculator{})
}
