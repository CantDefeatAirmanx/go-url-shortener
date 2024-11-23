package links_handler_v1

import (
	"errors"
	"net/http"
	link_repository_errors "url_shortener/internal/repository/link"
	link_repository_v1 "url_shortener/internal/repository/link/v1"
	link_service_v1 "url_shortener/internal/service/link/v1"
	"url_shortener/pkg/utils/http_helpers"
)

type LinksHandlerV1 struct {
	Service    *link_service_v1.LinkService
	Repository link_repository_v1.LinkRepository
}

type LinksHandlerV1Deps struct {
	Service    *link_service_v1.LinkService
	Repository link_repository_v1.LinkRepository
	Router     *http.ServeMux
}

const prefix = "/api/links/v1"

func NewLinksHandlerV1(deps LinksHandlerV1Deps) *LinksHandlerV1 {
	handler := LinksHandlerV1{Service: deps.Service, Repository: deps.Repository}

	deps.Router.HandleFunc("POST "+prefix, handler.CreateLink())
	deps.Router.HandleFunc("GET "+prefix+"/{"+goToParams.Hash+"}", handler.GoTo())
	deps.Router.HandleFunc("PATCH "+prefix+"/{"+updateLinksParams.Id+"}", handler.UpdateLink())
	deps.Router.HandleFunc("DELETE "+prefix+"/{"+deleteLinksParams.Id+"}", handler.DeleteLink())

	return &handler
}

func (handler *LinksHandlerV1) CreateLink() http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		reqPayload, err :=
			http_helpers.
				HandleJsonReqBody[LinkCreateRequestBody](&resWriter, req)
		if err != nil {
			return
		}

		link, err := handler.Service.CreateLink(reqPayload.Url)

		if err != nil {
			_ = http_helpers.WriteJsonRes(resWriter, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http_helpers.WriteJsonRes(resWriter, *link, http.StatusCreated)
	}
}

var goToParams = struct{ Hash string }{Hash: "hash"}

func (handler *LinksHandlerV1) GoTo() http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		hash := req.PathValue(goToParams.Hash)

		link, err := handler.Repository.Get(hash)

		if err == nil {
			http.Redirect(resWriter, req, link.Url, http.StatusTemporaryRedirect)
			return
		}

		if errors.Is(err, link_repository_errors.ErrRecordNotFound) {
			http_helpers.WriteJsonRes(resWriter, "Not found Error", http.StatusNotFound)
			return
		}

		http_helpers.WriteJsonRes(resWriter, "Internal Server Errro", http.StatusInternalServerError)
	}
}

var updateLinksParams = struct{ Id string }{Id: "id"}

func (handler *LinksHandlerV1) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

var deleteLinksParams = struct{ Id string }{Id: "id"}

func (handler *LinksHandlerV1) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
