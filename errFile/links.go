package errFile

import (
	"errors"
	"strings"
)

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
	return "", "", errors.New("bad format links")
}

func duplicatedLink(name1, name2 string) error {
	for _, link := range farm.Links {
		if name1 == link[0] && name2 == link[1] || name1 == link[1] && name2 == link[0] {
			return errors.New("duplicated link" + " " + link[0] + " "+ link[1])
		}
	}
	return nil
}