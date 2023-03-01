package queue

import "github.com/mrzalr/queue-api/internal/models"

type Repository interface {
	Find() ([]models.Queue, error)
	FindQueue() ([]models.Queue, error)
	Create(queue models.Queue) error
	FindLast() (models.Queue, error)
	GetCurrentQueueNumber() (int, error)
	Save(queue models.Queue) error
}

type Usecase interface {
	FindQueue() ([]models.Queue, error)
	Enqueue(queue models.Queue) ([]models.Queue, error)
	Dequeue() ([]models.Queue, error)
}
