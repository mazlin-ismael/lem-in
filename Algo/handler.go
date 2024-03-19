package algo

import (
	handler "lem-in/Handler"
)

func Handler(farmBase handler.FarmProperties) {
	initFarm(farmBase)
	farm.initRelations()
	farm.InitStepsToEnd()
	farm.initPaths()
	farm.optimalPaths()
}