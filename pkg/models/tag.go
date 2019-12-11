package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name string `json:"name"`
}

// AddTag Add a Tag
func AddTag(name string, state int) (bool, error) {
	tag := Tag{
		Name: name,
	}
	if err := db.Create(&tag).Error; err != nil {
		return false, err
	}
	return true, nil
}
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)
	fmt.Print(db == nil)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}
