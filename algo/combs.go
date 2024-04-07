package algo

import (
	"reflect"
	"slices"
)


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