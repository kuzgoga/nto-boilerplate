package services

import (
	"fmt"
	"github.com/kuzgoga/nto-boilerplate/internal/dialogs"
	"github.com/kuzgoga/nto-boilerplate/internal/models"
)

func InsertDefaultData() {
	insertPosts()
}

// Example of usage
//func insertProductTypes() {
//	InsertDefaultEntityData(&ProductTypeService{}, []ProductType{
//		{Id: 1, Name: "Сырые пиломатериалы"},
//		{Id: 2, Name: "Сухие пиломатериалы"},
//		{Id: 3, Name: "Строганные доски"},
//		{Id: 4, Name: "Рейки"},
//		{Id: 5, Name: "Брус"},
//		{Id: 6, Name: "Пеллеты"},
//	})
//}

func insertPosts() {
	InsertDefaultEntityData(&PostService{}, []models.Post{
		{Text: "В Кузбассе начали строить дома выше, чтобы их жители стали ближе к Богу."},
	})
}

func InsertDefaultEntityData[T any](service Service[T], entities []T) {
	for _, item := range entities {
		createdItem, err := service.Create(item)
		if err != nil {
			dialogs.ErrorDialog("Ошибка при вставке данных по умолчанию", fmt.Sprintf("Произошла ошибка при вставке значения %#v: %s", createdItem, err))
		}
	}
}
