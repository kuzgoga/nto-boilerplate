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

type OrderService struct {
}
type Order = models.Order

func (service *OrderService) Create(item Order) (Order, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Order.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *OrderService) GetAll() ([]*Order, error) {
	var orders []*Order
	orders, err := dal.Order.Preload(field.Associations).Find()
	return orders, err
}
func (service *OrderService) GetById(id uint) (*Order, error) {
	item, err := dal.Order.Preload(field.Associations).Where(dal.Order.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *OrderService) Update(item Order) (Order, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Order.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *OrderService) Delete(id uint) error {
	_, err := dal.Order.Unscoped().Where(dal.Order.Id.Eq(id)).Delete()
	return err
}
func (service *OrderService) Count() (int64, error) {
	amount, err := dal.Order.Count()
	return amount, err
}
func (service *OrderService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Order, error) {
	return utils.SortByOrder(fieldsSortingOrder, Order{})
}
func (service *OrderService) SearchByAllTextFields(phrase string) ([]*Order, error) {
	return utils.FindPhraseByStringFields[Order](phrase, Order{})
}
