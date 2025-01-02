package models

var Entities = []any{
	&Post{},
}

type Post struct {
	Id uint `gorm:"primaryKey"`
	Text string
}

