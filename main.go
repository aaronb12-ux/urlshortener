package main

import (

	"net/http"
	"fmt"
)

func createMapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return nil
}


func main() {

	mux := http.NewServeMux() //http res and res handling

	pathsToUrls := map[string]string { //set of mapped shortened urls to actual urls
		"/dogbreed" : "https://www.petfinder.com/dogs-and-puppies/breeds/",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
    
	
	


}


