package kodelabs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var sharedClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     30 * time.Second,
		MaxConnsPerHost:     10,
		MaxIdleConnsPerHost: 10,
	},
}

// SendPostRequest sends a POST request to the Kodelabs API
func SendPostRequest(message Message, payload any, url, authToken string, logger *zap.Logger) ([]byte, error) {
	var response Response

	customer := message.CustomerName
	siteName := message.SiteName
	gateway := message.Gateway
	controller := message.Controller
	deviceType := message.DeviceType
	controllerSerialNumber := message.ControllerSerialNumber
	deviceName := message.DeviceName
	deviceSerialNumber := message.DeviceSerialNumber

	logger.Info(fmt.Sprintf("Posting to Kodelabs -> %s :: %s :: %s :: %s :: %s :: %s :: %s :: %s", gateway, customer, siteName, controller, controllerSerialNumber, deviceType, deviceSerialNumber, deviceName))
	logger.Debug(fmt.Sprintf("Posting to Kodelabs -> %s :: %s :: %s :: %s :: %s :: %s :: %s :: %s", gateway, customer, siteName, controller, controllerSerialNumber, deviceType, deviceSerialNumber, deviceName), zap.String("url", url), zap.Any("data", payload))

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	startTime := time.Now()

	resp, err := sharedClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP error: Status Code %d, Body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	elapsedTime := time.Since(startTime)

	logger.Info(fmt.Sprintf("Response -> %s :: %s :: %s :: %s :: %s :: %s :: %s :: %s", gateway, customer, siteName, controller, controllerSerialNumber, deviceType, deviceSerialNumber, deviceName), zap.String("status", response.Status))
	logger.Debug(fmt.Sprintf("Response -> %s :: %s :: %s :: %s :: %s :: %s :: %s :: %s", gateway, customer, siteName, controller, controllerSerialNumber, deviceType, deviceSerialNumber, deviceName), zap.String("status", response.Status), zap.String("duration", elapsedTime.String()), zap.String("message", response.Message))

	return nil, nil
}
