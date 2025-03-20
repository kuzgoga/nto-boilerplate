package services

import (
	"github.com/kuzgoga/nto-boilerplate/internal/database"
	"gorm.io/gorm/clause"
)

type DryReport struct {
	Materials       []string
	MaterialsAmount int
}
type DryReportService struct{}

func (service DryReportService) GetDryMaterialsList() (DryReport, error) {
	var items []SushkaResult
	err := database.GetInstance().Preload(clause.Associations).Preload("ProductSpec.WoodSpecType").Find(&items).Error
	if err != nil {
		return DryReport{}, err
	}
	var materials []string
	for _, item := range items {
		materials = append(materials, item.ProductSpec.WoodSpecType.Name)
	}
	return DryReport{Materials: materials, MaterialsAmount: len(items)}, nil
}
