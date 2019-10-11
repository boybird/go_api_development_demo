package models

import "github.com/jinzhu/gorm"

// Tag 标签类
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Model(&Tag{}).Where("name = ? AND deleted_on = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return true, nil

}

// AddTag Add a tag
func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil

}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, err
}

// GetTagTotal counts the total number of tags based on the constraint
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	err := db.Where(maps).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil

}

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

// DeleteTag delete a tag
func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

// EditTag 编辑标签
func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// CleanAllTag 清空所有标签
func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("deleted_on = ?", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
