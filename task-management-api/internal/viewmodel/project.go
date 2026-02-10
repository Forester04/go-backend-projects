package viewmodel

type CreateProjectRequest struct {
	Body struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	} `json:"body" binding:"required"`
}

type CreateProjectResponse struct {
	Body struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"body"`
}

type GetProjectRequest struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

type GetProjectResponse struct {
	Body struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"body"`
}

type DeleteProjectRequest struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

type DeleteProjectResponse struct {
	Body struct {
		ID uint `json:"id"`
	} `json:"body"`
}
