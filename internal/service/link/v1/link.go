package link_service_v1

import (
	"url_shortener/internal/models"
	link_repository_v1 "url_shortener/internal/repository/link/v1"
	"url_shortener/internal/service/link/link_service_errors"
)

type LinkService struct {
	Repository link_repository_v1.LinkRepository
}

func NewLinkService(repository link_repository_v1.LinkRepository) *LinkService {
	return &LinkService{Repository: repository}
}

func (service *LinkService) CreateLink(url string) (*models.Link, error) {
	var resLink *models.Link

	for attempt := 1; attempt <= 10; attempt++ {
		linkCandidate := models.NewLink(url)
		createdLink, err := service.Repository.Create(linkCandidate)
		if err != nil {
			continue
		}

		resLink = createdLink
		break
	}

	if resLink == nil {
		return nil, link_service_errors.ErrCreateLink
	}

	return resLink, nil
}
