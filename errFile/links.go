package errFile

import (
	"errors"
	"strings"
)

// Initialize the links between rooms of the farm
func initLinks() error {
	for _, row := range farm.FileRows {
		// Check if the row is not a comment and is only one field
		if len(strings.Fields(row)) == 1 && !(row[0] == '#') {
			var names []string

			// Find rooms names of the current link
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
	return nil
}

// Check if the link format is good
func validLink(row string, names []string) (string, string, error) {
	if len(names)-1 < 0 {
		return "", "", errors.New("bad format links")
	}
	
	for i, name1 := range names[:len(names)-1] {
		for _, name2 := range names[i+1:] {

			if name1 + "-" + name2 == row || name2 + "-" + name1 == row {
				return name1, name2, nil
			}
		}
	}
	return "", "", errors.New("bad format links")
}

// Check if the link is not duplicated
func duplicatedLink(name1, name2 string) error {
	for _, link := range farm.Links {
		if name1 == link[0] && name2 == link[1] || name1 == link[1] && name2 == link[0] {
			return errors.New("duplicated link")
		}
	}
	return nil
}