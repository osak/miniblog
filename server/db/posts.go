package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type Post struct {
	Id        uuid.UUID
	Slug      string
	Title     string
	Body      string
	PostedAt  int64     `db:"posted_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PostStore interface {
	FindAll(conn *sqlx.Conn) ([]Post, error)
	FindById(conn *sqlx.Conn, id uuid.UUID) (*Post, error)
	FindBySlug(conn *sqlx.Conn, slug string) (*Post, error)
}

type PostStoreImpl struct{}

func (p *PostStoreImpl) FindAll(conn *sqlx.Conn) ([]Post, error) {
	var posts []Post
	err := conn.SelectContext(p.getDefaultContext(), &posts, "SELECT * FROM posts ORDER BY posted_at DESC")
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostStoreImpl) FindById(conn *sqlx.Conn, id uuid.UUID) (*Post, error) {
	var post Post
	err := conn.GetContext(p.getDefaultContext(), &post, "SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *PostStoreImpl) FindBySlug(conn *sqlx.Conn, slug string) (*Post, error) {
	var post Post
	err := conn.GetContext(p.getDefaultContext(), &post, "SELECT * FROM posts WHERE slug = ?", slug)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *PostStoreImpl) getDefaultContext() context.Context {
	return context.Background()
}
