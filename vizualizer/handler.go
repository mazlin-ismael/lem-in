package vizualizer

import (
	"net/http"
)

var port = ":2030"

func Handler() {
	http.Handle("/vizualizer/static/", http.StripPrefix("/vizualizer/static/", http.FileServer(http.Dir("vizualizer/static"))))
	http.HandleFunc("/", hostVizualiser)
	http.ListenAndServe(port, nil)
}
