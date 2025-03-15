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

type WorkshopService struct {
}
type Workshop = models.Workshop

func (service *WorkshopService) Create(item Workshop) (Workshop, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Workshop.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *WorkshopService) GetAll() ([]*Workshop, error) {
	var workshops []*Workshop
	workshops, err := dal.Workshop.Preload(field.Associations).Find()
	return workshops, err
}
func (service *WorkshopService) GetById(id uint) (*Workshop, error) {
	item, err := dal.Workshop.Preload(field.Associations).Where(dal.Workshop.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *WorkshopService) Update(item Workshop) (Workshop, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Workshop.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *WorkshopService) Delete(id uint) error {
	_, err := dal.Workshop.Unscoped().Where(dal.Workshop.Id.Eq(id)).Delete()
	return err
}
func (service *WorkshopService) Count() (int64, error) {
	amount, err := dal.Workshop.Count()
	return amount, err
}
func (service *WorkshopService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Workshop, error) {
	return utils.SortByOrder(fieldsSortingOrder, Workshop{})
}
func (service *WorkshopService) SearchByAllTextFields(phrase string) ([]*Workshop, error) {
	return utils.FindPhraseByStringFields[Workshop](phrase, Workshop{})
}
