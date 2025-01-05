package models

var Entities = []any{
	&Post{}, &Author{},
}

type Post struct {
	Id   uint `gorm:"primaryKey"`
	Text string
}

type Author struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
