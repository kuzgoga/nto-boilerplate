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

type PostavService struct {
}
type Postav = models.Postav

func (service *PostavService) Create(item Postav) (Postav, error) {
	if item.CenterOutPercent+item.OpilkiOutPercent+item.BacksideOutPercent != 100 {
		return Postav{}, errors.New("неверные входные данные, проценты выхода в сумме должны составлять 100%")
	}
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Postav.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *PostavService) GetAll() ([]*Postav, error) {
	var postavs []*Postav
	postavs, err := dal.Postav.Preload(field.Associations).Find()
	return postavs, err
}
func (service *PostavService) GetById(id uint) (*Postav, error) {
	item, err := dal.Postav.Preload(field.Associations).Where(dal.Postav.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *PostavService) Update(item Postav) (Postav, error) {
	if item.CenterOutPercent+item.OpilkiOutPercent+item.BacksideOutPercent != 100 {
		return Postav{}, errors.New("неверные входные данные, проценты выхода в сумме должны составлять 100%")
	}
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Postav.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *PostavService) Delete(id uint) error {
	_, err := dal.Postav.Unscoped().Where(dal.Postav.Id.Eq(id)).Delete()
	return err
}
func (service *PostavService) Count() (int64, error) {
	amount, err := dal.Postav.Count()
	return amount, err
}
func (service *PostavService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Postav, error) {
	return utils.SortByOrder(fieldsSortingOrder, Postav{})
}
func (service *PostavService) SearchByAllTextFields(phrase string) ([]*Postav, error) {
	return utils.FindPhraseByStringFields[Postav](phrase, Postav{})
}
