package controller

import "lalokal/infrastructure/injector"

type mainController struct {
	solvent injector.InjectorSolvent
}

func MainController(solvent *injector.InjectorSolvent) mainController {
	return mainController{solvent: *solvent}
}
