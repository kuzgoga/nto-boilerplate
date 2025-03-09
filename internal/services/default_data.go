package services

import (
	"app/internal/dialogs"
	"fmt"
	"time"
)

func InsertDefaultData() {
	insertPostTypes()
	insertAuthors()
	insertPosts()
	insertComments()
}

func InsertDefaultEntityData[T any](service Service[T], entities []T) {
	for _, item := range entities {
		createdItem, err := service.Create(item)
		if err != nil {
			dialogs.ErrorDialog("Ошибка при вставке данных по умолчанию", fmt.Sprintf("Произошла ошибка при вставке значения %#v: %s", createdItem, err))
		}
	}
}

func insertPosts() {
	InsertDefaultEntityData(&PostService{}, []Post{
		{
			Id:         1,
			Text:       "Жителям Кузбасса запретили болеть.",
			Deadline:   time.Now().Unix(),
			AuthorId:   1,
			PostTypeId: 1,
		},
		{
			Id:         2,
			Deadline:   time.Now().Add(time.Hour * 24 * 5).Unix(),
			Text:       "⚡️⚡️⚡️Дома будут летать.",
			AuthorId:   2,
			PostTypeId: 2,
		},
		{
			Id:         3,
			Deadline:   time.Now().Add(time.Hour * 24 * 6).Unix(),
			Text:       "В Кузбассе начали строить дома выше, чтобы жители были ближе к богу и солнцу.",
			AuthorId:   3,
			PostTypeId: 3,
		},
	})
}

func insertAuthors() {
	InsertDefaultEntityData(&AuthorService{}, []Author{
		{
			Id:   1,
			Name: "ИА Кузбасс",
		},
		{
			Id:   2,
			Name: "ASTRA",
		},
		{
			Id:   3,
			Name: "ЧТД",
		},
	})
}

func insertComments() {
	InsertDefaultEntityData(&CommentService{}, []Comment{
		{
			Id:       1,
			Text:     "Это просто замечательно!",
			AuthorId: 1,
			Posts:    []Post{{Id: 1}},
		},
		{
			Id:       2,
			Text:     "Я тоже думаю, что это замечательно!",
			AuthorId: 2,
			Posts:    []Post{{Id: 2}},
		},
		{
			Id:       3,
			Text:     "Я тоже думаю, что это замечательно!",
			AuthorId: 3,
			Posts:    []Post{{Id: 3}},
		},
	})
}

func insertPostTypes() {
	InsertDefaultEntityData(&PostTypeService{}, []PostType{
		{
			Id:   1,
			Name: "Общество",
		},
		{
			Id:   2,
			Name: "Политика",
		},
		{
			Id:   3,
			Name: "Экономика",
		},
	})
}
