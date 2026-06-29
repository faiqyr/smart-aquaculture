package domain

import (
	"time"
)

// PoolMonitoring represents the sensor data log for the aquaculture pool
type PoolMonitoring struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Temperature    float32   `gorm:"not null"`
	Humidity       float32   `gorm:"not null"`
	LightIntensity int       `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

// DeviceControl represents the current status of the devices
type DeviceControl struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	IsFeeding bool      `gorm:"not null;default:false"`
	LedStatus bool      `gorm:"not null;default:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// DeviceToken stores FCM tokens for push notifications
type DeviceToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
