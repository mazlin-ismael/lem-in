package errFile

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func contentFile() []string {
	fileContent, notFound := os.ReadFile("../FILES/fileForTests.txt")
	if notFound != nil {
		return nil
	}
	var rows []string = strings.Split(string(fileContent), "\n")
	return rows
}
func TestGetFile(t *testing.T) {
	os.Args = []string{}
	if getFile().Error() != "incorrect format args" {
		t.Fail()
	}

	os.Args = []string{".", "more", "2 args"}
	if getFile().Error() != "incorrect format args" {
		t.Fail()
	}

	os.Args = []string{".", "fileNotFound"}
	if getFile().Error() != "file not found" {
		t.Fail()
	}

	os.Args = []string{".", "../FILES/example00.txt"}
	if getFile() != nil {
		t.Fail()
	}
}

func TestCheckEndpoints(t *testing.T) {
	rows := contentFile()
	if rows == nil {
		t.Fail()
	}
	rowsCopy1 := slices.Clone(rows)
	rowsCopy2 := slices.Clone(rows)

	farm.FileRows = slices.Delete(rowsCopy1[1:], 0, 1)
	if checkEndpoints().Error() != "invalid endpoints" {
		t.Fail()
	}

	farm.FileRows = slices.Delete(rowsCopy2[1:], 4, 5)
	if checkEndpoints().Error() != "invalid endpoints" {
		t.Fail()
	}

	farm.FileRows = rows[1:]
	if checkEndpoints() != nil {
		t.Fail()
	}
}

func TestInitRooms(t *testing.T) {
	rows := contentFile()
	if rows == nil {
		t.Fail()
	}
	farm.FileRows = rows[1:]

	farm.FileRows[1] = "2 2 errorY"
	if initRooms().Error() != "invalid room" {
		t.Fail()
	}

	farm.Rooms = make(map[string]*Room)
	farm.FileRows[1] = "2 errorX 2"
	if initRooms().Error() != "invalid room" {
		t.Fail()
	}

	farm.Rooms = make(map[string]*Room)
	farm.FileRows[1] = "2 errorX errorY"
	if initRooms().Error() != "invalid room" {
		t.Fail()
	}

	farm.Rooms = make(map[string]*Room)
	farm.FileRows[1] = "0 0 3"
	farm.FileRows[2] = "0 0 3"
	if initRooms().Error() != "duplicated room" {
		t.Fail()
	}

	farm.Rooms = make(map[string]*Room)
	farm.FileRows[2] = "2 2 5"
	if initRooms() != nil {
		t.Fail()
	}
}

func TestInitEndpoints(t *testing.T) {
	start := farm.Rooms["0"]
	end := farm.Rooms["1"]

	delete(farm.Rooms, "0")
	if initEndpoints().Error() != "invalid endpoint" {
		t.Fail()
	}

	farm.Rooms["0"] = start
	farm.End.Name = ""
	delete(farm.Rooms, "1")
	if initEndpoints().Error() != "invalid endpoint" {
		t.Fail()
	}

	farm.Start.Name = ""
	farm.End.Name = ""
	delete(farm.Rooms, "0")
	if initEndpoints().Error() != "invalid endpoint" {
		t.Fail()
	}

	farm.Rooms["0"] = start
	farm.Rooms["1"] = end
	if initEndpoints() != nil {
		t.Fail()
	}
}

func TestInitLinks(t *testing.T) {
	farm.FileRows[6] = "0-8"
	if initLinks().Error() != "bad format links" {
		t.Fail()
	}

	farm.Links = nil
	farm.FileRows[6] = "8-0"
	if initLinks().Error() != "bad format links" {
		t.Fail()
	}

	farm.Links = nil
	farm.FileRows[6] = "8-"
	if initLinks().Error() != "bad format links" {
		t.Fail()
	}

	farm.Links = nil
	farm.FileRows[6] = " -2"
	if initLinks().Error() != "bad format links" {
		t.Fail()
	}

	farm.Links = nil
	farm.FileRows[6] = "0-2"
	farm.FileRows[7] = "0-2"
	if initLinks().Error() != "duplicated link" {
		t.Fail()
	}

	farm.Links = nil
	farm.FileRows[7] = "2-0"
	if initLinks().Error() != "duplicated link" {
		t.Fail()
	}

	farm.Links = nil
	farm.FileRows[7] = "2-3"
	if initLinks() != nil {
		t.Fail()
	}
}

func TestCheckLengthFile(t *testing.T) {
	countRows = 0
	initRooms()
	initLinks()
	if checkLengthFile() != nil {
		t.Fail()
	}

	farm.FileRows = append(farm.FileRows, "")
	countRows = 0
	initRooms()
	initLinks()
	if checkLengthFile().Error() != "invalid rows" {
		t.Fail()
	}

	farm.FileRows = slices.Replace(farm.FileRows, len(farm.FileRows)-1, len(farm.FileRows), "#comm")
	countRows = 0
	initRooms()
	initLinks()
	if checkLengthFile() != nil {
		t.Fail()
	}

	farm.FileRows = slices.Replace(farm.FileRows, len(farm.FileRows)-1, len(farm.FileRows), "\n")
	countRows = 0
	initRooms()
	initLinks()
	if checkLengthFile().Error() != "invalid rows" {
		t.Fail()
	}

	farm.FileRows = slices.Replace(farm.FileRows, len(farm.FileRows)-1, len(farm.FileRows), "L6 8 8")
	countRows = 0
	initRooms()
	initLinks()
	if checkLengthFile().Error() != "invalid rows" {
		t.Fail()
		return
	}
}

func TestNumberAnts(t *testing.T) {
	firstline = "-1"
	if numberAnts().Error() != "ants number can't be zero or negetive" {
		t.Fail()
	}

	firstline = "0"
	if numberAnts().Error() != "ants number can't be zero or negetive" {
		t.Fail()
	}

	firstline = "str"
	if numberAnts().Error() != "bad ants number format" {
		t.Fail()
	}

	firstline = "\n"
	if numberAnts().Error() != "bad ants number format" {
		t.Fail()
	}

	firstline = ""
	if numberAnts().Error() != "bad ants number format" {
		t.Fail()
	}

	firstline = "10"
	if numberAnts() != nil {
		t.Fail()
	}
}