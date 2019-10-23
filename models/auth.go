package models

import "github.com/jinzhu/gorm"

// Auth 用户认证
type Auth struct {
	ID       int    `gorm:"primary_Key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 检查登录
func CheckAuth(username string, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil

}
