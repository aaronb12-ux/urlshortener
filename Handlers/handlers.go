package handlers

import (
	"log"
	"net/http"
	"gopkg.in/yaml.v3"
)

    
type PathURL struct { //struct we unmarshal the YAML into
    Path string `yaml:"path"`
    URL  string `yaml:"url"`
}

func ParseYAML(YAML []byte) []PathURL { //unmarshal the YAML into the 'PathURL' struct. This creates a slice of 'PathURL' objects

	var parsed []PathURL

	if err := yaml.Unmarshal(YAML, &parsed); err != nil { 
		log.Fatal(err)
	}

	return parsed
}


func YamlToMap(yaml []PathURL) map[string]string { //convert the parsed yaml into a map for easier key value access

	m := make(map[string]string)

	for _, element := range yaml {
		m[element.Path] = element.URL
	}

	return m
}


func CreateMapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc { 
     
	return func (w http.ResponseWriter, r *http.Request) {
		
		path, exists := pathToUrls[r.URL.Path] //check if the current url path exists in the map 

		if exists { //if it does then redirect user
			http.Redirect(w, r, path, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r) 
		}
	}
}


func CreateYAMlHandler(YAML []byte, fallback http.Handler) (http.HandlerFunc) {
	 
	parsedYAML := ParseYAML(YAML)
	
	yamlToMap := YamlToMap(parsedYAML)

	return func (w http.ResponseWriter, r *http.Request) {
		
		path, exists := yamlToMap[r.URL.Path]

		if exists {
			http.Redirect(w, r, path, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}


func HomeHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The path you entered does not map to a URL"))
}