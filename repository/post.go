package repository

import (
	"course-forum/infra/database"
	"course-forum/models"
)

func GetPosts() (posts []models.Post, err error) {
	err = database.DB.Find(&posts).Error

	return posts, err
}

func FindPost(id uint) (post *models.Post, err error) {
	err = database.DB.First(&post, id).Error

	post.Views++
	err = database.DB.Updates(&post).Error

	return post, err
}

func CreatePost(post *models.Post) (err error) {
	err = database.DB.Create(&post).Error

	return err
}

func UpdatePost(post *models.Post) (err error) {
	err = database.DB.Updates(&post).Error

	return err
}

func DeletePost(post *models.Post) (err error) {
	err = database.DB.Delete(&post).Error

	return err
}
