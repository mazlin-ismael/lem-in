package main

import (
	"fmt"
	algo "lem-in/algo"
	errFile "lem-in/errFile"
	vizu2d "lem-in/visualizer"
)

func main() {
	farmInit := errFile.Handler()
	optimalComb, antsByPaths := algo.Handler(farmInit)
	fmt.Println(optimalComb, antsByPaths)
	vizu2d.WebHandler(farmInit, optimalComb, antsByPaths)
}