package excel

import (
	"app/internal/services"
	"github.com/xuri/excelize/v2"
	"log/slog"
	"reflect"
)

type TableHeaders struct {
	Headers              []string
	IgnoredFieldsIndices []uint
}

func ExportEntityToSpreadsheet[T any](filename, sheetName string, entity T, service services.Service[T]) error {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			slog.Error("Error while closing excel file: ", err)
		}
	}()

	return nil
}

func ExportHeaders(entity any) TableHeaders {
	headers := TableHeaders{}
	v := reflect.TypeOf(entity)
	for i := range v.NumField() {
		field := v.Field(i)
		displayName := field.Tag.Get("displayName")
		if displayName != "" {
			
		}
	}
	return headers
}
