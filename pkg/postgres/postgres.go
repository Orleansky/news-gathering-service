package postgres

import (
	"Anastasia/skillfactory/advanced/news-gathering-service/pkg/models"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func New(connstr string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}

	s := Store{
		db: db,
	}

	return &s, nil
}

func (s *Store) CreatePost(p models.Post) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO posts (title, content, published_at, link)
		VALUES ($1, $2, $3, $4)`,
		p.Title, p.Content, p.PublishedAt, p.Link)

	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Posts(n int) ([]models.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
		posts.id,
		posts.title,
		posts.content,
		posts.published_at,
		posts.link
		FROM posts
		ORDER BY DESC
		LIMIT ($1)`,
		n)

	if err != nil {
		return nil, err
	}

	posts := []models.Post{}

	for rows.Next() {
		var p models.Post
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PublishedAt,
			&p.Link,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}
	return posts, rows.Err()
}
