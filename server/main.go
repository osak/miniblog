package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/osak/miniblog/db"
	"github.com/osak/miniblog/resources"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
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

type PostController struct {
	dbx       sqlx.DB
	postStore db.PostStore
}

type PostModel struct {
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	BodyHtml string `json:"bodyHtml"`
	PostedAt string `json:"postedAt"`
}

func (p *PostController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := p.dbx.Connx(context.Background())
	if err != nil {
		fmt.Printf("Failed to connect to db: %v", err)
		w.WriteHeader(500)
		return
	}

	posts, err := p.postStore.FindAll(conn)
	if err != nil {
		fmt.Printf("Failed to fetch posts: %v", err)
		w.WriteHeader(500)
		return
	}

	postModels := make([]PostModel, len(posts))
	for i, post := range posts {
		postModels[i] = PostModel{
			Id:       post.Id.String(),
			Slug:     post.Slug,
			Title:    post.Title,
			BodyHtml: post.Body,
			PostedAt: time.UnixMilli(post.PostedAt).Format(time.RFC3339),
		}
	}

	response := map[string]interface{}{
		"posts": postModels,
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		fmt.Printf("Failed to encode response: %v", err)
		w.WriteHeader(500)
		return
	}
}

func main() {
	s := server{}
	dbx, err := resources.OpenDB()
	if err != nil {
		panic(err)
	}
	postController := PostController{
		dbx:       *dbx,
		postStore: &db.PostStoreImpl{},
	}

	http.HandleFunc("/", s.htmlHandler("index"))
	http.HandleFunc("/post", s.htmlHandler("post"))
	http.Handle("/posts/{slug}", &postController)
	http.Handle("/api/posts", &postController)
	http.HandleFunc("/static/js/post.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../ui/dist/post.js")
	})
	http.HandleFunc("/static/css/main.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../assets/css/main.css")
	})

	if err := http.ListenAndServe("localhost:8081", nil); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
