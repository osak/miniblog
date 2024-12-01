package main

import (
	"bytes"
	"context"
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
	Id       string
	Slug     string
	Title    string
	BodyHtml template.HTML
	PostedAt string
}

func (p *PostController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob("templates/*.html"))
	slug := r.PathValue("slug")
	if slug == "" {
		conn, err := p.dbx.Connx(context.Background())
		if err != nil {
			println("Failed to connect to db")
			w.WriteHeader(500)
			return
		}
		posts, err := p.postStore.FindAll(conn)
		if err != nil {
			fmt.Printf("Failed to fetch posts: %v", err)
			w.WriteHeader(500)
			return
		}
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Printf("Failed to load timezone: %v", err)
			w.WriteHeader(500)
			return
		}
		models := make([]PostModel, 0, len(posts))
		for _, post := range posts {
			models = append(models, PostModel{
				Id:       post.Id.String(),
				Slug:     post.Slug,
				Title:    post.Title,
				BodyHtml: template.HTML(post.Body),
				PostedAt: time.UnixMilli(post.PostedAt).In(jst).Format("2006-01-02 15:04:05 -0700"),
			})
		}
		params := map[string]interface{}{
			"Posts": models,
		}

		contentBuf := bytes.Buffer{}
		err = templates.ExecuteTemplate(&contentBuf, "posts.template.html", params)
		if err != nil {
			fmt.Printf("Failed to render template: %v", err)
			w.WriteHeader(500)
			return
		}

		err = templates.ExecuteTemplate(w, "_layout.template.html", map[string]interface{}{
			"Content": template.HTML(contentBuf.String()),
		})
		if err != nil {
			fmt.Printf("Failed to render template: %v", err)
			return
		}
	} else {
		w.WriteHeader(404)
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
	http.Handle("/posts", &postController)
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
