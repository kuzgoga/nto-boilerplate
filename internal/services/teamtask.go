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

type TeamTaskService struct {
}
type TeamTask = models.TeamTask

func (service *TeamTaskService) Create(item TeamTask) (TeamTask, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.TeamTask.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *TeamTaskService) GetAll() ([]*TeamTask, error) {
	var teamtasks []*TeamTask
	teamtasks, err := dal.TeamTask.Preload(field.Associations).Find()
	return teamtasks, err
}
func (service *TeamTaskService) GetById(id uint) (*TeamTask, error) {
	item, err := dal.TeamTask.Preload(field.Associations).Where(dal.TeamTask.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *TeamTaskService) Update(item TeamTask) (TeamTask, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.TeamTask.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *TeamTaskService) Delete(id uint) error {
	_, err := dal.TeamTask.Unscoped().Where(dal.TeamTask.Id.Eq(id)).Delete()
	return err
}
func (service *TeamTaskService) Count() (int64, error) {
	amount, err := dal.TeamTask.Count()
	return amount, err
}
func (service *TeamTaskService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*TeamTask, error) {
	return utils.SortByOrder(fieldsSortingOrder, TeamTask{})
}
func (service *TeamTaskService) SearchByAllTextFields(phrase string) ([]*TeamTask, error) {
	return utils.FindPhraseByStringFields[TeamTask](phrase, TeamTask{})
}
