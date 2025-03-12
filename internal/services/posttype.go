package services

import (
	"app/internal/dal"
	"app/internal/database"
	"app/internal/dialogs"
	"app/internal/extras/excel"
	"app/internal/models"
	"app/internal/utils"
	"errors"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"strconv"
)

type PostTypeService struct {
}
type PostType = models.PostType

func (service *PostTypeService) Create(item PostType) (PostType, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.PostType.Preload(field.Associations).Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *PostTypeService) GetAll() ([]*PostType, error) {
	var posttypes []*PostType
	posttypes, err := dal.PostType.Preload(field.Associations).Find()
	return posttypes, err
}
func (service *PostTypeService) GetById(id uint) (*PostType, error) {
	item, err := dal.PostType.Preload(field.Associations).Where(dal.PostType.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *PostTypeService) Update(item PostType) (PostType, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.PostType.Preload(field.Associations).Save(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)
	if err != nil {
		return item, err
	}
	return item, err
}
func (service *PostTypeService) Delete(id uint) error {
	_, err := dal.PostType.Unscoped().Where(dal.PostType.Id.Eq(id)).Delete()
	return err
}
func (service *PostTypeService) Count() (int64, error) {
	amount, err := dal.PostType.Count()
	return amount, err
}

func (service *PostTypeService) ImportFromExcel() error {
	filepath, err := dialogs.OpenFileDialog("Импорт данных")
	if err != nil {
		return err
	}
	err = excel.ImportFromSpreadsheet(filepath, excel.Importer{
		SheetName: "Тип поста",
		Loader: func(rowIndex int, row []string) error {
			id, err := strconv.Atoi(row[0])
			if err != nil {
				return err
			}

			_, err = service.Create(PostType{
				Id:   uint(id),
				Name: row[1],
			})
			if err != nil {
				return err
			}
			return nil
		},
	})
	if err != nil {
		return err
	}
	return nil
}
func (service *PostTypeService) SortedByOrder(fieldsSortOrder map[string]string) ([]*PostType, error) {
	return utils.SortByOrder(fieldsSortOrder, PostType{})
}
