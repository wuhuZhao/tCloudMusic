package dao

import (
	"tCloudMusic_rpc/biz/model"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// FindUserById: give a uid to return a user model
func (u *UserDao) FindUserById(id uint) *model.User {
	user := model.User{}
	u.db.Find(&user, id)
	return &user
}

// FindUserByUserNameAndPassword: give a username and password, check a user is or isn't exsited
func (u *UserDao) FindUserByUserNameAndPassword(username, password string) *model.User {
	user := model.User{}
	u.db.Where("username = ? and password = ?", username, password).Find(&user)
	return &user
}
