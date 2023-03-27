package repository

import (
	"course-forum/infra/database"
	"course-forum/models"
)

func GetTags() (tags []models.Tag, err error) {
	err = database.DB.Find(&tags).Error

	return tags, err
}

func FindTag(id uint) (tag *models.Tag, err error) {
	err = database.DB.First(&tag, id).Error

	return tag, err
}

func CreateTag(tag *models.Tag) (err error) {
	err = database.DB.Create(&tag).Error

	return err
}

func UpdateTag(tag *models.Tag) (err error) {
	err = database.DB.First(&models.Tag{}, tag.ID).Error

	if err == nil {
		err = database.DB.Updates(&tag).Error
	}

	return err
}

func DeleteTag(tag *models.Tag) (err error) {
	err = database.DB.Delete(&tag).Error

	return err
}
