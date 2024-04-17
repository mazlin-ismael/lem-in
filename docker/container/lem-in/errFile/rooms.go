package errFile

import (
	"errors"
	"strconv"
	"strings"
)

// Check if the endpoints start and end are valids
func checkEndpoints() error {
	var countsEndpoints = []int{0, 0}
	// Loop for finding the endpoints
	for pos, row := range farm.FileRows {
		if row == "##start" {
			countsEndpoints[0]++
			farm.Start.Row = pos
		} else if row == "##end" {
			countsEndpoints[1]++
			farm.End.Row = pos
		}
	}
	// Check if there is exactly one start and one end
	if countsEndpoints[0] != 1 || countsEndpoints[1] != 1 {
		return errors.New("invalid endpoints")
	}
	return nil
}

// Initialize the rooms in the farm
func initRooms() error {
	for pos, row := range farm.FileRows {
		rowSplit := strings.Fields(row)

		// Skipping the comments and rows starting with 'L'
		if len(rowSplit) == 0 || rowSplit[0][0] == '#' || rowSplit[0][0] == 'L' {
			continue
		}

		// Check if the row has room properties
		if len(rowSplit) == 3 {
			x, errX := strconv.Atoi(rowSplit[1])
			y, errY := strconv.Atoi(rowSplit[2])
			if errX != nil || errY != nil {
				return errors.New("invalid room")
			}

			// Check if the room is duplicated
			_, exist := farm.Rooms[rowSplit[0]]
			if exist {
				return errors.New("duplicated room")
			}
			// Add room to farm
			farm.Rooms[rowSplit[0]] = &Room{x, y, pos, rowSplit[0], nil, nil, 0, 0}
		}
	}
	return nil
}

// Initialize endpoints start and end
func initEndpoints() error {
	for nameRoom, room := range farm.Rooms {
		if (room.Row)-1 == farm.Start.Row {
			farm.Start.Name = nameRoom
		} else if (room.Row)-1 == farm.End.Row {
			farm.End.Name = nameRoom
		}
	}
	if farm.End.Name == "" || farm.Start.Name == "" {
		return errors.New("invalid endpoint")
	}
	return nil
}