package vizualizer

import (
	farmer "lem-in/errFile"
	"net/http"
	"fmt"
)

var port = ":2030"

// Handle web requests for the visualizer
func WebHandler(farm farmer.FarmProperties, optimalComb [][]string, antsByPaths []int) {
	
	// Initialize multiplicators and rooms
	multiX, multiY := multiplicatorsInit(farm.Rooms)
	initNewsRooms(farm.Rooms, multiX, multiY)

	// Initialize links, endpoints, combinations and ants by paths
	initLinks(farm)
	initEndpoints(farm)
	initComb(optimalComb)
	initAntsByPaths(antsByPaths)
	
	// Serve static files, launch 2D and 3D visualizer and start the HTTP server
	http.Handle("/visualizer/static/", http.StripPrefix("/visualizer/static/", http.FileServer(http.Dir("visualizer/static"))))
	visualizer2DLaunch()
	visualizer3DLaunch()
	http.ListenAndServe(port, nil)
}

// Launch the 2D visualizer
func visualizer2DLaunch() {
	http.HandleFunc("/2d", hostVizualiser2D)
	fmt.Println("\nvisualizer 2d launch on: http://localhost" + port + "/2d")
}

// Launch the 3D visualizer
func visualizer3DLaunch() {
	http.HandleFunc("/3d", hostVizualiser3D)
	fmt.Println("visualizer 3d launch on: http://localhost" + port + "/3d")
}