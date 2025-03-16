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

type PostService struct {
}
type Post = models.Post

func (service *PostService) Create(item Post) (Post, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Post.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
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
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Post.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *PostService) Delete(id uint) error {
	_, err := dal.Post.Unscoped().Where(dal.Post.Id.Eq(id)).Delete()
	return err
}
func (service *PostService) Count() (int64, error) {
	amount, err := dal.Post.Count()
	return amount, err
}
func (service *PostService) SortedByOrder(fieldsSortingOrder []utils.SortField) ([]*Post, error) {
	return utils.SortByOrder(fieldsSortingOrder, Post{})
}
func (service *PostService) SearchByAllTextFields(phrase string) ([]*Post, error) {
	return utils.FindPhraseByStringFields[Post](phrase, Post{})
}
