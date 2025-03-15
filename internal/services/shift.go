package services

import (
	"app/internal/dal"
	"app/internal/database"
	"app/internal/models"
	"app/internal/utils"
	"errors"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type ShiftService struct {
}
type Shift = models.Shift

func (service *ShiftService) Create(item Shift) (Shift, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Shift.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *ShiftService) GetAll() ([]*Shift, error) {
	var shifts []*Shift
	shifts, err := dal.Shift.Preload(field.Associations).Find()
	return shifts, err
}
func (service *ShiftService) GetById(id uint) (*Shift, error) {
	item, err := dal.Shift.Preload(field.Associations).Where(dal.Shift.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *ShiftService) Update(item Shift) (Shift, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Shift.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *ShiftService) Delete(id uint) error {
	_, err := dal.Shift.Unscoped().Where(dal.Shift.Id.Eq(id)).Delete()
	return err
}
func (service *ShiftService) Count() (int64, error) {
	amount, err := dal.Shift.Count()
	return amount, err
}
func (service *ShiftService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Shift, error) {
	return utils.SortByOrder(fieldsSortingOrder, Shift{})
}
func (service *ShiftService) SearchByAllTextFields(phrase string) ([]*Shift, error) {
	return utils.FindPhraseByStringFields[Shift](phrase, Shift{})
}
