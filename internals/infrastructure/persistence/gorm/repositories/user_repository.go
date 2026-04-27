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

func (r *UserRepository) Create(user models.User) error {
	user.Id = uuid.Must(uuid.NewV7()).String()
	entity := entity_mappers.ToUserEntity(user)
	return r.db.Create(&entity).Error
}

func (r *UserRepository) GetById(id int) (*models.User, error) {
	var entity entities.User
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	data := entity_mappers.ToUserDomain(entity)
	return &data, nil
}

func (r *UserRepository) Update(user models.User) error {
	entity := entity_mappers.ToUserEntity(user)
	return r.db.Save(&entity).Error
}

func (r *UserRepository) Delete(id int) error {
	return r.db.Delete(&entities.User{}, "id = ?", id).Error
}
