package link_repository_v1

import (
	"errors"
	"url_shortener/internal/models"
	link_repository_errors "url_shortener/internal/repository/link"
	"url_shortener/pkg/gorm_db"

	"gorm.io/gorm"
)

type LinkRepositoryGorm struct {
	Db *gorm_db.Db
}

var singleTon *LinkRepositoryGorm

// follows singleTon pattern
func NewLinkRepositoryGorm(db *gorm_db.Db) *LinkRepositoryGorm {
	if singleTon != nil {
		return singleTon
	}

	singleTon := &LinkRepositoryGorm{Db: db}

	return singleTon
}

func (repository *LinkRepositoryGorm) Create(link *models.Link) (*models.Link, error) {
	res := repository.Db.Create(link)

	if res.Error != nil {
		return nil, res.Error
	}

	return link, nil
}

func (repository *LinkRepositoryGorm) Get(hash string) (*models.Link, error) {
	var link models.Link

	res := repository.Db.First(&link, "hash = ?", hash)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, link_repository_errors.ErrRecordNotFound
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &link, nil
}
