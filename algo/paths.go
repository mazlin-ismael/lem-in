package algo

import (
	errFile "lem-in/errFile"
	"math"
	"reflect"
)

func (farm *FarmProperties) initPaths() {
	start := farm.Rooms[farm.Start.Name]
	end := farm.Rooms[farm.End.Name]
	current := start
	var prevRoom *errFile.Room

	for _, room := range farm.Rooms {
		room.PrevRoom = nil
	}

	for {
		for current != end {
			for current.NextPos >= len(current.LinkedRooms) {
				current = current.PrevRoom
				if current == nil {
					return
				}
				current.LinkedRooms[current.NextPos].PrevRoom = nil
				current.NextPos++
			}

			if current.LinkedRooms[current.NextPos].PrevRoom == nil && current.LinkedRooms[current.NextPos] != start || current.LinkedRooms[current.NextPos] == end {
				prevRoom = current
				current = current.LinkedRooms[current.NextPos]
				current.PrevRoom = prevRoom
				current.NextPos = 0
			} else {
				current.NextPos++
			}
		}

		savePath := current
		savePaths(savePath)

		current = current.PrevRoom
		current.NextPos++

	}
}

func savePaths(current *errFile.Room) {
	var path []string
	var inArray bool
		for current != nil {
			path = append([]string{current.Name}, path...)
			current = current.PrevRoom
		}

		for _, savePath := range paths {
			if reflect.DeepEqual(savePath, path) {
				inArray = true
			}
		}
		if !inArray {
			paths = append(paths, path)
		}
}

func (farm *FarmProperties) optimalPaths() [][][]string {
	maxPath := int(math.Min(float64(len(farm.Rooms[farm.Start.Name].LinkedRooms)), float64(len(farm.Rooms[farm.End.Name].LinkedRooms))))
	if farm.Ants < maxPath {
		maxPath = farm.Ants
	}
	
	combsPaths, bestCombsPaths := initFirstComb(paths)
	var bestCombPaths [][]string

	for i := 1; i < maxPath; i++ {
		combsPaths, bestCombPaths = initCombs(combsPaths, paths)
		bestCombsPaths = append(bestCombsPaths, bestCombPaths)
	}
	
	return bestCombsPaths
}