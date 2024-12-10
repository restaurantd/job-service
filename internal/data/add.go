package data

import "job-service/internal/data/models"

func (b DataBase) Add(w models.Worker) error {
	if err := b.Create(&w).Error; err != nil {
		return err
	}

	return nil
}
