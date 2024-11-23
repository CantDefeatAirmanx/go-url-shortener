package link_repository_v1

import "url_shortener/internal/models"

type LinkRepository interface {
	Create() (int, error)
	Get(id int) (*models.Link, error)
}
