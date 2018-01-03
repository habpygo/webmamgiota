package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type MaMWebApp struct {
	mampage *controllers.MamSetup
}

//Ld is layout definition
var Ld string

//Td is template definition
var Td string

var Layout string

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	if templateName == "index.html" {
		Ld = filepath.Join("web", "templates", "layout-login.html")
		Layout = "layout-login"
	} else {
		Ld = filepath.Join("web", "templates", "layout.html")
		Layout = "layout"
	}
	Td = filepath.Join("web", "templates", templateName)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(Td)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}
	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	resultTemplate, err := template.ParseFiles(Td, Ld)
	if err != nil {
		// Log the detailed error
		fmt.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err := resultTemplate.ExecuteTemplate(w, Layout, data); err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
