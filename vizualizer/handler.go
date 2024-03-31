package vizualizer

import (
	farmer "lem-in/Handler"
	"net/http"
	"fmt"
)

var port = ":2030"

func WebHandler(farm farmer.FarmProperties, optimalComb [][]string, antsByPaths []int) {
	multiX, multiY := multiplicatorsInit(farm.Rooms)
	initNewsRooms(farm.Rooms, multiX, multiY)
	initLinks(farm)
	initEndpoints(farm)
	initComb(optimalComb)
	initAntsByPaths(antsByPaths)
	
	http.Handle("/vizualizer/static/", http.StripPrefix("/vizualizer/static/", http.FileServer(http.Dir("vizualizer/static"))))
	http.HandleFunc("/", hostVizualiser)
	fmt.Println("\nvisualizer launch on: http://localhost" + port)
	http.ListenAndServe(port, nil)
}