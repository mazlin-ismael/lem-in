package algo

import (
	errFile "lem-in/errFile"
	"os"
	"slices"
	"testing"
)

func TestCheckPossiblePath(t *testing.T) {

	os.Args = []string{".", "../FILES/fileForTests.txt"}
	initFarm(errFile.Handler())
	
	links := slices.Clone(farm.Links)
	farm.Links = append(farm.Links[0:1], farm.Links[2:]...)
	farm.initRelations()

	if checkPossiblePath().Error() != "no path between start and end" {
		t.Fail()
	}

	farm.Links = links
	for _, room := range farm.Rooms {
		room.LinkedRooms = nil
		room.PrevRoom = nil
	}

	farm.initRelations()
	if checkPossiblePath() != nil {
		t.Fail()
	}
}