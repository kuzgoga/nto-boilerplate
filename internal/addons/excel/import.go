package excel

import (
	"fmt"
	"log/slog"

	"github.com/xuri/excelize/v2"
)

type Importer struct {
	SheetName       string
	Loader          func(rowIndex int, row []string) error
	StartFromRowIdx int
}

func ImportEntitiesFromSpreadsheet(path string, importers ...Importer) error {
	for _, importer := range importers {
		err := ImportEntitiesFromSpreadsheet(path, importer)
		if err != nil {
			return err
		}
	}
	return nil
}

func ImportFromSpreadsheet(filepath string, importer Importer) error {
	f, err := excelize.OpenFile(filepath)
	defer func() {
		err := f.Close()
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to close file: %s", err))
		}
	}()

	if err != nil {
		return err
	}

	rows, err := f.GetRows(importer.SheetName)

	if err != nil {
		return err
	}

	for i, row := range rows {
		err := importer.Loader(i, row)
		if err != nil {
			return err
		}
	}

	return nil
}
