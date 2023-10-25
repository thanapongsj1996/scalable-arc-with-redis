package model

type Member struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsActive int    `json:"is_active"`
}
