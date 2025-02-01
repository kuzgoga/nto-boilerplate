package excel

import (
	"app/internal/dialogs"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log/slog"
	"reflect"
	"slices"
	"strings"
	"time"
)

type TableHeaders struct {
	Headers              []string
	IgnoredFieldsIndices []int
}

func isPrimitive(valueType reflect.Type) bool {
	switch valueType.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return true
	default:
		return false
	}
}

func ExportEntityToSpreadsheet[T any](filename, sheetName string, entity T, provider func() ([]*T, error)) error {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			slog.Error("Error while closing excel file: ", err)
		}
	}()

	if _, err := file.NewSheet(sheetName); err != nil {
		return err
	}
	if err := file.DeleteSheet("Sheet1"); err != nil {
		return err
	}

	headers, err := WriteHeaders(sheetName, entity, file)
	if err != nil {
		return err
	}

	items, err := provider()
	if err != nil {
		return err
	}

	// TODO: process composite objects
	// TODO: appearance
	for i, item := range items {
		structValue := reflect.ValueOf(item).Elem()

		for j := 0; j < structValue.NumField(); j++ {
			if slices.Contains(headers.IgnoredFieldsIndices, j) {
				continue
			}
			field := structValue.Field(j)
			if isPrimitive(field.Type()) {
				fieldValue := field.Interface()
				cellName, err := GetCellNameByIndices(j, i+1)
				if err != nil {
					return err
				}

				cellType := structValue.Type().Field(j).Tag.Get(CellTypeTag)

				var cellValue any

				switch cellType {
				case TimestampTag:
					cellValue = time.Unix(field.Int(), 0)
				case DurationTag:
					cellValue = time.Duration(field.Int())
				default:
					cellValue = fieldValue
				}

				slog.Debug("Field %s value: %v, %s\n", cellName, cellValue, cellType)
				err = file.SetCellValue(sheetName, cellName, cellValue)
				if err != nil {
					return err
				}
			}
		}
	}

	if err := WriteData(file, filename); err != nil {
		return err
	}

	return nil
}

func GetHeaderCellNameByIndex(column int) (string, error) {
	colName, err := excelize.ColumnNumberToName(column + 1)
	if err != nil {
		return "", err
	}
	cellName, err := excelize.JoinCellName(colName, 1)
	if err != nil {
		return "", err
	}
	return cellName, nil
}

func GetCellNameByIndices(column int, row int) (string, error) {
	colName, err := excelize.ColumnNumberToName(column + 1)
	if err != nil {
		return "", err
	}
	cellName, err := excelize.JoinCellName(colName, row+1)
	if err != nil {
		return "", err
	}
	return cellName, nil
}

func ExportHeaders(entity any) TableHeaders {
	headers := TableHeaders{}
	v := reflect.TypeOf(entity)
	for i := range v.NumField() {
		field := v.Field(i)
		displayName := field.Tag.Get("displayName")
		if displayName != "" {
			headers.Headers = append(headers.Headers, displayName)
		} else {
			headers.IgnoredFieldsIndices = append(headers.IgnoredFieldsIndices, i)
		}
	}
	return headers
}

func WriteHeaders(sheetName string, entity any, file *excelize.File) (TableHeaders, error) {
	headers := ExportHeaders(entity)
	for i, header := range headers.Headers {
		cellName, err := GetHeaderCellNameByIndex(i)
		if err != nil {
			return headers, err
		}

		err = file.SetCellValue(sheetName, cellName, header)
		if err != nil {
			return headers, err
		}
	}
	err := ApplyStyleHeaders(file, sheetName, headers)
	if err != nil {
		return headers, err
	}
	return headers, nil
}

func GetStyleId(f *excelize.File, style *excelize.Style) (int, error) {
	styleId, err := f.NewStyle(style)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании стиля: %w", err)
	}

	return styleId, nil
}

func LoadHeadersStyle(file *excelize.File) (int, error) {
	headersStyle := excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      true,
			VertAlign: "center",
		},
	}
	return GetStyleId(file, &headersStyle)
}

func ApplyStyleHeaders(file *excelize.File, sheetName string, headers TableHeaders) error {
	styleId, err := LoadHeadersStyle(file)
	if err != nil {
		return err
	}

	cellName, err := GetHeaderCellNameByIndex(len(headers.Headers) - 1)
	if err != nil {
		return err
	}

	err = file.SetCellStyle(sheetName, "A1", cellName, styleId)
	if err != nil {
		return err
	}

	return nil
}

func WriteData(file *excelize.File, filename string) error {
	filepath, err := dialogs.SaveFileDialog("Экспорт данных", filename)

	if !strings.HasSuffix(filepath, ".xlsx") {
		filepath += ".xlsx"
	}

	if err != nil {
		return err
	}
	if err := file.SaveAs(filepath); err != nil {
		return err
	}
	return nil
}
