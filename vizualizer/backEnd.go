package vizualizer

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func hostVizualiser(w http.ResponseWriter, r *http.Request) {
	var dataView DataViews = DataViews{
		Rooms: rooms,
		Links: links,
	}
	
	tpl, errTpl := template.ParseFiles("vizualizer/templates/index.html")
	if errTpl != nil {
		fmt.Println("errTpl", errTpl)
		os.Exit(1)
	}
	tpl.Execute(w, dataView)
}