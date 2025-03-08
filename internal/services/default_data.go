package services

import (
	"app/internal/dialogs"
	"fmt"
	"time"
)

func InsertDefaultData() {
	insertAuthors()
	insertPosts()
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
			Id:       1,
			Text:     "Жителям Кузбасса запретили болеть.",
			Deadline: time.Now().Unix(),
			AuthorId: 1,
		},
		{
			Id:       2,
			Deadline: time.Now().Add(time.Hour * 24 * 5).Unix(),
			Text:     "⚡️⚡️⚡️Дома будут летать.",
			AuthorId: 2,
		},
		{
			Id:       3,
			Deadline: time.Now().Add(time.Hour * 24 * 6).Unix(),
			Text:     "В Кузбассе начали строить дома выше, чтобы жители были ближе к богу и солнцу.",
			AuthorId: 3,
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
