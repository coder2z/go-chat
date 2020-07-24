package repositories

import (
	"auth/auth-svr/models"
	"github.com/jinzhu/gorm"
)

type UserRepositoryImp interface {
	GetUserByName(name string) (*models.User, error)
	AddUser(user *models.User) error
	UpdateUser(user *models.User) error
	DelUser(id int) error
}

func NewUserRepository() UserRepositoryImp {
	return &UserManagerRepository{
		Db: models.DB,
	}
}

type UserManagerRepository struct {
	Db *gorm.DB
}

func (u *UserManagerRepository) GetUserByName(name string) (user *models.User, err error) {
	user = &models.User{}
	err = u.Db.Where("name=?", name).Find(user).Error
	return
}

func (u *UserManagerRepository) AddUser(user *models.User) (err error) {
	err = u.Db.Create(user).Error
	return
}

func (u *UserManagerRepository) UpdateUser(user *models.User) (err error) {
	err = u.Db.Model(&models.User{}).Update(user).Error
	return
}

func (u *UserManagerRepository) DelUser(id int) (err error) {
	err = u.Db.Where("id=?", id).Delete(&models.User{}).Error
	return
}
