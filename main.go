package main

import (
	handler "lem-in/Handler"
	algo "lem-in/Algo"
	vizu2d   "lem-in/vizualizer"
)

func main() {
	farmInit := handler.Handler()
	algo.Handler(farmInit)
	vizu2d.Handler()
}