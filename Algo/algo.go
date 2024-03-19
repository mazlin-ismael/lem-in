package algo

import (
	"errors"
	"fmt"
	handler "lem-in/Handler"
	"reflect"
	"math"
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
	actual := farm.Rooms[farm.Start.Name]
	var prevRoom *handler.Room
	end := farm.Rooms[farm.End.Name]
	var indRoom int

	for actual != end {
		if indRoom == len(actual.LinkedRooms) {
			indRoom = 0
			if actual.PrevRoom == nil {
				return errors.New("no path between start and end")
			}
			actual = actual.PrevRoom
		}
		
		if actual.LinkedRooms[indRoom].PrevRoom == nil && actual.LinkedRooms[indRoom].Name != farm.Start.Name {
			prevRoom = actual
			actual = actual.LinkedRooms[indRoom]
			actual.PrevRoom = prevRoom
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

func (farm *FarmProperties) optimalPaths() {
	maxPath := math.Min(float64(len(farm.Rooms[farm.Start.Name].LinkedRooms)), float64(len(farm.Rooms[farm.End.Name].LinkedRooms)))
	for _, path := range paths {
		fmt.Println(path)
	}
	fmt.Println(maxPath)
}