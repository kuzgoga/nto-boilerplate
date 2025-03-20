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

type WoodSpecService struct {
}
type WoodSpec = models.WoodSpec

func (service *WoodSpecService) Create(item WoodSpec) (WoodSpec, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.WoodSpec.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *WoodSpecService) GetAll() ([]*WoodSpec, error) {
	var woodspecs []*WoodSpec
	woodspecs, err := dal.WoodSpec.Preload(field.Associations).Find()
	return woodspecs, err
}
func (service *WoodSpecService) GetById(id uint) (*WoodSpec, error) {
	item, err := dal.WoodSpec.Preload(field.Associations).Where(dal.WoodSpec.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *WoodSpecService) Update(item WoodSpec) (WoodSpec, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.WoodSpec.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)
	if err != nil {
		return item, err
	}
	return item, err
}
func (service *WoodSpecService) Delete(id uint) error {
	_, err := dal.WoodSpec.Unscoped().Where(dal.WoodSpec.Id.Eq(id)).Delete()
	return err
}
func (service *WoodSpecService) Count() (int64, error) {
	amount, err := dal.WoodSpec.Count()
	return amount, err
}
func (service *WoodSpecService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*WoodSpec, error) {
	return utils.SortByOrder(fieldsSortingOrder, WoodSpec{})
}
func (service *WoodSpecService) SearchByAllTextFields(phrase string) ([]*WoodSpec, error) {
	return utils.FindPhraseByStringFields[WoodSpec](phrase, WoodSpec{})
}
