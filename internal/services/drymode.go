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

type DryModeService struct {
}
type DryMode = models.DryMode

func (service *DryModeService) Create(item DryMode) (DryMode, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.DryMode.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *DryModeService) GetAll() ([]*DryMode, error) {
	var drymodes []*DryMode
	drymodes, err := dal.DryMode.Preload(field.Associations).Find()
	return drymodes, err
}
func (service *DryModeService) GetById(id uint) (*DryMode, error) {
	item, err := dal.DryMode.Preload(field.Associations).Where(dal.DryMode.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *DryModeService) Update(item DryMode) (DryMode, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.DryMode.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *DryModeService) Delete(id uint) error {
	_, err := dal.DryMode.Unscoped().Where(dal.DryMode.Id.Eq(id)).Delete()
	return err
}
func (service *DryModeService) Count() (int64, error) {
	amount, err := dal.DryMode.Count()
	return amount, err
}
func (service *DryModeService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*DryMode, error) {
	return utils.SortByOrder(fieldsSortingOrder, DryMode{})
}
func (service *DryModeService) SearchByAllTextFields(phrase string) ([]*DryMode, error) {
	return utils.FindPhraseByStringFields[DryMode](phrase, DryMode{})
}
