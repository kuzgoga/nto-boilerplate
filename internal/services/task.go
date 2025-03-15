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

type TaskService struct {
}
type Task = models.Task

func (service *TaskService) Create(item Task) (Task, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Task.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *TaskService) GetAll() ([]*Task, error) {
	var tasks []*Task
	tasks, err := dal.Task.Preload(field.Associations).Find()
	return tasks, err
}
func (service *TaskService) GetById(id uint) (*Task, error) {
	item, err := dal.Task.Preload(field.Associations).Where(dal.Task.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *TaskService) Update(item Task) (Task, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Task.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *TaskService) Delete(id uint) error {
	_, err := dal.Task.Unscoped().Where(dal.Task.Id.Eq(id)).Delete()
	return err
}
func (service *TaskService) Count() (int64, error) {
	amount, err := dal.Task.Count()
	return amount, err
}
func (service *TaskService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Task, error) {
	return utils.SortByOrder(fieldsSortingOrder, Task{})
}
func (service *TaskService) SearchByAllTextFields(phrase string) ([]*Task, error) {
	return utils.FindPhraseByStringFields[Task](phrase, Task{})
}
