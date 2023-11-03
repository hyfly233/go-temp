package model

type (
	Test struct {
		Id string `json:"id" binding:"required,min=1,max=10"`
	}
)
