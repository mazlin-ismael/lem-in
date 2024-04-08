package main

import (
	"os"
	algo "lem-in/algo"
	errFile "lem-in/errFile"
	vizu2d "lem-in/visualizer"
)

func main() {
	farmInit := errFile.Handler()
	optimalComb, antsByPaths := algo.Handler(farmInit)
	if len(os.Args) == 2 {
		vizu2d.WebHandler(farmInit, optimalComb, antsByPaths)
	} 
}