package domainservices

import (
	"strconv"

	"auptex.com/botnova/internals/domain/models"
)

func ApplyCalibrationToCommand(calibrationProfile models.CalibrationProfile, command models.Command) models.Command {
	cal, ok := calibrationProfile.Commands[command.Name]
	if !ok {
		return command
	}

	// Apply calibration to parameters
	for key, value := range command.Params {
		if calParam, ok := cal.Parameters[key]; ok {
			switch v := value.(type) {
			case float64:
				calibratedValue := v
				if calParam.Scale != nil {
					calibratedValue = calibratedValue * *calParam.Scale
				}

				if calParam.Offset != nil {
					calibratedValue = calibratedValue + *calParam.Offset
				}

				if calParam.Min != nil && calibratedValue < *calParam.Min {
					calibratedValue = *calParam.Min
				}

				if calParam.Max != nil && calibratedValue > *calParam.Max {
					calibratedValue = *calParam.Max
				}
				command.Params[key] = calibratedValue

				//If there's an append or prepend string, we convert the calibrated value to a string and add the append/prepend
				if calParam.Append != nil {
					command.Params[key] = strconv.FormatFloat(calibratedValue, 'f', 2, 64) + *calParam.Append
				}
				if calParam.Prepend != nil {
					command.Params[key] = *calParam.Prepend + strconv.FormatFloat(calibratedValue, 'f', 2, 64)
				}
			case int:
				calibratedValue := float64(v)
				if calParam.Scale != nil {
					calibratedValue = calibratedValue * *calParam.Scale
				}

				if calParam.Offset != nil {
					calibratedValue = calibratedValue + *calParam.Offset
				}

				if calParam.Min != nil && calibratedValue < *calParam.Min {
					calibratedValue = *calParam.Min
				}

				if calParam.Max != nil && calibratedValue > *calParam.Max {
					calibratedValue = *calParam.Max
				}
				command.Params[key] = int(calibratedValue)

				//If there's an append or prepend string, we convert the calibrated value to a string and add the append/prepend
				if calParam.Append != nil {
					command.Params[key] = strconv.Itoa(int(calibratedValue)) + *calParam.Append
				}
				if calParam.Prepend != nil {
					command.Params[key] = *calParam.Prepend + strconv.Itoa(int(calibratedValue))
				}

			case string:
				// Handle string parameters with prepend/append
				if calParam.Append != nil {
					command.Params[key] = command.Params[key].(string) + *calParam.Append
				}
				if calParam.Prepend != nil {
					command.Params[key] = *calParam.Prepend + command.Params[key].(string)
				}
			default:
				// For non-numeric parameters, we could apply other types of calibration if needed
				// For now, we'll just leave them unchanged
			}
		}
	}

	return command
}
