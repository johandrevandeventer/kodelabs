package kodelabs

import "time"

type Message struct {
	State                string
	Version              string
	CustomerName         string
	SiteName             string
	Controller           string
	DeviceType           string
	ControllerIdentifier string
	DeviceName           string
	DeviceIdentifier     string
	Data                 map[string]interface{}
	Timestamp            time.Time
}
