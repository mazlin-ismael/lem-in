package Handler

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func getFile() error {
	if len(os.Args) != 2 {
		return errors.New("unspecified file")
	}
	file, notFound := os.ReadFile(os.Args[1])
	if notFound != nil {
		return errors.New("file not found")
	}
	rows := strings.Split(string(file), "\n")
	farm.FileRows = rows[1:]
	firstline = rows[0]
	return nil
}

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

func initLinks() error {
	for _, row := range farm.FileRows {
		if len(strings.Fields(row)) == 1 && !(row[0] == '#') {
			var names []string

			for roomName := range farm.Rooms {
				if strings.Contains(row, roomName) {
					names = append(names, roomName)
				}
			}

			name1, name2, badFormat := validLink(row, names)
			if badFormat != nil {
				return badFormat
			}

			validLink := duplicatedLink(name1, name2)
				if validLink != nil {
					return validLink
				}
			farm.Links = append(farm.Links, [2]string{name1, name2})

		}
	}
	countRows = countRows + len(farm.Links)
	return nil
}

func validLink(row string, names []string) (string, string, error) {
	for i, name1 := range names[:len(names)-1] {
		for _, name2 := range names[i+1:] {

			if name1 + "-" + name2 == row || name2 + "-" + name1 == row {
				return name1, name2, nil
			}
		}
	}
	return "", "", errors.New("bad format")
}

func duplicatedLink(name1, name2 string) error {
	for _, link := range farm.Links {
		if name1 == link[0] && name2 == link[1] || name1 == link[1] && name2 == link[0] {
			return errors.New("duplicated link" + " " + link[0] + " "+ link[1])
		}
	}
	return nil
}

func checkLengthFile() error {
	if countRows != len(farm.FileRows) {
		return errors.New("invalid rows")
	}
	return nil
}

func numberAnts() error {
	numberAnts, notInt := strconv.Atoi(firstline)
	if notInt != nil {
		return errors.New("bad ants number format")
	}
	farm.Ants = numberAnts
	return nil
}