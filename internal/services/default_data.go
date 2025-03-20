package services

import (
	"fmt"
	"github.com/kuzgoga/nto-boilerplate/internal/dialogs"
	"github.com/kuzgoga/nto-boilerplate/internal/models"
	"time"
)

func InsertDefaultData() {
	insertWoodSpecTypes()
	insertWoodSpecs()
	insertExporters()
	insertReceivers()
	InsertWorkResults()
	insertSushkaResults()
	insertDryModes()
}

func insertWoodSpecTypes() {
	InsertDefaultEntityData(&WoodSpecTypeService{}, []models.WoodSpecType{
		{
			Id:   1,
			Name: "Круглый лес",
		},
		{
			Id:   2,
			Name: "Сырые пиломатериалы",
		},
		{
			Id:   3,
			Name: "Сухие пиломатериалы",
		},
		{
			Id:   4,
			Name: "Опилки",
		},
		{
			Id:   5,
			Name: "Пеллеты",
		},
	})
}

func insertWoodSpecs() {
	InsertDefaultEntityData(&WoodSpecService{}, []models.WoodSpec{
		{
			Id:             1,
			WoodSpecTypeId: 1,
			Poroda:         "Пиловочник",
			Sort:           "1 сорт",
			DlinaBrevna:    "6 м",
			DiameterBrevna: "24 см",
		},
		{
			Id:             2,
			WoodSpecTypeId: 2,
			Poroda:         "Сосна",
			Sort:           "сорт А",
			DlinaBrevna:    "6 м",
			SechenieDoski:  "100*20 мм",
		},
		{
			Id:               3,
			WoodSpecTypeId:   3,
			Poroda:           "Пиломатериалы сухие",
			Sort:             "сорт А",
			DlinaBrevna:      "6 м",
			SechenieDoski:    "100 * 20 мм",
			ProcentVlajnosti: 18,
		},
		{
			Id:             4,
			WoodSpecTypeId: 4,
			Sort:           "1 сорт",
			DiameterGranyl: "8 мм",
		},
		{
			Id:             5,
			WoodSpecTypeId: 5,
			Poroda:         "Сосна",
		},
	})
}

func insertExporters() {
	InsertDefaultEntityData(&ExporterService{}, []models.Exporter{
		{
			Id:          1,
			Name:        "Кемлес",
			Description: "описание",
			Receivers:   nil,
		},
	})
	InsertDefaultEntityData(&ExporterService{}, []models.Exporter{
		{
			Id:          2,
			Name:        "Сиблес",
			Description: "лучший поставщик",
			Receivers:   nil,
		},
	})
}

func insertReceivers() {
	InsertDefaultEntityData(&ReceiverService{}, []models.Receiver{
		{
			Id:               1,
			ExporterId:       1,
			ExporterQuantity: 3,
			FactoryQuantity:  4,
			Description:      "test",
			SpecId:           1,
		},
		{
			Id:               2,
			ExporterId:       2,
			ExporterQuantity: 5,
			FactoryQuantity:  6,
			Description:      "any",
			SpecId:           3,
		},
	})
}

func InsertWorkResults() {
	InsertDefaultEntityData(&WorkResultService{}, []models.WorkResult{
		{
			Id:               1,
			MaterialId:       1,
			MaterialQuantity: 10,
			ProductSpecId:    2,
			WorkDate:         time.Now().Add(time.Hour * 48).UnixMilli(),
		},
		{
			Id:               2,
			MaterialId:       1,
			MaterialQuantity: 0,
			ProductSpecId:    2,
			WorkDate:         time.Now().Add(time.Hour * 96).UnixMilli(),
		},
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

func insertSushkaResults() {
	InsertDefaultEntityData(&SushkaResultService{}, []models.SushkaResult{
		{
			Id:               1,
			Description:      "сушка",
			MaterialSpecId:   2,
			MaterialQuantity: 8,
			ProductSpecId:    3,
			ProductAmount:    9,
			WorkDate:         time.Now().Add(time.Hour * 48).UnixMilli(),
		},
		{
			Id:               2,
			Description:      "сушка",
			MaterialSpecId:   2,
			MaterialQuantity: 10,
			ProductSpecId:    3,
			ProductAmount:    10,
			WorkDate:         time.Now().Add(time.Hour * 96).UnixMilli(),
		},
		{
			Id:               3,
			Description:      "сушка",
			MaterialSpecId:   2,
			MaterialQuantity: 14,
			ProductSpecId:    3,
			ProductAmount:    16,
			WorkDate:         time.Now().Add(time.Hour * 128).UnixMilli(),
		},
	})
}

func insertDryModes() {
	InsertDefaultEntityData(&DryModeService{}, []models.DryMode{
		{
			Id:              1,
			Name:            "Легкий",
			WetMaterialsId:  2,
			HumidityPercent: 30,
			ProcentYsyshki:  20,
		},
		{
			Id:              2,
			Name:            "Средний",
			WetMaterialsId:  2,
			WetMaterials:    models.WoodSpec{},
			HumidityPercent: 40,
			ProcentYsyshki:  30,
		},
		{
			Id:              3,
			Name:            "Сильный",
			WetMaterialsId:  2,
			WetMaterials:    models.WoodSpec{},
			HumidityPercent: 40,
			ProcentYsyshki:  30,
		},
	})
}
