package usecase

import (
	"fmt"

	"github.com/mrzalr/queue-api/internal/models"
	"github.com/mrzalr/queue-api/internal/queue"
	"gorm.io/gorm"
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
	lastQueueNumber, err := u.repository.GetCurrentQueueNumber()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return []models.Queue{}, err
		}
	}

	queue.Number = lastQueueNumber + 1
	err = u.repository.Create(queue)
	if err != nil {
		return []models.Queue{}, err
	}

	return u.repository.FindQueue()
}

func (u *usecase) Dequeue() ([]models.Queue, error) {
	last, err := u.repository.FindLast()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []models.Queue{}, fmt.Errorf("queue is empty")
		} else {
			return []models.Queue{}, err
		}
	}

	last.IsDone = true
	err = u.repository.Save(last)
	if err != nil {
		return []models.Queue{}, err
	}

	return u.repository.FindQueue()
}
