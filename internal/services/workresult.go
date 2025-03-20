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

type WorkResultService struct {
}
type WorkResult = models.WorkResult

func (service *WorkResultService) Create(item WorkResult) (WorkResult, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.WorkResult.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *WorkResultService) GetAll() ([]*WorkResult, error) {
	var workresults []*WorkResult
	workresults, err := dal.WorkResult.Preload(field.Associations).Find()
	return workresults, err
}
func (service *WorkResultService) GetById(id uint) (*WorkResult, error) {
	item, err := dal.WorkResult.Preload(field.Associations).Where(dal.WorkResult.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *WorkResultService) Update(item WorkResult) (WorkResult, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.WorkResult.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)
	if err != nil {
		return item, err
	}
	return item, err
}
func (service *WorkResultService) Delete(id uint) error {
	_, err := dal.WorkResult.Unscoped().Where(dal.WorkResult.Id.Eq(id)).Delete()
	return err
}
func (service *WorkResultService) Count() (int64, error) {
	amount, err := dal.WorkResult.Count()
	return amount, err
}
func (service *WorkResultService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*WorkResult, error) {
	return utils.SortByOrder(fieldsSortingOrder, WorkResult{})
}
func (service *WorkResultService) SearchByAllTextFields(phrase string) ([]*WorkResult, error) {
	return utils.FindPhraseByStringFields[WorkResult](phrase, WorkResult{})
}
func (service *WorkResultService) GetCompMaterialSpecs() ([]*WoodSpec, error) {
	items, err := dal.WoodSpec.Preload(field.Associations).Preload(dal.WoodSpec.WoodSpecType).Find()
	if err != nil {
		return nil, err
	}
	var filtered []*WoodSpec
	for _, item := range items {
		if item.WoodSpecTypeId == 1 {
			filtered = append(filtered, item)
		}
	}
	return filtered, err
}
func (service *WorkResultService) GetCompProductSpecs() ([]*WoodSpec, error) {
	items, err := dal.WoodSpec.Preload(field.Associations).Preload(dal.WoodSpec.WoodSpecType).Find()
	if err != nil {
		return nil, err
	}
	var filtered []*WoodSpec
	for _, item := range items {
		if item.WoodSpecTypeId == 2 {
			filtered = append(filtered, item)
		}
	}
	return filtered, err
}
