package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name string `json:"name"`
}

// AddTag Add a Tag
func AddTag(name string) (bool, error) {
	tag := Tag{
		Name: name,
	}
	if err := db.FirstOrCreate(&tag, Tag{Name: name}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var tags []Tag
	err := db.Find(&tags).Offset(pageNum).Limit(pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tags, err
	}
	return tags, err
}

func CheckTag(where interface{}) (bool, Tag) {
	var tags Tag
	err := db.Where(where).First(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, tags
	}
	if tags.ID > 0 {
		return true, tags
	}
	return false, tags
}

func EditTags(id int, data interface{}) bool {
	print(id)
	err := db.Model(&Tag{}).Where("id= ?", id).Updates(data).Error
	return err == nil
}

func DeleteTags(id int) bool {
	err := db.Unscoped().Delete(&Tag{}, id).Error
	return err == nil
}
