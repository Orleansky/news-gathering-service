package repo

import (
	"Anastasia/skillfactory/advanced/news-gathering-service/pkg/models"
)

type Interface interface {
	CreatePost(models.Post) error
	Posts(int) ([]models.Post, error)
}
