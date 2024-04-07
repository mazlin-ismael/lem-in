package main

import (
	errFile "lem-in/errFile"
	algo "lem-in/algo"
	vizu2d   "lem-in/vizualizer"
)

func main() {
	farmInit := errFile.Handler()
	optimalComb, antsByPaths := algo.Handler(farmInit)
	vizu2d.WebHandler(farmInit, optimalComb, antsByPaths)
}