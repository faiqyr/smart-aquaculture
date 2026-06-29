package repository

import (
	"smart-aquaculture-backend/domain"

	"gorm.io/gorm"
)

type MonitoringRepository interface {
	Save(monitoring *domain.PoolMonitoring) error
	GetLatest() (*domain.PoolMonitoring, error)
	GetHistory(limit int) ([]domain.PoolMonitoring, error)
}

type monitoringRepository struct {
	db *gorm.DB
}

func NewMonitoringRepository(db *gorm.DB) MonitoringRepository {
	return &monitoringRepository{db}
}

func (r *monitoringRepository) Save(monitoring *domain.PoolMonitoring) error {
	return r.db.Create(monitoring).Error
}

func (r *monitoringRepository) GetLatest() (*domain.PoolMonitoring, error) {
	var monitorings []domain.PoolMonitoring
	err := r.db.Order("created_at desc").Limit(1).Find(&monitorings).Error
	if err != nil {
		return nil, err
	}
	if len(monitorings) == 0 {
		return nil, nil // Jika kosong, kembalikan nil tanpa error
	}
	return &monitorings[0], nil
}

func (r *monitoringRepository) GetHistory(limit int) ([]domain.PoolMonitoring, error) {
	var history []domain.PoolMonitoring
	err := r.db.Order("created_at desc").Limit(limit).Find(&history).Error
	return history, err
}
