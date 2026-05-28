package repository

import (
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/models"
	"gorm.io/gorm"
)

// UserRepository defines the interface (methods) our repo must have
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
}

// userRepo is the concrete implementation
type userRepo struct {
	db *gorm.DB
}

// NewUserRepository is a constructor function
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}
