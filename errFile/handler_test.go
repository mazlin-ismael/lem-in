package errFile

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func contentFile() []string {
	fileContent, notFound := os.ReadFile("../FILES/test.txt")
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

	farm.FileRows[1] = "2 errorX 2"
	if initRooms().Error() != "invalid room" {
		t.Fail()
	}
	
	farm.FileRows[1] = "2 errorX errorY"
	if initRooms().Error() != "invalid room" {
		t.Fail()
	}

	farm.FileRows[1] = "0 0 3"
	farm.FileRows[2] = "0 0 3"
	if initRooms().Error() != "duplicated room" {
		t.Fail()
	}
}
