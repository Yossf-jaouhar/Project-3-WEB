package funcs

import (
	"log"             
	"net/http"      
	"text/template"   
)

// Page struct holds data for rendering the HTML template
type Page struct {
	Text     string
	Banner   string 
	AsciiArt string 
}

// Global variable to hold the page data
var Pagedata = &Page{}

// ArtHandler handles incoming requests to generate ASCII art
func ArtHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Please use POST.", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data.", http.StatusBadRequest)
		return
	}

	// Retrieve the user input from the form
	UserText := r.FormValue("usertext")
	if UserText == "" {
		http.Error(w, "No text provided in the form.", http.StatusBadRequest)
		return
	}

	// Retrieve the selected banner style from the form
	Banner := r.FormValue("Banner")
	if Banner == "" {
		http.Error(w, "No banner style selected.", http.StatusBadRequest)
		return
	}

	// Generate ASCII art based on the selected banner style
	result, err := ReadB(Banner)
	if err != nil {
		http.Error(w, "Error reading banner.", http.StatusInternalServerError)
		log.Printf("Error reading banner: %v", err) // Log the error for debugging
		return
	}

	// Process the user text and generated result to create ASCII art
	asciiart := TreatData(result, UserText)

	// Populate the page data with user input and generated ASCII art
	Pagedata.Text = UserText
	Pagedata.Banner = Banner
	Pagedata.AsciiArt = asciiart

	// Load the HTML template from the specified file
	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template.", http.StatusInternalServerError)
		log.Printf("Template loading error: %v", err) 
		return
	}

	// Execute the template and pass the populated page data to render it
	err = tmpl.Execute(w, Pagedata)
	if err != nil {
		http.Error(w, "Error rendering template.", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err) 
		return
	}
}
