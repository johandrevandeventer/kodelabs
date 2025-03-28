package kodelabs

import (
	"time"
)

// ConvertToInverterAPIv2Payload converts a Message to a payload that can be sent to the Kodelabs Inverter API v2
func ConvertToInverterAPIv2Payload(payload Message) []map[string]any {
	var points []map[string]any

	if payload.Timestamp.IsZero() {
		payload.Timestamp = time.Now()
	}

	if payload.Data["SerialNo1"] == nil {
		deviceIdentifier := payload.DeviceIdentifier
		deviceIdentifierPoint := map[string]any{
			"ts":         payload.Timestamp.Unix(),
			"registerNr": "SerialNo1",
			"value":      deviceIdentifier,
		}

		points = append(points, deviceIdentifierPoint)
	}

	for key, value := range payload.Data {
		point := map[string]any{
			"ts":         payload.Timestamp.Unix(),
			"registerNr": key,
			"value":      value,
		}

		points = append(points, point)
	}

	return points
}
