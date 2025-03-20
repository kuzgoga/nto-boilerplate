package services

import (
	"errors"
	"github.com/kuzgoga/nto-boilerplate/internal/dal"
	"github.com/kuzgoga/nto-boilerplate/internal/database"
	"github.com/kuzgoga/nto-boilerplate/internal/models"
	"github.com/kuzgoga/nto-boilerplate/internal/utils"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SushkaResultService struct {
}
type SushkaResult = models.SushkaResult

func (service *SushkaResultService) Create(item SushkaResult) (SushkaResult, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.SushkaResult.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *SushkaResultService) GetAll() ([]*SushkaResult, error) {
	var sushkaresults []*SushkaResult
	err := database.GetInstance().Preload(clause.Associations).Preload("ProductSpec.WoodSpecType").Preload("Material").Find(&sushkaresults)
	if err != nil {
		return nil, err.Error
	}
	return sushkaresults, nil
}
func (service *SushkaResultService) GetById(id uint) (*SushkaResult, error) {
	item, err := dal.SushkaResult.Preload(field.Associations).Where(dal.SushkaResult.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *SushkaResultService) Update(item SushkaResult) (SushkaResult, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.SushkaResult.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)
	if err != nil {
		return item, err
	}
	return item, err
}
func (service *SushkaResultService) Delete(id uint) error {
	_, err := dal.SushkaResult.Unscoped().Where(dal.SushkaResult.Id.Eq(id)).Delete()
	return err
}
func (service *SushkaResultService) Count() (int64, error) {
	amount, err := dal.SushkaResult.Count()
	return amount, err
}
func (service *SushkaResultService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*SushkaResult, error) {
	return utils.SortByOrder(fieldsSortingOrder, SushkaResult{})
}
func (service *SushkaResultService) SearchByAllTextFields(phrase string) ([]*SushkaResult, error) {
	return utils.FindPhraseByStringFields[SushkaResult](phrase, SushkaResult{})
}
