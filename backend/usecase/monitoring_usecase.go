package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"smart-aquaculture-backend/domain"
	"smart-aquaculture-backend/infrastructure"
	"smart-aquaculture-backend/repository"

	"firebase.google.com/go/v4/messaging"
)

type MonitoringUsecase interface {
	SaveTelemetryData(temp float32, humidity float32, light int) error
	GetLatestMonitoring() (*domain.PoolMonitoring, error)
	GetMonitoringHistory(limit int) ([]domain.PoolMonitoring, error)
}

type monitoringUsecase struct {
	monitoringRepo repository.MonitoringRepository
	tokenRepo      repository.TokenRepository
}

func NewMonitoringUsecase(repo repository.MonitoringRepository, tokenRepo repository.TokenRepository) MonitoringUsecase {
	return &monitoringUsecase{
		monitoringRepo: repo,
		tokenRepo:      tokenRepo,
	}
}

func (u *monitoringUsecase) SaveTelemetryData(temp float32, humidity float32, light int) error {
	data := &domain.PoolMonitoring{
		Temperature:    temp,
		Humidity:       humidity,
		LightIntensity: light,
	}

	err := u.monitoringRepo.Save(data)
	if err != nil {
		log.Printf("Error saving telemetry data: %v", err)
		return err
	}

	log.Printf("Telemetry data saved: Temp=%.2f, Humidity=%.2f, Light=%d", temp, humidity, light)

	// Telegram Alert Logic
	if temp > 35.0 {
		go sendTelegramAlert(temp)
		go u.sendFirebaseAlert(temp)
	}

	return nil
}

func sendTelegramAlert(temp float32) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		log.Println("Telegram credentials not configured, skipping alert")
		return
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	message := fmt.Sprintf("⚠️ PERINGATAN: Suhu air kolam terlalu panas! %.2f°C", temp)

	payload := map[string]string{
		"chat_id": chatID,
		"text":    message,
	}
	
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Failed to send Telegram alert: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Telegram alert sent successfully!")
	} else {
		log.Printf("Failed to send Telegram alert, status code: %d", resp.StatusCode)
	}
}

func (u *monitoringUsecase) sendFirebaseAlert(temp float32) {
	if infrastructure.FirebaseApp == nil {
		return
	}
	client, err := infrastructure.FirebaseApp.Messaging(context.Background())
	if err != nil {
		log.Printf("Error getting Messaging client: %v\n", err)
		return
	}

	tokens, err := u.tokenRepo.GetAllTokens()
	if err != nil || len(tokens) == 0 {
		return
	}

	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: "⚠️ Peringatan Suhu Panas!",
			Body:  fmt.Sprintf("Suhu kolam saat ini: %.2f°C", temp),
		},
	}

	br, err := client.SendMulticast(context.Background(), message)
	if err != nil {
		log.Printf("Gagal mengirim Firebase Push: %v", err)
		return
	}
	log.Printf("Firebase message sent. Berhasil: %d, Gagal: %d\n", br.SuccessCount, br.FailureCount)
}

func (u *monitoringUsecase) GetLatestMonitoring() (*domain.PoolMonitoring, error) {
	return u.monitoringRepo.GetLatest()
}

func (u *monitoringUsecase) GetMonitoringHistory(limit int) ([]domain.PoolMonitoring, error) {
	return u.monitoringRepo.GetHistory(limit)
}
