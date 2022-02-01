package models

type Todo struct {
	Base
	Title 			string `json:"title"`
	Description 	string `json:"description"`
}

func (m Todo) TableName() string {
	return "todo"
}