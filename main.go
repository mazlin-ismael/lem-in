package main

import (
	handler "lem-in/Handler"
	algo "lem-in/Algo"
)

func main() {
	farmInit := handler.Handler()
	algo.Handler(farmInit)
}