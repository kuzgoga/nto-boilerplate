package services

import (
	"app/internal/dal"
	"app/internal/models"
	"errors"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type AuthorService struct{}
type Author = models.Author

func (service *AuthorService) Create(item Author) (Author, error) {
	err := dal.Author.Preload(field.Associations).Create(&item)
	return item, err
}
func (service *AuthorService) GetAll() ([]*Author, error) {
	var authors []*Author
	authors, err := dal.Author.Preload(field.Associations).Find()
	return authors, err
}
func (service *AuthorService) GetById(id uint) (*Author, error) {
	item, err := dal.Author.Preload(field.Associations).Where(dal.Author.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *AuthorService) Update(item Author) (Author, error) {
	err := dal.Author.Preload(field.Associations).Save(&item)
	return item, err
}
func (service *AuthorService) Delete(item Author) (Author, error) {
	_, err := dal.Author.Unscoped().Preload(field.Associations).Delete(&item)
	return item, err
}
func (service *AuthorService) Count() (int64, error) {
	amount, err := dal.Author.Count()
	return amount, err
}
