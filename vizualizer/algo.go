package vizualizer

import (
	"lem-in/Handler"
	farm "lem-in/Handler"
)

func multiplicatorsInit(rooms map[string]*farm.Room) (float64, float64) {
	var multiX, multiY int
	var first bool = true

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

func initNewsRooms(roomsFarm map[string]*farm.Room, multiX, multiY float64) {
	for name, roomFarm := range roomsFarm {
		rooms[name] = Room{
			Name: name,
			X:    float64(roomFarm.X) * multiX,
			Y:    float64(roomFarm.Y) * multiY,
		}
	}
}

func initLinks(farm farm.FarmProperties) {
	for _, link := range farm.Links {
		links = append(links, link)
	}
}

func initEndpoints(farm Handler.FarmProperties) {
	start = farm.Start.Name
	end = farm.End.Name
}