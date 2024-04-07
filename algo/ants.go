package algo

import (
	"fmt"
	"slices"
)

func movingAnts(bestsCombs [][][]string) ([][]string, []int) {
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
	bestAntsByPathSaved := slices.Clone(bestAntsByPath)
	displayPathAnts(shortestComb, bestAntsByPath)
	return shortestComb, bestAntsByPathSaved
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
	var rankAnt = 1
	for !endDisplaying {
		
		for i, ant := range antsByPath {
			if ant > 0 {
				ants = append(ants, Ant{
					Pos:	 	0,
					PathNum: 	i,
					Rank: 		rankAnt,
				})
				antsByPath[i]--
				rankAnt++
			}
		}
		
		endDisplaying = true
		for i, ant := range ants {
			if ant.Pos < len(bestComb[ant.PathNum]) {
				if i >= len(ants) {
					break
				}
				fmt.Print("L", ant.Rank, "-", bestComb[ant.PathNum][ant.Pos])
				endDisplaying = false
				ants[i].Pos++
				if i != len(ants)-1 {
					fmt.Print(" ")
				} else {
					fmt.Print("\n")
				}
			} else {
				ants = slices.Delete(ants, i, i+1)
			}
		}
	}
}