package links_v1

import (
	"fmt"
	"net/http"
)

type LinksHandlerV1 struct{}

const prefix = "/api/links/v1"

func NewLinksHandlerV1(router *http.ServeMux) *LinksHandlerV1 {
	handler := LinksHandlerV1{}

	router.HandleFunc(fmt.Sprintf("POST %s", prefix), handler.CreateLink())
	router.HandleFunc(fmt.Sprintf("GET %s/{alias}", prefix), handler.GetLinks())
	router.HandleFunc(fmt.Sprintf("PATCH %s/{id}", prefix), handler.UpdateLink())
	router.HandleFunc(fmt.Sprintf("DELETE %s/{id}", prefix), handler.DeleteLinks())

	return &handler
}

func (hander *LinksHandlerV1) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinksHandlerV1) GetLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinksHandlerV1) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinksHandlerV1) DeleteLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
