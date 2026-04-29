package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User) error {
	user.Id = uuid.Must(uuid.NewV7()).String()

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	entity := entity_mappers.ToUserEntity(user)
	return r.db.Create(&entity).Error
}

func (r *UserRepository) GetById(id string) (*models.User, error) {
	var entity entities.User
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	data := entity_mappers.ToUserDomain(entity)
	return &data, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var entity entities.User
	if err := r.db.First(&entity, "email = ?", email).Error; err != nil {
		return nil, err
	}

	data := entity_mappers.ToUserDomain(entity)
	return &data, nil
}

func (r *UserRepository) Update(user models.User) error {
	updates := map[string]any{
		"name":  user.Name,
		"email": user.Email,
	}

	if user.Password != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updates["password"] = string(bytes)
	}

	return r.db.Model(&entities.User{}).Where("id = ?", user.Id).Updates(updates).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&entities.User{}, "id = ?", id).Error
}
