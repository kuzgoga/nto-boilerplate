package services

import (
	"errors"
	"github.com/kuzgoga/nto-boilerplate/internal/dal"
	"github.com/kuzgoga/nto-boilerplate/internal/database"
	"github.com/kuzgoga/nto-boilerplate/internal/models"
	"github.com/kuzgoga/nto-boilerplate/internal/utils"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type ExporterService struct {
}
type Exporter = models.Exporter

func (service *ExporterService) Create(item Exporter) (Exporter, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Exporter.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *ExporterService) GetAll() ([]*Exporter, error) {
	var exporters []*Exporter
	exporters, err := dal.Exporter.Preload(field.Associations).Find()
	return exporters, err
}
func (service *ExporterService) GetById(id uint) (*Exporter, error) {
	item, err := dal.Exporter.Preload(field.Associations).Where(dal.Exporter.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *ExporterService) Update(item Exporter) (Exporter, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Exporter.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)
	if err != nil {
		return item, err
	}
	return item, err
}
func (service *ExporterService) Delete(id uint) error {
	_, err := dal.Exporter.Unscoped().Where(dal.Exporter.Id.Eq(id)).Delete()
	return err
}
func (service *ExporterService) Count() (int64, error) {
	amount, err := dal.Exporter.Count()
	return amount, err
}
func (service *ExporterService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Exporter, error) {
	return utils.SortByOrder(fieldsSortingOrder, Exporter{})
}
func (service *ExporterService) SearchByAllTextFields(phrase string) ([]*Exporter, error) {
	return utils.FindPhraseByStringFields[Exporter](phrase, Exporter{})
}
