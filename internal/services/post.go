package services

import (
	"app/internal/dal"
	"app/internal/dialogs"
	excel "app/internal/extras/excel"
	"app/internal/models"
	"errors"
	"fmt"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type PostService struct{}
type Post = models.Post

func (service *PostService) Create(item Post) (Post, error) {
	err := dal.Post.Preload(field.Associations).Create(&item)
	return item, err
}
func (service *PostService) GetAll() ([]*Post, error) {
	var posts []*Post
	posts, err := dal.Post.Preload(field.Associations).Find()
	return posts, err
}
func (service *PostService) GetById(id uint) (*Post, error) {
	item, err := dal.Post.Preload(field.Associations).Where(dal.Post.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *PostService) Update(item Post) (Post, error) {
	err := dal.Post.Preload(field.Associations).Save(&item)
	return item, err
}
func (service *PostService) Delete(item Post) (Post, error) {
	_, err := dal.Post.Unscoped().Preload(field.Associations).Delete(&item)
	return item, err
}
func (service *PostService) Count() (int64, error) {
	amount, err := dal.Post.Count()
	return amount, err
}

func (service *PostService) ExportToExcel() {
	err := excel.ExportEntityToSpreadsheet("report.xlsx", "Посты", Post{}, service.GetAll)
	if err != nil {
		dialogs.ErrorDialog("Ошибка экспорта", fmt.Sprintf("Ошибка при экспорте данных: %s", err))
	}
}
