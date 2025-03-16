package models

var Entities = []any{&Post{}}

type Post struct {
	Id        uint   `gorm:"primaryKey" ui:"hidden"`
	Text      string `ui:"label:Текст;"`
	CreatedAt uint   `gorm:"autoCreateTime" ui:"readonly"`
}
