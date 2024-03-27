package Handler

type FarmProperties struct {
	FileRows []string
	Start    Endpoint
	End      Endpoint
	Links    [][2]string
	Rooms    map[string]*Room
	Ants     int
}

type Room struct {
	X           int
	Y           int
	Row         int
	Name        string
	LinkedRooms []*Room
	PrevRoom    *Room
	NextPos     int
	StepsToEnd  int
}

type Endpoint struct {
	Name string
	Row  int
}

var farm FarmProperties = FarmProperties{
	Rooms: make(map[string]*Room),
}

var (
	countRows int
	firstline string
)
