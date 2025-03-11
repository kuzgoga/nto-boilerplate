package models

var Entities = []any{
	&Post{}, &Author{}, &PostType{}, &Comment{},
}

type PostType struct {
	Id   uint   `gorm:"primaryKey" ui:"hidden"`
	Name string `ui:"label:Название;"`
}

type Post struct {
	Id         uint      `gorm:"primaryKey" ui:"hidden;label:\"Номер поста\""`
	Text       string    `ui:"label:Текст"`
	Deadline   int64     `ui:"label:Дедлайн;datatype:datetime;"`
	CreatedAt  int64     `gorm:"autoCreateTime" ui:"label:Время создания;readonly;datatype:datetime;"`
	AuthorId   uint      `ui:"hidden" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Author     Author    `ui:"label:Автор; field:Name;"`
	PostTypeId uint      `ui:"hidden; excel:Номер типа поста;"`
	PostType   PostType  `ui:"label:Тип поста; field:Name;" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments   []Comment `ui:"label:Комментарии; field:Text;" gorm:"many2many:comments_post;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Author struct {
	Id       uint      `gorm:"primaryKey" ui:"hidden"`
	Name     string    `ui:"label:Имя;"`
	Posts    []Post    `ui:"label:Посты; field:Text;" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments []Comment `ui:"label:Комментарии; field:Text;" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Comment struct {
	Id       uint   `gorm:"primaryKey" ui:"hidden"`
	Text     string `ui:"label:Текст;"`
	AuthorId uint   `ui:"hidden"`
	Author   Author `ui:"label:Автор; field:Name;" gorm:"foreignKey:AuthorId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Posts    []Post `ui:"label:Посты; field:Text;" gorm:"many2many:comments_post;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
