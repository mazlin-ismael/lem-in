package vizualizer

type Room struct {
	Name 	string
	X		int
	Y		int
}

type DataViews struct {
	Rooms	map[string]Room
	Links	[][2]string
}

var rooms map[string]Room = make(map[string]Room)
var links [][2]string

