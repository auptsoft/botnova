package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User, passwordHash string) (*models.User, error) {
	user.Id = uuid.Must(uuid.NewV7()).String()
	entity := entity_mappers.ToUserEntity(user, passwordHash)
	if err := r.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	created := entity_mappers.ToUserDomain(entity)
	return &created, nil
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

func (r *UserRepository) GetAuthByEmail(email string) (*models.UserAuth, error) {
	var entity entities.User
	if err := r.db.First(&entity, "email = ?", email).Error; err != nil {
		return nil, err
	}

	authData := entity_mappers.ToUserAuthDomain(entity)
	return &authData, nil
}

func (r *UserRepository) Update(user models.User) error {
	updates := map[string]any{
		"name":  user.Name,
		"email": user.Email,
	}

	return r.db.Model(&entities.User{}).Where("id = ?", user.Id).Updates(updates).Error
}

func (r *UserRepository) UpdatePassword(userID string, passwordHash string) error {
	return r.db.Model(&entities.User{}).Where("id = ?", userID).Update("password", passwordHash).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&entities.User{}, "id = ?", id).Error
}
