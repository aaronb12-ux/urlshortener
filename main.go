package main

import (
	"fmt"
	"log"
	"net/http"
	"gopkg.in/yaml.v3"
)

    
type PathURL struct {
    Path string `yaml:"path"`
    URL  string `yaml:"url"`
}

func parseYAML(YAML []byte) []PathURL { //unmarshal the YAML into the 'PathURL' struct. This creates a slice of 'PathURL' objects

	var parsed []PathURL

	if err := yaml.Unmarshal(YAML, &parsed); err != nil { 
		log.Fatal(err)
	}

	return parsed

}

func yamlToMap(yaml []PathURL) map[string]string { //convert the parsed yaml into a map for easier key value access

	m := make(map[string]string)

	for _, element := range yaml {
		m[element.URL] = element.Path	
	}

	return m
}

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


func createYAMlHandler(YAML []byte, fallback http.Handler) (http.HandlerFunc) {
	 
	parsedYAML := parseYAML(YAML)
	
	yamlToMap := yamlToMap(parsedYAML)

	return func (w http.ResponseWriter, r *http.Request) {
		path, exists := yamlToMap[r.URL.Path]

		if exists {
			http.Redirect(w, r, path, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
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
	//yamlHandler, err := createYAMlHandler([]byte(yaml), mapHandler)

	yamlHandler := createYAMlHandler([]byte(yaml), mapHandler)
	
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)

}


