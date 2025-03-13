package infra

import (
	userDomain "qualthea-api/internal/app/user/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id int) (userDomain.User, error)
	GetByUsername(username string) (*userDomain.User, error)
	GetAll() ([]userDomain.User, error)
}

type UserRepositoryImpl struct {
	Db   *gorm.DB
	Repo UserRepository
}

func NewRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// Get user by ID
func (r *UserRepositoryImpl) GetByID(id int) (userDomain.User, error) {
	var user userDomain.User

	result := r.Db.First(&user, id)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// Get user by username
func (r *UserRepositoryImpl) GetByUsername(username string) (*userDomain.User, error) {
	var user userDomain.User

	result := r.Db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// Get all users
func (r *UserRepositoryImpl) GetAll() ([]userDomain.User, error) {
	var users []userDomain.User

	result := r.Db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
