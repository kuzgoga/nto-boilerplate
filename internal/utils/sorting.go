package utils

import (
	"app/internal/database"
	"errors"
	"fmt"
	"github.com/kuzgoga/fogg"
	"gorm.io/gorm/clause"
	"reflect"
	"strings"
)

// SortByOrder Order items by specified field and a sort type
// Example: SortByOrder(map[string]string{"name": "ASC"}, &models.Post{})
// ASC - по возрастанию (или от А до Я)
// DESC - по убыванию (или от Я до А)
func SortByOrder[T any](fieldsSortOrder map[string]string, entity T) ([]*T, error) {
	var (
		orderQuery []string
		items      []*T
		joins      []string
	)

	for name, order := range fieldsSortOrder {
		structInfo := reflect.ValueOf(entity).Type()
		field, fieldExist := structInfo.FieldByName(name)

		if !fieldExist {
			return nil, errors.New(fmt.Sprintf("Field `%s` not found", name))
		}

		if strings.ToUpper(order) != "ASC" && strings.ToUpper(order) != "DESC" {
			return nil, errors.New(fmt.Sprintf("Field `%s` can only be sorted by ASC or DESC", name))
		}

		tag, err := fogg.Parse(string(field.Tag))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Failed to parse tag for `%s` failed: %s", name, err))
		}

		if !tag.HasTag("ui") {
			return nil, errors.New(fmt.Sprintf("Field `%s` doesn't have ui tag", name))
		}

		if field.Type.Kind() == reflect.Slice {
			return nil, errors.New(fmt.Sprintf("Field `%s` is array and cannot be used for sorting", name))
		}

		fieldPath := tag.GetTag("ui").GetParamOr("field", "")
		if fieldPath == "" {
			orderQuery = append(orderQuery, fmt.Sprintf("%s %s", name, order))
		} else {
			fieldsPathParts := strings.Split(fieldPath, ".")

			if len(fieldsPathParts) > 1 {
				return nil, errors.New(fmt.Sprintf("Too complex fieldPath for structure `%s`", name))
			}

			if len(fieldsPathParts) == 0 {
				return nil, errors.New(fmt.Sprintf("Invalid field path for `%s` field", name))
			}

			joinPathParts := append([]string{field.Type.Name()}, fieldsPathParts...)
			joinField := strings.Join(joinPathParts, ".")
			joinTable := field.Type.Name()
			joins = append(joins, joinTable)
			orderQuery = append(orderQuery, fmt.Sprintf("%s %s", joinField, order))
		}
	}

	db := database.GetInstance()

	for _, join := range joins {
		db = db.Joins(join)
	}

	if len(orderQuery) != 0 {
		db = db.Order(strings.Join(orderQuery, ", "))
	}

	result := db.Preload(clause.Associations).Find(&items)

	if result.Error != nil {
		return items, result.Error
	}

	return items, nil
}
