package errFile

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Get the argument file, create a version split by rows + check the file format
func getFile() error {
	if len(os.Args) != 2 {
		return errors.New("incorrect format args")
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

// Check the number of rows
func checkLengthFile() error {
	countRows = countRows + len(farm.Links)
	countRows = countRows + len(farm.Rooms)

	for _, row := range farm.FileRows { 
		if len(row) > 0 && row[0] == '#' {
			countRows++
		}
	}
	
	if countRows != len(farm.FileRows) {
		return errors.New("invalid rows")
	}
	return nil
}

// Get the number of ants and check if it is valid
func numberAnts() error {
	numberAnts, notInt := strconv.Atoi(firstline)
	if notInt != nil {
		return errors.New("bad ants number format")
	}
	if numberAnts <= 0 {
		return errors.New("ants number can't be zero or negetive")
	}
	farm.Ants = numberAnts
	return nil
}