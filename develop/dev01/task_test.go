package dev01

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockNTPClient struct {
	mock.Mock
}

func (m *MockNTPClient) Time(server string) (time.Time, error) {
	args := m.Called(server)
	return args.Get(0).(time.Time), args.Error(1)
}

func TestSetNTPServer(t *testing.T) {
	originalServer := mtime
	defer func() { mtime = originalServer }() // Восстановление исходного значения после теста

	testServer := "1.test-ntp-server.org"
	SetNTPServer(testServer)
	if mtime != testServer {
		t.Errorf("Expected NTP server to be '%s', but got '%s'", testServer, mtime)
	}
}

func TestGetTime(t *testing.T) {
	mockNTP := new(MockNTPClient)
	expectedTime := time.Now().Truncate(time.Second) // Округляем до секунд
	mockNTP.On("Time", mock.Anything).Return(expectedTime, nil)

	// Подмена реализации ntp.Time на мок-объект

	actualTime, err := GetTime()
	if err != nil {
		t.Fatal("Failed to get time:", err)
	}

	if !actualTime.Truncate(time.Second).Equal(expectedTime) { // Округляем до секунд при сравнении
		t.Errorf("Expected time '%v', but got '%v'", expectedTime, actualTime)
	}
}
