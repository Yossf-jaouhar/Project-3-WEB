package main

import (
	"fmt"
	"log"
	"net/http"

	"art/funcs"
)

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", funcs.HomeHandler)
	http.HandleFunc("/Art", funcs.ArtHandler)


	fmt.Printf("Server is running at http://localhost%s/\n" , ":8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
