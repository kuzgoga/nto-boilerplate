package utils

import (
	"app/internal/database"
	"errors"
	"fmt"
	"reflect"
)

// SortByOrder Order items by specified field and a sort type
// Example: SortByOrder(map[string]string{"name": "ASC"}, &models.Post{})
// ASC - по возрастанию (от А до Я)
// DESC - по убыванию (от Я до А)
func SortByOrder[T any](fieldsSortOrder map[string]string, entity T) ([]*T, error) {
	var (
		orderQuery string
		items      []*T
	)

	for name, order := range fieldsSortOrder {
		structInfo := reflect.ValueOf(entity).Type()
		_, fieldExist := structInfo.FieldByName(name)

		if !fieldExist {
			return nil, errors.New(fmt.Sprintf("Field `%s` not found", name))
		}

		if order != "ASC" && order != "DESC" {
			return nil, errors.New(fmt.Sprintf("Field `%s` can only be sorted by ASC or DESC", name))
		}

		orderQuery += fmt.Sprintf("%s %s", name, order)
	}

	result := database.GetInstance().Order(orderQuery).Find(&items)
	if result.Error != nil {
		return items, result.Error
	}
	return items, nil
}
