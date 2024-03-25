package vizualizer

import (
	"net/http"
	farmer "lem-in/Handler"
)

var port = ":2030"

func WebHandler(farm farmer.FarmProperties, optimalComb [][]string, antsByPaths []int) {
	multiX, multiY := multiplicatorsInit(farm.Rooms)
	initNewsRooms(farm.Rooms, multiX, multiY)
	initLinks(farm)
	initEndpoints(farm)
	
	http.Handle("/vizualizer/static/", http.StripPrefix("/vizualizer/static/", http.FileServer(http.Dir("vizualizer/static"))))
	http.HandleFunc("/", hostVizualiser)
	http.ListenAndServe(port, nil)
}