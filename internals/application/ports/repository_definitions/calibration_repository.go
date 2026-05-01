package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type CalibrationRepository interface {
	Create(calibration *models.CalibrationProfile) error
	GetById(id string) (*models.CalibrationProfile, error)
	GetByRobotId(robotId string) (*models.CalibrationProfile, error)
	Update(calibration *models.CalibrationProfile) error
	Delete(id string) error
}
