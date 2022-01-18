package users

import (
	"project-e-commerces/entities"

	"gorm.io/gorm"
)

type UserInterface interface {
	// Register and Login
	Register(user entities.User) (entities.User, error)
	Login(email, password string) (entities.User, error)
	// CRUD
	Get(userId int) (entities.User, error)
	Update(customer entities.User) (entities.User, error)
	Delete(userId int) error
}

type UserStructRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserStructRepository {
	return &UserStructRepository{db}
}

func (ur *UserStructRepository) Register(user entities.User) (entities.User, error) {
	if err := ur.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserStructRepository) Login(email, password string) (entities.User, error) {
	login := entities.User{
		Email:    email,
		Password: password,
	}
	if err := ur.db.First(&login).Error; err != nil {
		return login, nil
	}
	return login, nil
}

func (ur *UserStructRepository) Get(userId int) (entities.User, error) {
	var user entities.User
	if err := ur.db.Find(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil

}
func (ur *UserStructRepository) Update(user entities.User) (entities.User, error) {
	err := ur.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (ur *UserStructRepository) Delete(id int) error {
	var user entities.User
	err := ur.db.Unscoped().Delete(&user, id).Error
	if err != nil {
		return err
	}
	return err
}
