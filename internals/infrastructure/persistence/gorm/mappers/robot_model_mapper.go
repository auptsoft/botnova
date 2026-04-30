package entity_mappers

import (
	"encoding/json"

	"auptex.com/botnova/internals/domain/models"
	gormentities "auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToRobotModelDomain(e gormentities.RobotModel) models.RobotModel {

	model := models.RobotModel{
		Id:        e.Id,
		ModelID:   e.ModelID,
		ModelName: e.ModelName,
		Version:   e.Version,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}

	json.Unmarshal([]byte(e.CommandsJson), &model.Commands)
	json.Unmarshal([]byte(e.PropertiesJson), &model.Properties)

	return model
}

func ToRobotModelEntity(m models.RobotModel) gormentities.RobotModel {
	commandsBytes, _ := json.Marshal(m.Commands)
	commandjson := string(commandsBytes)

	propertiesBytes, _ := json.Marshal(m.Properties)
	propertiesJson := string(propertiesBytes)

	return gormentities.RobotModel{
		Id:             m.Id,
		ModelID:        m.ModelID,
		ModelName:      m.ModelName,
		Version:        m.Version,
		CommandsJson:   commandjson,
		PropertiesJson: propertiesJson,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}
