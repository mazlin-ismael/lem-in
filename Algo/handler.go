package algo

import (
	handler "lem-in/Handler"
)

func Handler(farmBase handler.FarmProperties) {
	InitFarm(farmBase)
	farm.InitRelations()
}