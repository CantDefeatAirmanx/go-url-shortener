package link_repository_v1

import (
	"url_shortener/internal/models"
	"url_shortener/pkg/gorm_db"
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

func (repository *LinkRepositoryGorm) Create() (int, error) {
	panic("TODO: Implement")
}

func (repository *LinkRepositoryGorm) Get(id int) (*models.Link, error) {
	panic("TODO: Implement")
}
