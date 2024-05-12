package handler

type CreateResourceRequest struct {
	Name        string  `json:"name" binding:"required"`
	Capacity    uint32  `json:"capacity" binding:"required,min=1"`
	Description *string `json:"description"`
}
