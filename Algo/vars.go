package algo

import (
	handler "lem-in/Handler"
)

type FarmProperties handler.FarmProperties

var farm FarmProperties

var paths [][]string

type Ant struct {
	Pos 	int
	PathNum	int
	Rank 	int
}