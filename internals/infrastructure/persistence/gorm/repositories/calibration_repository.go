package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type CalibrationRepository struct {
	db *gorm.DB
}

func NewCalibrationRepository(db *gorm.DB) *CalibrationRepository {
	return &CalibrationRepository{db: db}
}

func (r *CalibrationRepository) Create(calibration *models.CalibrationProfile) error {
	calibration.Id = uuid.Must(uuid.NewV7()).String()
	entity, err := entity_mappers.ToCalibrationEntity(calibration)
	if err != nil {
		return err
	}
	return r.db.Create(&entity).Error
}

func (r *CalibrationRepository) GetById(id string) (*models.CalibrationProfile, error) {
	var entity entities.CalibrationEntity
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	data, err := entity_mappers.ToCalibrationDomain(&entity)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CalibrationRepository) GetByRobotId(robotId string) (*models.CalibrationProfile, error) {
	var entity entities.CalibrationEntity
	if err := r.db.First(&entity, "robot_id = ?", robotId).Error; err != nil {
		return nil, err
	}

	data, err := entity_mappers.ToCalibrationDomain(&entity)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CalibrationRepository) Update(calibration *models.CalibrationProfile) error {
	entity, err := entity_mappers.ToCalibrationEntity(calibration)
	if err != nil {
		return err
	}
	return r.db.Save(&entity).Error
}

func (r *CalibrationRepository) Delete(id string) error {
	return r.db.Delete(&entities.CalibrationEntity{}, "id = ?", id).Error
}
