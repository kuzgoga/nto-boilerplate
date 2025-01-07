package models

var Entities = []any{
	&Post{}, &Author{},
}

type Post struct {
	Id        uint   `gorm:"primaryKey"`
	Text      string `displayName:"Текст"`
	CreatedAt int64  `gorm:"autoCreateTime" displayName:"Дата публикации" cellType:"timestamp"`
}

// Author A sample of comment
type Author struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
