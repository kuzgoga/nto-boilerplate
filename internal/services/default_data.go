package services

import (
	"app/internal/dialogs"
	"fmt"
)

func InsertDefaultData() {
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
			Id:   1,
			Text: "Жителям Кузбасса запретили болеть.",
		},
		{
			Id:   2,
			Text: "⚡️⚡️⚡️Дома будут летать.",
		},
		{
			Id:   3,
			Text: "В Кузбассе начали строить дома выше, чтобы жители были ближе к богу и солнцу.",
		},
	})
}
