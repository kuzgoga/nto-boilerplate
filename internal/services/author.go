package services

import (
	"app/internal/dal"
	"app/internal/database"
	"app/internal/models"
	"errors"

	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type AuthorService struct{}
type Author = models.Author

func (service *AuthorService) Create(item Author) (Author, error) {
	ReplaceEmptySlicesWithNil(&item)
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
	ReplaceEmptySlicesWithNil(&item)

	_, err := dal.Author.Updates(&item)
	if err != nil {
		return item, err
	}

	err = UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}

func (service *AuthorService) Delete(id uint) error {
	_, err := dal.Author.Unscoped().Where(dal.Author.Id.Eq(id)).Delete()
	return err
}
func (service *AuthorService) Count() (int64, error) {
	amount, err := dal.Author.Count()
	return amount, err
}
