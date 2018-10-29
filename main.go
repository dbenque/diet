package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

func init() {
	go initializeTemplates()
	defineRoutes()
}

var templates = make(map[string]*template.Template)

// Base template is 'theme.html'  Can add any variety of content fillers in /layouts directory
func initializeTemplates() {
	t := time.NewTicker(time.Second)
	for range t.C {
		layouts, err := filepath.Glob("templates/*.html")
		if err != nil {
			log.Fatal(err)
		}

		for _, layout := range layouts {
			templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout, "templates/layouts/theme.html", "templates/layouts/navbar.html", "templates/layouts/sidenav.html"))
		}
	}
}

func defineRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/patient", patientHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		Title   string
		Message string
	}{
		Title:   "WelcomePage",
		Message: "Mon Message",
	}

	templates["welcome.html"].ExecuteTemplate(w, "outerTheme", page)
}

func patientHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		Title   string
		Message string
	}{
		Title:   "WelcomePage",
		Message: "Mon Message",
	}

	templates["patient.html"].ExecuteTemplate(w, "outerTheme", page)
}

var BUILDTIME string
var VERSION string
var COMMIT string
var BRANCH string

func mainHeader() {
	fmt.Println("Program go-ui started at: " + time.Now().String())
	fmt.Println("BUILDTIME=" + BUILDTIME)
	fmt.Println("VERSION=" + VERSION)
	fmt.Println("COMMIT=" + COMMIT)
	fmt.Println("BRANCH=" + BRANCH)
	fmt.Println("-------")
}

func main() {
	mainHeader()
	flag.Parse()
	log.Fatal(http.ListenAndServe(*addr, nil))
}
