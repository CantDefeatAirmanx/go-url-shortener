package links_handler_v1

import (
	"net/http"
	link_service_v1 "url_shortener/internal/service/link/v1"
	"url_shortener/pkg/utils/http_helpers"
)

type LinksHandlerV1 struct {
	Service *link_service_v1.LinkService
}

type LinksHandlerV1Deps struct {
	Service *link_service_v1.LinkService
	Router  *http.ServeMux
}

const prefix = "/api/links/v1"

func NewLinksHandlerV1(deps LinksHandlerV1Deps) *LinksHandlerV1 {
	handler := LinksHandlerV1{Service: deps.Service}

	deps.Router.HandleFunc("POST "+prefix, handler.CreateLink())
	deps.Router.HandleFunc("GET "+prefix+"/{"+getLinksParams.Id+"}", handler.GetLinks())
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

		http_helpers.WriteJsonRes(resWriter, *link, http.StatusOK)
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
