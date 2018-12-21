package models


type Notes struct{
	Id int `json:"id"`
	Name string	`json:"name"`
	Details string	`json:"details"`
	CreatedAt string	`json:"created_at"`
	UpdatedAt string	`json:"updated_at"`
}
