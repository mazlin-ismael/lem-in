package vizualizer

import (
	farmer "lem-in/errFile"
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
	
	http.Handle("/visualizer/static/", http.StripPrefix("/visualizer/static/", http.FileServer(http.Dir("visualizer/static"))))
	visualizer2DLaunch()
	visualizer3DLaunch()
	http.ListenAndServe(port, nil)
}

func visualizer2DLaunch() {
	http.HandleFunc("/2d", hostVizualiser2D)
	fmt.Println("\nvisualizer 2d launch on: http://localhost" + port + "/2d")
}

func visualizer3DLaunch() {
	http.HandleFunc("/3d", hostVizualiser3D)
	fmt.Println("visualizer 3d launch on: http://localhost" + port + "/3d")
}