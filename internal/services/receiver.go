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

type ReceiverService struct {
}
type Receiver = models.Receiver

func (service *ReceiverService) Create(item Receiver) (Receiver, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Receiver.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *ReceiverService) GetAll() ([]*Receiver, error) {
	var receivers []*Receiver
	receivers, err := dal.Receiver.Preload(field.Associations).Find()
	return receivers, err
}
func (service *ReceiverService) GetById(id uint) (*Receiver, error) {
	item, err := dal.Receiver.Preload(field.Associations).Where(dal.Receiver.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *ReceiverService) Update(item Receiver) (Receiver, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Receiver.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)
	if err != nil {
		return item, err
	}
	return item, err
}
func (service *ReceiverService) Delete(id uint) error {
	_, err := dal.Receiver.Unscoped().Where(dal.Receiver.Id.Eq(id)).Delete()
	return err
}
func (service *ReceiverService) Count() (int64, error) {
	amount, err := dal.Receiver.Count()
	return amount, err
}
func (service *ReceiverService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Receiver, error) {
	return utils.SortByOrder(fieldsSortingOrder, Receiver{})
}
func (service *ReceiverService) SearchByAllTextFields(phrase string) ([]*Receiver, error) {
	return utils.FindPhraseByStringFields[Receiver](phrase, Receiver{})
}
