package link_service_v1

import link_repository_v1 "url_shortener/internal/repository/link/v1"

type LinkService struct {
	repository link_repository_v1.LinkRepository
}

func NewLinkService(repository link_repository_v1.LinkRepository) *LinkService {
	return &LinkService{repository: repository}
}
