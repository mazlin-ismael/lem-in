package main

import (
	algo "lem-in/algo"
	errFile "lem-in/errFile"
	vizu2d "lem-in/visualizer"
)

func main() {
	farmInit := errFile.Handler()
	optimalComb, antsByPaths := algo.Handler(farmInit)
	vizu2d.WebHandler(farmInit, optimalComb, antsByPaths)
}