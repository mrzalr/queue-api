package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

type ResponseWithMeta struct {
	Response
	Meta Meta `json:"meta"`
}

type Meta struct {
	TotalPages      int  `json:"total_pages"`
	CurrentPages    int  `json:"current_pages"`
	TotalItems      int  `json:"total_items"`
	HasNextPage     bool `json:"has_next_page"`
	HasPreviousPage bool `json:"has_previous_page"`
}
