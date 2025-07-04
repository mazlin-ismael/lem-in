package algo

import (
	"errors"
	errFile "lem-in/errFile"
)

// Initialize the farm with the provided base properties
func initFarm(farmBase errFile.FarmProperties) {
	farm = FarmProperties(farmBase)
}

// Initialize the relations between the rooms in the farm
func (farm *FarmProperties) initRelations() {
	for _, link := range farm.Links {
		farm.Rooms[link[0]].LinkedRooms = append(farm.Rooms[link[0]].LinkedRooms, farm.Rooms[link[1]])
		farm.Rooms[link[1]].LinkedRooms = append(farm.Rooms[link[1]].LinkedRooms, farm.Rooms[link[0]])
	}
}

// Check if there is a possible path between the endpoints start and end
func checkPossiblePath() error {
	current := farm.Rooms[farm.Start.Name]
	var prevRoom *errFile.Room
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