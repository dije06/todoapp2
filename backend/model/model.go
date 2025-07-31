package model

type Todo struct {
	ID    int    `json:"ID" gorm:"primaryKey;autoIncrement"`
	Title string `json:title`
}
