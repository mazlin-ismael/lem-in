package vizualizer

type Room struct {
	Name 	string
	X		float64
	Y		float64
}

type DataViews struct {
	Rooms	map[string]Room
	Links	[][2]string
}

var rooms map[string]Room = make(map[string]Room)
var links [][2]string

