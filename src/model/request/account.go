package request

type Account struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Mobile string `json:"mobile"`
}
