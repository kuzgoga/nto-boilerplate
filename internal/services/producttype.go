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

type ProductTypeService struct {
}
type ProductType = models.ProductType

func (service *ProductTypeService) Create(item ProductType) (ProductType, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.ProductType.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *ProductTypeService) GetAll() ([]*ProductType, error) {
	var producttypes []*ProductType
	producttypes, err := dal.ProductType.Preload(field.Associations).Find()
	return producttypes, err
}
func (service *ProductTypeService) GetById(id uint) (*ProductType, error) {
	item, err := dal.ProductType.Preload(field.Associations).Where(dal.ProductType.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *ProductTypeService) Update(item ProductType) (ProductType, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.ProductType.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *ProductTypeService) Delete(id uint) error {
	_, err := dal.ProductType.Unscoped().Where(dal.ProductType.Id.Eq(id)).Delete()
	return err
}
func (service *ProductTypeService) Count() (int64, error) {
	amount, err := dal.ProductType.Count()
	return amount, err
}
func (service *ProductTypeService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*ProductType, error) {
	return utils.SortByOrder(fieldsSortingOrder, ProductType{})
}
func (service *ProductTypeService) SearchByAllTextFields(phrase string) ([]*ProductType, error) {
	return utils.FindPhraseByStringFields[ProductType](phrase, ProductType{})
}
