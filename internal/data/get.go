package data

import "job-service/internal/data/models"

func (b DataBase) GetAdv(title string) (models.Adv, error) {
	var advs models.Adv

	if err := b.Where(
		models.Adv{Title: title},
	).Find(&advs).Error; err != nil {
		return advs, err
	}

	return advs, nil
}

func (b DataBase) GetWorkers(title string) ([]models.Worker, error) {
	var workers []models.Worker

	if err := b.Where(
		models.Worker{Advtitle: title},
	).Find(&workers).Error; err != nil {
		return workers, err
	}

	return workers, nil
}
