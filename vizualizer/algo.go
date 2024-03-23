package vizualizer

import (
	farm "lem-in/Handler"
)

func multiplicatorsInit(rooms map[string]*farm.Room) (float64, float64) {
	var multiX, multiY int
	var first bool = true

	for _, room := range rooms {
		if first {
			multiX, multiY = room.X, room.Y
			first = false
		} else if room.X > multiX {
			multiX = room.X
		} else if room.Y > multiY {
			multiY = room.Y
		}
	}
	return 100/float64(multiX), 100/float64(multiY)
}

func initNewsRooms(roomsFarm map[string]*farm.Room, multiX, multiY float64) []Room {
	var rooms []Room
	for name, roomFarm := range roomsFarm {
		rooms = append(rooms, Room{
			Name: name,
			X:    int(float64(roomFarm.X) * multiX),
			Y:    int(float64(roomFarm.Y) * multiY),
		})
	}
	return rooms
}