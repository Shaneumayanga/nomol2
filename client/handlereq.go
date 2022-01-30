package client

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Shaneumayanga/nomol/sites"
)

var home = template.Must(template.ParseFiles("./static/index.html"))

func handleReq(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet && req.RequestURI == "/" {
		sites := sites.GetSites()
		home.Execute(w, map[string]interface{}{
			"Sites": sites,
		})
	} else if req.Method == http.MethodPost && req.RequestURI == "/new" {
		newSite := req.FormValue("site")
		log.Println(newSite)
	}
}
