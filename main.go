package main

import (
	"net/http"
	"fmt"

)

func createMapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
     
	return func (w http.ResponseWriter, r *http.Request) {
		
		path, exists := pathToUrls[r.URL.Path] //check if the current url path exists in the map 

		if exists { //if it does then redirect user
			http.Redirect(w, r, path, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func createYAMlHandler(YAML []byte, fallback http.Handler) (http.HandlerFunc, error) {

	//parse YAML

	//convert to map
}





func main() {

	mux := http.NewServeMux() //http res and res handling

	pathsToUrls := map[string]string { //set of mapped shortened urls to actual urls
		"/dogbreed" : "https://www.petfinder.com/dogs-and-puppies/breeds/",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := createMapHandler(pathsToUrls, mux)

	yaml := `
			- path: /urlshort
  				url: https://github.com/gophercises/urlshort
			- path: /urlshort-final
  				url: https://github.com/gophercises/urlshort/tree/solution
			`
	yamlHandler, err := createYAMlHandler([]byte(yaml), mapHandler)

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)

}


