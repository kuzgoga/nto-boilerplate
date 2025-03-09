package models

var Entities = []any{
	&Post{}, &Author{}, &PostType{},
}

type PostType struct {
	Id   uint   `gorm:"primaryKey" ui:"hidden"`
	Name string `ui:"label:Название;"`
}

type Post struct {
	Id        uint     `gorm:"primaryKey" ui:"hidden"`
	Text      string   `displayName:"Текст" ui:"label:Текст"`
	Deadline  int64    `ui:"label:Дедлайн;datatype:datetime;"`
	CreatedAt int64    `gorm:"autoCreateTime" ui:"readonly;datatype:datetime;"`
	AuthorId  uint     `ui:"hidden" gorm:"foreignKey:Id;references:AuthorId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Author    Author   `ui:"label:Автор; field:Name;"`
	TypeId    uint     `ui:"hidden"`
	PostType  PostType `ui:"label:Тип; field:Name;" gorm:"foreignKey:Id;references:TypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Author struct {
	Id    uint   `gorm:"primaryKey" ui:"hidden"`
	Name  string `ui:"label:Имя;"`
	Posts []Post `ui:"label:Посты; field:Text;" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
