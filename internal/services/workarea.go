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

type WorkAreaService struct {
}
type WorkArea = models.WorkArea

func (service *WorkAreaService) Create(item WorkArea) (WorkArea, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.WorkArea.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *WorkAreaService) GetAll() ([]*WorkArea, error) {
	var workareas []*WorkArea
	workareas, err := dal.WorkArea.Preload(field.Associations).Find()
	return workareas, err
}
func (service *WorkAreaService) GetById(id uint) (*WorkArea, error) {
	item, err := dal.WorkArea.Preload(field.Associations).Where(dal.WorkArea.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *WorkAreaService) Update(item WorkArea) (WorkArea, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.WorkArea.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *WorkAreaService) Delete(id uint) error {
	_, err := dal.WorkArea.Unscoped().Where(dal.WorkArea.Id.Eq(id)).Delete()
	return err
}
func (service *WorkAreaService) Count() (int64, error) {
	amount, err := dal.WorkArea.Count()
	return amount, err
}
func (service *WorkAreaService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*WorkArea, error) {
	return utils.SortByOrder(fieldsSortingOrder, WorkArea{})
}
func (service *WorkAreaService) SearchByAllTextFields(phrase string) ([]*WorkArea, error) {
	return utils.FindPhraseByStringFields[WorkArea](phrase, WorkArea{})
}
