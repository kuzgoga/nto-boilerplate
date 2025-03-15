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

type PrepTaskService struct {
}
type PrepTask = models.PrepTask

func (service *PrepTaskService) Create(item PrepTask) (PrepTask, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.PrepTask.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *PrepTaskService) GetAll() ([]*PrepTask, error) {
	var preptasks []*PrepTask
	preptasks, err := dal.PrepTask.Preload(field.Associations).Find()
	return preptasks, err
}
func (service *PrepTaskService) GetById(id uint) (*PrepTask, error) {
	item, err := dal.PrepTask.Preload(field.Associations).Where(dal.PrepTask.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *PrepTaskService) Update(item PrepTask) (PrepTask, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.PrepTask.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *PrepTaskService) Delete(id uint) error {
	_, err := dal.PrepTask.Unscoped().Where(dal.PrepTask.Id.Eq(id)).Delete()
	return err
}
func (service *PrepTaskService) Count() (int64, error) {
	amount, err := dal.PrepTask.Count()
	return amount, err
}
func (service *PrepTaskService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*PrepTask, error) {
	return utils.SortByOrder(fieldsSortingOrder, PrepTask{})
}
func (service *PrepTaskService) SearchByAllTextFields(phrase string) ([]*PrepTask, error) {
	return utils.FindPhraseByStringFields[PrepTask](phrase, PrepTask{})
}
