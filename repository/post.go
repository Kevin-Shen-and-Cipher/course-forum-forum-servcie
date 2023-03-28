package repository

import (
	"course-forum/infra/database"
	"course-forum/models"

	"gorm.io/gorm/clause"
)

func GetPosts() (posts []models.Post, err error) {
	err = database.DB.Preload(clause.Associations).Find(&posts).Order("id DESC").Error

	return posts, err
}

func FindPost(id uint) (post *models.Post, err error) {
	err = database.DB.First(&post, id).Error

	if err == nil {
		post.Views++
		err = database.DB.Preload(clause.Associations).Updates(&post).Error
	}

	return post, err
}

func CreatePost(post *models.Post) (err error) {
	err = database.DB.Create(&post).Error

	return err
}

func UpdatePost(post *models.Post) (err error) {
	err = database.DB.First(&models.Post{}, post.ID).Error

	if err == nil {
		err = database.DB.Updates(&post).Error
	}

	return err
}

func DeletePost(post *models.Post) (err error) {
	err = database.DB.Select(clause.Associations).Delete(&post).Error

	return err
}
