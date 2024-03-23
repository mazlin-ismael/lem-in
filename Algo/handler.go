package algo

import (
	handler "lem-in/Handler"
)

func Handler(farmBase handler.FarmProperties) ([][]string, []int) {
	initFarm(farmBase)
	farm.initRelations()
	farm.InitStepsToEnd()
	farm.initPaths()
	bestsCombs := farm.optimalPaths()
	return(movingAnts(bestsCombs))
}