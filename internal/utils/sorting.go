package utils

import (
	"app/internal/database"
	"errors"
	"fmt"
	"github.com/kuzgoga/fogg"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gorm.io/gorm/clause"
	"reflect"
	"strings"
)

import "gorm.io/gorm"

func GetTableName(model any, db *gorm.DB) (string, error) {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(model); err != nil {
		return "", err
	}
	return stmt.Schema.Table, nil
}

func GetColumnName(model any, fieldName string, db *gorm.DB) (string, error) {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(model); err != nil {
		return "", err
	}

	field := stmt.Schema.LookUpField(fieldName)
	if field == nil {
		return "", fmt.Errorf("field '%s' not found", fieldName)
	}

	return field.DBName, nil
}

var p = message.NewPrinter(language.Russian)

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
			return nil, errors.New(p.Sprintf("Field %s not found", name))
		}

		if strings.ToUpper(order) != "ASC" && strings.ToUpper(order) != "DESC" {
			return nil, errors.New(p.Sprintf("Field `%s` can only be sorted by ASC or DESC", name))
		}

		tag, err := fogg.Parse(string(field.Tag))
		if err != nil {
			return nil, errors.New(p.Sprintf("Failed to parse tag for `%s` failed: %s", name, err))
		}

		if !tag.HasTag("ui") {
			return nil, errors.New(p.Sprintf("Field `%s` doesn't have ui tag", name))
		}

		if field.Type.Kind() == reflect.Slice {
			return nil, errors.New(p.Sprintf("Field `%s` is array and cannot be used for sorting", name))
		}

		fieldPath := tag.GetTag("ui").GetParamOr("field", "")
		if fieldPath == "" {
			tableName, err := GetTableName(entity, database.GetInstance())
			if err != nil {
				return nil, errors.New(p.Sprintf("Failed to get table name: %s", err))
			}
			columnName, err := GetColumnName(entity, field.Name, database.GetInstance())
			if err != nil {
				return nil, errors.New(p.Sprintf("Failed to get column name: %s", err))
			}
			orderQuery = append(orderQuery, fmt.Sprintf("%s.%s %s", tableName, columnName, order))
		} else {
			fieldsPathParts := strings.Split(fieldPath, ".")

			if len(fieldsPathParts) > 1 {
				return nil, errors.New(p.Sprintf("Too complex fieldPath for structure `%s`", name))
			}

			if len(fieldsPathParts) == 0 {
				return nil, errors.New(p.Sprintf("Invalid field path for `%s` field", name))
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
