package vizualizer

import (
	farm "lem-in/errFile"
)

// Initialize the multiplicator based on the maximum x and y coordinates of room
func multiplicatorsInit(rooms map[string]*farm.Room) (float64, float64) {
	var multiX, multiY int
	var first bool = true

	// Find the maximum x and y coordinates of rooms
	for _, room := range rooms {
		if first {
			multiX, multiY = room.X, room.Y
			first = false

		} else {
			if room.X > multiX {
				multiX = room.X
			}

			if room.Y > multiY {
			multiY = room.Y
			}
		}
	}
	return 100/float64(multiX), 100/float64(multiY)
}

// Initialize the rooms with adjusted coordinates
func initNewsRooms(roomsFarm map[string]*farm.Room, multiX, multiY float64) {
	for name, roomFarm := range roomsFarm {
		rooms[name] = Room{
			Name: name,
			X:    float64(roomFarm.X) * multiX,
			Y:    float64(roomFarm.Y) * multiY,
		}
	}
}

// Initialize links between rooms
func initLinks(farm farm.FarmProperties) {
	for _, link := range farm.Links {
		links = append(links, link)
	}
}

// Initialize endpoints start and end
func initEndpoints(farm farm.FarmProperties) {
	start = farm.Start.Name
	end = farm.End.Name
}

// Initialize paths combinations
func initComb(optimalComb [][]string) {
	selectComb = optimalComb
}

// Initialize the number of ants for each paths
func initAntsByPaths(antsByPaths []int) {
	antsComb = antsByPaths
}

// Initialize the data view for the visualization
func initDataView() DataViews {
	var dataView DataViews = DataViews{
		Rooms:	rooms,
		Links:	links,
		Start: 	start,
		End:	end,
		Comb: 	selectComb,
		Ants:   antsComb,
	}
	return dataView
}