package usecase

import (
	"github.com/mrzalr/queue-api/internal/models"
	"github.com/mrzalr/queue-api/internal/queue"
)

type usecase struct {
	repository queue.Repository
}

func New(repository queue.Repository) queue.Usecase {
	return &usecase{repository}
}

func (u *usecase) FindQueue() ([]models.Queue, error) {
	return u.repository.FindQueue()
}

func (u *usecase) Enqueue(queue models.Queue) ([]models.Queue, error) {
	err := u.repository.Create(queue)
	if err != nil {
		return []models.Queue{}, err
	}

	return u.repository.FindQueue()
}

func (u *usecase) Dequeue() ([]models.Queue, error) {
	last, err := u.repository.FindLast()
	if err != nil {
		return []models.Queue{}, err
	}

	last.IsDone = true
	err = u.repository.Save(last)
	if err != nil {
		return []models.Queue{}, err
	}

	return u.repository.FindQueue()
}
