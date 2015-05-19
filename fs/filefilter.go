package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
)

func main() {

	p := "/var/testfiles"

	// Get file names.
	glb := path.Join(p, "*.html")
	log.Printf("DEBUG: glb = %s\n", glb)
	files, err := filepath.Glob(glb)
	if err != nil {
		log.Println("Could not get files.")
	}
	log.Printf("DEBUG: files = %#v\n", files)

	templates := files[:0]
	for _, f := range files {
		switch filepath.Base(f) {
		case "layout.html":
			fmt.Println("Found layout.html file!")
		case "404.html":
			fmt.Println("Found 404.html file!")
		default:
			templates = append(templates, f)
		}
	}

	fmt.Printf("%#v\n", templates)
}
