package errFile

// All-in-one check error function
func checkingFile() {
	CheckFunc(getFile)
	CheckFunc(checkEndpoints)
	CheckFunc(initRooms)
	CheckFunc(initEndpoints)
	CheckFunc(initLinks)
	CheckFunc(checkLengthFile)
	CheckFunc(numberAnts)
}

// Launch checkingFile function and return farm
func Handler() FarmProperties {
	checkingFile()
	return farm
}
