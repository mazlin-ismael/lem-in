package algo

import (
	errFile "lem-in/errFile"
)

func Handler(farmBase errFile.FarmProperties) ([][]string, []int) {
	initFarm(farmBase)
	farm.initRelations()
	errFile.CheckFunc(CheckPossiblePath)
	farm.initPaths()
	bestsCombs := farm.optimalPaths()
	return(movingAnts(bestsCombs))
}