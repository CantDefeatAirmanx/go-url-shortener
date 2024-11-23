package links_handler_v1

import (
	"net/http"
	link_service_v1 "url_shortener/internal/service/link/v1"
)

type LinksHandlerV1 struct{}

type LinksHandlerV1Deps struct {
	Service *link_service_v1.LinkService
	Router  *http.ServeMux
}

const prefix = "/api/links/v1"

func NewLinksHandlerV1(deps LinksHandlerV1Deps) *LinksHandlerV1 {
	handler := LinksHandlerV1{}

	deps.Router.HandleFunc("POST "+prefix, handler.CreateLink())
	deps.Router.HandleFunc("GET "+prefix+"/{"+getLinksParams.Id+"}", handler.GetLinks())
	deps.Router.HandleFunc("PATCH "+prefix+"/{"+updateLinksParams.Id+"}", handler.UpdateLink())
	deps.Router.HandleFunc("DELETE "+prefix+"/{"+deleteLinksParams.Id+"}", handler.DeleteLink())

	return &handler
}

func (handler *LinksHandlerV1) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

var getLinksParams = struct{ Id string }{Id: "id"}

func (handler *LinksHandlerV1) GetLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
