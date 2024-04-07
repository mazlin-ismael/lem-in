package errFile

import (
	"errors"
	"strconv"
	"strings"
)

func checkEndpoints() error {
	var countsEndpoints = []int{0, 0}
	for pos, row := range farm.FileRows {
		if row == "##start" {
			countsEndpoints[0]++
			farm.Start.Row = pos
		} else if row == "##end" {
			countsEndpoints[1]++
			farm.End.Row = pos
		}
	}
	if countsEndpoints[0] != 1 || countsEndpoints[1] != 1 {
		return errors.New("invalid endpoints")
	}
	return nil
}

func initRooms() error {
	for pos, row := range farm.FileRows {
		rowSplit := strings.Fields(row)

		if len(rowSplit) == 0 || rowSplit[0][0] == '#' || rowSplit[0][0] == 'L' {
			if len(rowSplit) > 0 && rowSplit[0][0] == '#' {
				countRows++
			}
			continue
		}

		if len(rowSplit) == 3 {
			x, errX := strconv.Atoi(rowSplit[1])
			y, errY := strconv.Atoi(rowSplit[2])
			if errX != nil || errY != nil {
				return errors.New("invalid room")
			}

			_, exist := farm.Rooms[rowSplit[0]]
			if exist {
				return errors.New("duplicated room")
			}
			farm.Rooms[rowSplit[0]] = &Room{x, y, pos, rowSplit[0], nil, nil, 0, 0}
		}
	}
	countRows = countRows + len(farm.Rooms)
	return nil
}

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