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

type WorkerService struct {
}
type Worker = models.Worker

func (service *WorkerService) Create(item Worker) (Worker, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Worker.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *WorkerService) GetAll() ([]*Worker, error) {
	var workers []*Worker
	workers, err := dal.Worker.Preload(field.Associations).Find()
	return workers, err
}
func (service *WorkerService) GetById(id uint) (*Worker, error) {
	item, err := dal.Worker.Preload(field.Associations).Where(dal.Worker.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *WorkerService) Update(item Worker) (Worker, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Worker.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *WorkerService) Delete(id uint) error {
	_, err := dal.Worker.Unscoped().Where(dal.Worker.Id.Eq(id)).Delete()
	return err
}
func (service *WorkerService) Count() (int64, error) {
	amount, err := dal.Worker.Count()
	return amount, err
}
func (service *WorkerService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Worker, error) {
	return utils.SortByOrder(fieldsSortingOrder, Worker{})
}
func (service *WorkerService) SearchByAllTextFields(phrase string) ([]*Worker, error) {
	return utils.FindPhraseByStringFields[Worker](phrase, Worker{})
}
