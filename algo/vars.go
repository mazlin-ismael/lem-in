package algo

import (
	errFile "lem-in/errFile"
)

type FarmProperties errFile.FarmProperties

var farm FarmProperties

var paths [][]string

type Ant struct {
	Pos 	int
	PathNum	int
	Rank 	int
}