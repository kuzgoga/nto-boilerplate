package utils

import (
	"app/internal/database"
	"fmt"
	"github.com/kuzgoga/fogg"
	"gorm.io/gorm/clause"
	"reflect"
)

func FindPhraseByAllFields[T any](phrase string, entity T) ([]*T, error) {
	db := database.GetInstance().Preload(clause.Associations)
	structType := reflect.TypeOf(entity)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		tag, err := fogg.Parse(string(field.Tag))
		if err != nil {
			return nil, fmt.Errorf("ошибка при разборе тэга '%s': %w", field.Name, err)
		}

		if field.Type.Kind() == reflect.Pointer {
			field.Type = field.Type.Elem()
		}

		if field.Type.Kind() == reflect.String {
			db.Where(fmt.Sprintf("`%s` like ?", field.Name), "%"+phrase+"%")
		} else {
			if tag.HasTag("ui") {
				uiTag := tag.GetTag("ui")
				nestedFieldPath := uiTag.GetParamOr("field", "")
				if nestedFieldPath != "" {
					db.Preload(fmt.Sprintf("%s.%s", field.Name, nestedFieldPath))
				}
			}
		}

	}
	return nil, nil
}
