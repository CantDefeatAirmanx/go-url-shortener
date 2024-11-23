package links_handler_v1

type LinkCreateRequestBody struct {
	Url string `json:"url" validate:"required,url"`
}
