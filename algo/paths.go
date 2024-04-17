package algo

import (
	errFile "lem-in/errFile"
	"math"
	"reflect"
)

// Initialize the paths in the farm
func (farm *FarmProperties) initPaths() {
	start := farm.Rooms[farm.Start.Name]
	end := farm.Rooms[farm.End.Name]
	current := start
	var prevRoom *errFile.Room

	// Reset PrevRoom for all rooms
	for _, room := range farm.Rooms {
		room.PrevRoom = nil
	}

	// Go through all rooms to find paths to the end
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

// savePaths saves the paths found during initPaths
func savePaths(current *errFile.Room) {
	var path []string
	var inArray bool
		for current != nil {
			path = append([]string{current.Name}, path...)
			current = current.PrevRoom
		}

		// Check if the path already exists in paths
		for _, savePath := range paths {
			if reflect.DeepEqual(savePath, path) {
				inArray = true
			}
		}
		// If not, append it to paths
		if !inArray {
			paths = append(paths, path)
		}
}

// Find the optimal paths based on farm properties
func (farm *FarmProperties) optimalPaths() [][][]string {
	// Determine the maximum number of paths to consider
	maxPath := int(math.Min(float64(len(farm.Rooms[farm.Start.Name].LinkedRooms)), float64(len(farm.Rooms[farm.End.Name].LinkedRooms))))
	if farm.Ants < maxPath {
		maxPath = farm.Ants
	}
	
	// Initialize combinations and bestCombPaths
	combsPaths, bestCombsPaths := initFirstComb(paths)
	var bestCombPaths [][]string

	for i := 1; i < maxPath; i++ {
		combsPaths, bestCombPaths = initCombs(combsPaths, paths)
		bestCombsPaths = append(bestCombsPaths, bestCombPaths)
	}
	
	return bestCombsPaths
}