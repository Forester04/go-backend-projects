package viewmodel

type CreateTaskRequest struct {
	Body struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Status      string `json:"status,omitempty"`
	} `json:"body" binding:"required"`
}

type CreateTaskResponse struct {
	Body struct {
		ID    uint   `json:"id"`
		Title string `json:"name"`
	} `json:"body"`
}
type GetTaskRequest struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

type GetTaskResponse struct {
	Body struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
	}
}

type DeleteTaskRequest struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

type DeleteTaskResponse struct {
	Body struct {
		ID uint `json:"id"`
	} `json:"body"`
}
