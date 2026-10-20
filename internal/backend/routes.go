package me

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func renderPage(
	w http.ResponseWriter,
	r *http.Request,
	page string,
	context any,
) {
	files := []string{
		filepath.Join("templates", "page.html"),
		filepath.Join("templates", "snippets", "ribbon.html"),
		filepath.Join("templates", "snippets", "utils.html"),
		filepath.Join("templates", "pages", page+".html"),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Could not load page", http.StatusInternalServerError)
		return
	}
	if r.Header.Get("HX-Request") == "true" {
		err = tmpl.ExecuteTemplate(w, "content", context)
	} else {
		err = tmpl.ExecuteTemplate(w, "page.html", context)
	}
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Could not render page", http.StatusInternalServerError)
	}
}

func renderProject(
	w http.ResponseWriter,
	r *http.Request,
	project *Project,
) {
	files := []string{
		filepath.Join("templates", "page.html"),
		filepath.Join("templates", "pages", "project.html"),
		filepath.Join("templates", "snippets", "ribbon.html"),
		filepath.Join("templates", "snippets", "utils.html"),
		filepath.Join("templates", "snippets", "content_blocks.html"),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templateName := "page.html"
	if r.Header.Get("HX-Request") == "true" {
		templateName = "content"
	}
	err = tmpl.ExecuteTemplate(w, templateName, project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "home", nil)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "projects", ProjectDB.Projects)
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	endPoint := strings.TrimPrefix(r.URL.Path, "/projects/")
	log.Println("Target endpoint:", endPoint)
	for _, project := range ProjectDB.Projects {
		if len(project.EndPoint) == 0 {
			project.EndPoint = GenerateEndPoint(project.Title)
			log.Println("Generated endpoint:", project.EndPoint)
		}
		if project.EndPoint == endPoint {
			log.Println("Found endpoint")
			renderProject(w, r, project)
			return
		}
	}
	http.NotFound(w, r)
}

func activitiesHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "activities", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "about", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "contact", nil)
}

func CreateHandlers() {
	// Create content servers
	staticFiles := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	projectFiles := http.FileServer(http.Dir("./data"))
	http.Handle("/data/", http.StripPrefix("/data/", projectFiles))
	// Register routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/projects/", projectHandler)
	http.HandleFunc("/activities", activitiesHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
}
