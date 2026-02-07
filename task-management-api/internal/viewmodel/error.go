package viewmodel

type BadRequestErrorResponse struct {
	Body struct {
		Message string            `json:"message"`
		Context map[string]string `json:"context"`
	}
}

type InternalServerErrorResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}
