package Handler

type FarmProperties struct {
	FileRows []string
	Start    Endpoint
	End      Endpoint
	Links    [][2]string
	Rooms    map[string]*Room
	Ants	int
}

type Room struct {
	x 			int
	y 			int
	Row 		int
	Name 		string
	LinkedRooms []*Room
	PrevRoom	*Room
	StepToEnd	int
}

type Endpoint struct {
	Name	string
	Row		int
}


var farm FarmProperties = FarmProperties{
	Rooms: make(map[string]*Room),
}

var countRows int
var firstline string

