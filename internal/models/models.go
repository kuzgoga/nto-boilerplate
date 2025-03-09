package models

var Entities = []any{
	&Post{}, &Author{},
}

type Post struct {
	Id        uint   `gorm:"primaryKey" ui:"hidden"`
	Text      string `displayName:"Текст" ui:"label=Текст"`
	Deadline  int64  `ui:"label=Дедлайн"`
	CreatedAt int64  `gorm:"autoCreateTime" ui:"hidden"`
	AuthorId  uint
	Author    Author `ui:"label=Автор, data=Author, field=[Name]" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
}

type Author struct {
	Id    uint   `gorm:"primaryKey" ui:"hidden"`
	Name  string `ui:"label=Имя"`
	Posts []Post `ui:"label=Посты, data=Post, field=[Text]" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
}

// TODO: correct processing the semicolon (get attention to quotes)
