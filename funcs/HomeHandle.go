package funcs

import (
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Please use GET.", http.StatusMethodNotAllowed)
		return
	}

	// Check if the URL path is not the root ("/")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}


	
	Pagedata.Text = ""
	Pagedata.Banner = "Banner"
	Pagedata.AsciiArt = ""
	// Execute the template without any data
	err = tmpl.Execute(w, Pagedata)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
}
