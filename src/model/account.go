package model

type Account struct {
	ID     string `form:"user" json:"id" binding:"required"`
	Name   string `form:"password" json:"name" binding:"required"`
	Mobile string `json:"mobile"`
}
