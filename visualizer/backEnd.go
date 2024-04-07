package vizualizer

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func hostVizualiser2D(w http.ResponseWriter, r *http.Request) {
	templateSet("index2d.html", w, initDataView())
}

func hostVizualiser3D(w http.ResponseWriter, r *http.Request) {
	templateSet("index3d.html", w, initDataView())
}


func templateSet(file string, w http.ResponseWriter, dataView DataViews) {
	tpl, errTpl := template.ParseFiles("visualizer/templates/" + file)
	if errTpl != nil {
		fmt.Println("errTpl", errTpl)
		os.Exit(1)
	}
	tpl.Execute(w, dataView)
}