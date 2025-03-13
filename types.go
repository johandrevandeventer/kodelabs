package kodelabs

import "time"

type Message struct {
	State                  string
	Version                string
	CustomerName           string
	SiteName               string
	Gateway                string
	Controller             string
	DeviceType             string
	ControllerSerialNumber string
	DeviceName             string
	DeviceSerialNumber     string
	Data                   map[string]interface{}
	Timestamp              time.Time
}
