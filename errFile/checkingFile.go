package errFile

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

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
	if numberAnts == 0 {
		return errors.New("ants number can't be zero")
	}
	farm.Ants = numberAnts
	return nil
}