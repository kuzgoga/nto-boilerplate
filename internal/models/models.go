package models

var Entities = []any{
	&Post{}, &Author{},
}

type Post struct {
	Id        uint   `gorm:"primaryKey" ui:"hidden"`
	Text      string `displayName:"Текст" ui:"label=Текст"`
	CreatedAt int64  `gorm:"autoCreateTime" ui:"hidden"`
}

type Author struct {
	Id   uint   `gorm:"primaryKey" ui:"hidden"`
	Name string `ui:"label=Имя"`
}
