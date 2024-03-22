package algo

import (
	"errors"
	"fmt"
	handler "lem-in/Handler"
	"math"
	"reflect"
	"slices"
)

func initFarm(farmBase handler.FarmProperties) {
	farm = FarmProperties(farmBase)
}

func (farm *FarmProperties) initRelations() {
	for _, link := range farm.Links {
		farm.Rooms[link[0]].LinkedRooms = append(farm.Rooms[link[0]].LinkedRooms, farm.Rooms[link[1]])
		farm.Rooms[link[1]].LinkedRooms = append(farm.Rooms[link[1]].LinkedRooms, farm.Rooms[link[0]])
	}
	handler.CheckFunc(checkPossiblePath)
}

func checkPossiblePath() error {
	current := farm.Rooms[farm.Start.Name]
	var prevRoom *handler.Room
	end := farm.Rooms[farm.End.Name]
	var indRoom int

	for current != end {
		if indRoom == len(current.LinkedRooms) {
			indRoom = 0
			if current.PrevRoom == nil {
				return errors.New("no path between start and end")
			}
			current = current.PrevRoom
		}

		if current.LinkedRooms[indRoom].PrevRoom == nil && current.LinkedRooms[indRoom].Name != farm.Start.Name {
			prevRoom = current
			current = current.LinkedRooms[indRoom]
			current.PrevRoom = prevRoom
			indRoom = 0
		} else {
			indRoom++
		}
	}
	return nil
}

func (farm *FarmProperties) InitStepsToEnd() {
	end := farm.Rooms[farm.End.Name]
	var steps int = 1
	var shortsPathsInit bool = true

	for _, room := range end.LinkedRooms {
		room.StepsToEnd = steps
	}

	for shortsPathsInit {
		shortsPathsInit = false

		for _, room := range farm.Rooms {
			if room.StepsToEnd == steps {
				for _, room := range room.LinkedRooms {
					if (room.StepsToEnd > steps+1 || room.StepsToEnd == 0) && room != end {
						room.StepsToEnd = steps + 1
						shortsPathsInit = true
					}
				}
			}
		}
		steps++
	}
}

func (farm *FarmProperties) initPaths() {
	start := farm.Rooms[farm.Start.Name]
	end := farm.Rooms[farm.End.Name]
	current := start
	var prevRoom *handler.Room

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
		current = current.PrevRoom
		current.NextPos++

		var path []string
		var inArray bool
		for savePath != nil {
			path = append([]string{savePath.Name}, path...)
			savePath = savePath.PrevRoom
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

func initFirstComb(alonePaths [][]string) ([][][]string, [][][]string) {
	var combsPaths [][][]string
	var shortestPath int = len(alonePaths[0])
	var bestPath []string = alonePaths[0]
	var bestsCombsPaths [][][]string

	for _, path := range alonePaths {
		combsPaths = append(combsPaths, [][]string{path})

		if len(path) < shortestPath {
			shortestPath = len(path)
			bestPath = path
		}
	}

	bestsCombsPaths = append(bestsCombsPaths, [][]string{bestPath})
	return combsPaths, bestsCombsPaths
}

func initCombs(currentCombs[][][]string, pathsToAdd [][]string) ([][][]string, [][]string) {
	var combsPaths [][][]string
	for i := 0; i < len(pathsToAdd); i++ {
		for j := 0; j < len(currentCombs); j++ {

			var isInComb bool
			for _, pathToAdd := range pathsToAdd[i] {

				for c, currentComb := range currentCombs[j] {
					currentComb = currentComb[1:len(currentComb)-1]
					if slices.Contains(currentComb, pathToAdd) || reflect.DeepEqual(currentCombs[j][c], pathsToAdd[i]) {
						isInComb = true
						break
					}
				}

				if isInComb {
					break
				}
			}
			if !isInComb {
				newComb := append(currentCombs[j], pathsToAdd[i])
				combsPaths = append(combsPaths, newComb)
			}

		}
	}

	return  combsPaths, initBestComb(combsPaths)
}

func initBestComb(combsPaths [][][]string) [][]string {
	var shortestComb int
	var bestComb [][]string

	for i, combPath := range combsPaths {
		var comparedComb int

		for _, path := range combPath {
			if i == 0 {
				shortestComb += len(path)
			} else {
				comparedComb += len(path)
			}
		}

		if i == 0 {
			bestComb = combPath
		} else if comparedComb < shortestComb {
			shortestComb = comparedComb
			bestComb = combPath
		}
	}
	return bestComb
}

func movingAnts(bestsCombs [][][]string) {
	var shortestComb [][]string
	var bestAntsByPath   []int
	for i, bestComb := range bestsCombs {
		combSorted, antsByPath := antsToSend(bestComb)
		if i == 0 {
			shortestComb = combSorted
			bestAntsByPath = antsByPath
		} else if antsByPath[0] + len(combSorted[0])-1 < bestAntsByPath[0] + len(shortestComb[0])-1 {
			shortestComb = combSorted
			bestAntsByPath = antsByPath
		}
	}
	displayPathAnts(shortestComb, bestAntsByPath)
}

func antsToSend(comb [][]string) (combSorted [][]string, antsByPath []int){
	for i := 0; i < len(comb)-1; i++ {
		for j := i+1; j < len(comb); j++ {
			if len(comb[i]) > len(comb[j]) {
				comb[i], comb[j] = comb[j], comb[i]
			}
		}
	}

	var gapLinks []int
	for _, path := range comb {
		gapLinks = append(gapLinks, len(path) - len(comb[0]))
	}

	var count int
	var numbersAnts []int = make([]int, len(gapLinks))

	for i := 0; i < farm.Ants; {
		for j, gapLink := range gapLinks {
			if gapLink <= count {
				numbersAnts[j]++
				i++
			}
			if i == farm.Ants {
				return comb, numbersAnts
			}
		}
		count++
	}
	return comb, numbersAnts
}

func displayPathAnts(bestComb [][]string, antsByPath []int) {
	var ants []Ant
	var endDisplaying bool

	for i := range bestComb {
		bestComb[i] = bestComb[i][1:]
	}

	for !endDisplaying {
		for i, ant := range antsByPath {
			if ant > 0 {
				ants = append(ants, Ant{
					Pos: 0,
					PathNum: i,
				})
				antsByPath[i]--
			}
		}

		endDisplaying = true
		for i, ant := range ants {
			if ant.Pos < len(bestComb[ant.PathNum]) {
				fmt.Print("L", i+1, "-", bestComb[ant.PathNum][ant.Pos])
				endDisplaying = false
				ants[i].Pos++
				if i != len(ants)-1 {
					fmt.Print(" ")
				} else {
					fmt.Print("\n")
				}
			}
		}
	}

}