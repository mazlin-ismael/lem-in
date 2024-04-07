package errFile

func checkingFile() {
	CheckFunc(getFile)
	CheckFunc(checkEndpoints)
	CheckFunc(initRooms)
	CheckFunc(initEndpoints)
	CheckFunc(initLinks)
	CheckFunc(checkLengthFile)
	CheckFunc(numberAnts)
}

func Handler() FarmProperties {
	checkingFile()
	return farm
}
