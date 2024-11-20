package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

type server struct {
	templates map[string]*template.Template
}

func loadTemplate(name string) (*template.Template, error) {
	f, err := os.Open("templates/" + name + ".template.html")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(buf))
}

func (s *server) htmlHandler(name string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		layoutTemplate, err := loadTemplate("_layout")
		if err != nil {
			w.WriteHeader(500)
			return
		}
		contentTemplate, err := loadTemplate(name)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		contentBuf := bytes.Buffer{}
		if err := contentTemplate.Execute(&contentBuf, nil); err != nil {
			w.WriteHeader(500)
			return
		}

		params := map[string]interface{}{
			"Content": template.HTML(contentBuf.String()),
		}
		w.WriteHeader(200)
		if err := layoutTemplate.Execute(w, params); err != nil {
			return
		}
	}
}

func main() {
	s := server{}

	http.HandleFunc("/", s.htmlHandler("index"))
	http.HandleFunc("/post", s.htmlHandler("post"))
	http.HandleFunc("/static/js/post.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../ui/dist/post.js")
	})
	http.HandleFunc("/static/css/main.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/css/main.css")
	})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
