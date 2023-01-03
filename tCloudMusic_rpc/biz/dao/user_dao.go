package dao

import (
	"tCloudMusic_rpc/biz/model"

	"github.com/cloudwego/kitex/pkg/klog"
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

// InsertUse: insert user to db
func (u *UserDao) InsertUser(user *model.User) (*model.User, error) {
	result := u.db.Create(user)
	if result.RowsAffected != 0 {
		klog.Errorf("gorm insert user error: %v\n", result.Error.Error())
	}
	return user, result.Error
}
