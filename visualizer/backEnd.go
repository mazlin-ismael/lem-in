package vizualizer

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Host the 2D visualizer
func hostVizualiser2D(w http.ResponseWriter, r *http.Request) {
	templateSet("index2d.html", w, initDataView())
}

// Host the 3D visualizer
func hostVizualiser3D(w http.ResponseWriter, r *http.Request) {
	templateSet("index3d.html", w, initDataView())
}

// Set the template and execute it with dataView
func templateSet(file string, w http.ResponseWriter, dataView DataViews) {
	tpl, errTpl := template.ParseFiles("visualizer/templates/" + file)
	if errTpl != nil {
		fmt.Println("errTpl", errTpl)
		os.Exit(1)
	}
	tpl.Execute(w, dataView)
}