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
		deviceSerialNumber := payload.DeviceSerialNumber
		deviceSerialNumberPoint := map[string]any{
			"ts":         payload.Timestamp.Unix(),
			"registerNr": "SerialNo1",
			"value":      deviceSerialNumber,
		}

		points = append(points, deviceSerialNumberPoint)
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
