package entity_mappers

import (
	"auptex.com/botnova/internals/domain/models"
	gormentities "auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToUserDomain(e gormentities.User) models.User {
	return models.User{
		Id:        e.Id,
		Name:      e.Name,
		Email:     e.Email,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToUserAuthDomain(e gormentities.User) models.UserAuth {
	return models.UserAuth{
		UserID:       e.Id,
		Email:        e.Email,
		PasswordHash: e.Password,
	}
}

func ToUserEntity(m models.User, passwordHash string) gormentities.User {
	return gormentities.User{
		Id:        m.Id,
		Name:      m.Name,
		Email:     m.Email,
		Password:  passwordHash,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
