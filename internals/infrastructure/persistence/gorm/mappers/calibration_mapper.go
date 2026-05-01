package entity_mappers

import (
	"encoding/json"

	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToCalibrationEntity(model *models.CalibrationProfile) (*entities.CalibrationEntity, error) {
	commandsJSON, err := json.Marshal(model.Commands)
	if err != nil {
		return nil, err
	}

	return &entities.CalibrationEntity{
		Id:           model.Id,
		RobotId:      model.RobotId,
		Enabled:      model.Enabled,
		Version:      model.Version,
		CommandsJSON: string(commandsJSON),
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}, nil
}

func ToCalibrationDomain(entity *entities.CalibrationEntity) (*models.CalibrationProfile, error) {
	var commands map[string]models.CommandCalibration
	if err := json.Unmarshal([]byte(entity.CommandsJSON), &commands); err != nil {
		return nil, err
	}

	return &models.CalibrationProfile{
		Id:       entity.Id,
		RobotId:  entity.RobotId,
		Enabled:  entity.Enabled,
		Version:  entity.Version,
		Commands: commands,
	}, nil
}
