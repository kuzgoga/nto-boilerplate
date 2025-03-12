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

type CommentService struct {
}
type Comment = models.Comment

func (service *CommentService) Create(item Comment) (Comment, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	err := dal.Comment.Create(&item)
	if err != nil {
		return item, err
	}
	err = utils.AppendAssociations(database.GetInstance(), &item)
	return item, err
}
func (service *CommentService) GetAll() ([]*Comment, error) {
	var comments []*Comment
	comments, err := dal.Comment.Preload(field.Associations).Find()
	return comments, err
}
func (service *CommentService) GetById(id uint) (*Comment, error) {
	item, err := dal.Comment.Preload(field.Associations).Where(dal.Comment.Id.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return item, nil
}
func (service *CommentService) Update(item Comment) (Comment, error) {
	utils.ReplaceEmptySlicesWithNil(&item)
	_, err := dal.Author.Updates(&item)
	if err != nil {
		return item, err
	}
	err = utils.UpdateAssociations(database.GetInstance(), &item)

	if err != nil {
		return item, err
	}

	return item, err
}
func (service *CommentService) Delete(id uint) error {
	_, err := dal.Comment.Unscoped().Where(dal.Comment.Id.Eq(id)).Delete()
	return err
}
func (service *CommentService) Count() (int64, error) {
	amount, err := dal.Comment.Count()
	return amount, err
}
func (service *CommentService) SortedByOrder(fieldsSortOrder map[string]string) ([]*Comment, error) {
	return utils.SortByOrder(fieldsSortOrder, Comment{})
}
