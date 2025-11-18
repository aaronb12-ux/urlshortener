package main

import (
	"fmt"
	"net/http"
	"aaronb.com/urlshortener/handlers"

)
func homeHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The path you entered does not map to a url"))
}

func main() {

	mux := http.NewServeMux() //http res and res handling
	mux.HandleFunc("/", homeHandler)

	pathsToUrls := map[string]string { //set of mapped shortened urls to actual urls
		"/dogbreed" : "https://www.petfinder.com/dogs-and-puppies/breeds/",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/go":     "https://gophercises.com/",
		
	}

	mapHandler := handlers.CreateMapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
- path: /goooooooo
  url: https://gophercises.com/
`
	yamlHandler := handlers.CreateYAMlHandler([]byte(yaml), mapHandler)
	
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)

}


