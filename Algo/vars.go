package algo

import (
	handler "lem-in/Handler"
)

type FarmProperties handler.FarmProperties

var farm FarmProperties
var visitedRoom map[*handler.Room]handler.Room
