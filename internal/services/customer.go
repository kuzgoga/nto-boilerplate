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

type CustomerService struct {
}
type Customer = models.Customer

func (service *CustomerService) Create(item Customer) (Customer, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Customer.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *CustomerService) GetAll() ([]*Customer, error) {
	var customers []*Customer
	customers, err := dal.Customer.Preload(field.Associations).Find()
	return customers, err
}
func (service *CustomerService) GetById(id uint) (*Customer, error) {
	item, err := dal.Customer.Preload(field.Associations).Where(dal.Customer.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *CustomerService) Update(item Customer) (Customer, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Customer.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *CustomerService) Delete(id uint) error {
	_, err := dal.Customer.Unscoped().Where(dal.Customer.Id.Eq(id)).Delete()
	return err
}
func (service *CustomerService) Count() (int64, error) {
	amount, err := dal.Customer.Count()
	return amount, err
}
func (service *CustomerService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Customer, error) {
	return utils.SortByOrder(fieldsSortingOrder, Customer{})
}
func (service *CustomerService) SearchByAllTextFields(phrase string) ([]*Customer, error) {
	return utils.FindPhraseByStringFields[Customer](phrase, Customer{})
}
