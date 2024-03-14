package algo

import (
	"errors"
	"fmt"
	handler "lem-in/Handler"
)

func InitFarm(farmBase handler.FarmProperties) {
	farm = FarmProperties(farmBase)
}

func (farm *FarmProperties) InitRelations() {
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

func (farm *FarmProperties) RoomsStepToEnd() {
	actual := farm.Rooms[farm.End.Name]
	farm.Rooms[farm.End.Name].PrevRoom = nil
	var prevRoom *handler.Room
	var indRoom int
	var stepsToEnd int = 0

	for {
		if indRoom >= len(actual.LinkedRooms) {
			indRoom = 0
			stepsToEnd--
			if stepsToEnd == -1 {
				break
			}
			actual = actual.PrevRoom
		}
		
		if (actual.LinkedRooms[indRoom].StepToEnd == 0 || actual.LinkedRooms[indRoom].StepToEnd > stepsToEnd) && actual.LinkedRooms[indRoom] != actual.PrevRoom {
			fmt.Println(actual.Name)
			stepsToEnd++
			prevRoom = actual
			actual = actual.LinkedRooms[indRoom]
			actual.PrevRoom = prevRoom
			actual.StepToEnd = stepsToEnd
		} else {
			indRoom++
		}
	}
}	