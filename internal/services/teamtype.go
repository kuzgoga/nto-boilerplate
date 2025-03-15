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

type TeamTypeService struct {
}
type TeamType = models.TeamType

func (service *TeamTypeService) Create(item TeamType) (TeamType, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.TeamType.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *TeamTypeService) GetAll() ([]*TeamType, error) {
	var teamtypes []*TeamType
	teamtypes, err := dal.TeamType.Preload(field.Associations).Find()
	return teamtypes, err
}
func (service *TeamTypeService) GetById(id uint) (*TeamType, error) {
	item, err := dal.TeamType.Preload(field.Associations).Where(dal.TeamType.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *TeamTypeService) Update(item TeamType) (TeamType, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.TeamType.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *TeamTypeService) Delete(id uint) error {
	_, err := dal.TeamType.Unscoped().Where(dal.TeamType.Id.Eq(id)).Delete()
	return err
}
func (service *TeamTypeService) Count() (int64, error) {
	amount, err := dal.TeamType.Count()
	return amount, err
}
func (service *TeamTypeService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*TeamType, error) {
	return utils.SortByOrder(fieldsSortingOrder, TeamType{})
}
func (service *TeamTypeService) SearchByAllTextFields(phrase string) ([]*TeamType, error) {
	return utils.FindPhraseByStringFields[TeamType](phrase, TeamType{})
}
