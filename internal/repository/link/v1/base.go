package link_repository_v1

import "url_shortener/internal/models"

type LinkRepository interface {
	Create(link *models.Link) (*models.Link, error)
	Get(hash string) (*models.Link, error)
}
