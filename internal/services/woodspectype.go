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

type WoodSpecTypeService struct {
}
type WoodSpecType = models.WoodSpecType

func (service *WoodSpecTypeService) Create(item WoodSpecType) (WoodSpecType, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.WoodSpecType.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *WoodSpecTypeService) GetAll() ([]*WoodSpecType, error) {
	var woodspectypes []*WoodSpecType
	woodspectypes, err := dal.WoodSpecType.Preload(field.Associations).Find()
	return woodspectypes, err
}
func (service *WoodSpecTypeService) GetById(id uint) (*WoodSpecType, error) {
	item, err := dal.WoodSpecType.Preload(field.Associations).Where(dal.WoodSpecType.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *WoodSpecTypeService) Update(item WoodSpecType) (WoodSpecType, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.WoodSpecType.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)
	if err != nil {
		return item, err
	}
	return item, err
}
func (service *WoodSpecTypeService) Delete(id uint) error {
	_, err := dal.WoodSpecType.Unscoped().Where(dal.WoodSpecType.Id.Eq(id)).Delete()
	return err
}
func (service *WoodSpecTypeService) Count() (int64, error) {
	amount, err := dal.WoodSpecType.Count()
	return amount, err
}
func (service *WoodSpecTypeService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*WoodSpecType, error) {
	return utils.SortByOrder(fieldsSortingOrder, WoodSpecType{})
}
func (service *WoodSpecTypeService) SearchByAllTextFields(phrase string) ([]*WoodSpecType, error) {
	return utils.FindPhraseByStringFields[WoodSpecType](phrase, WoodSpecType{})
}
