package mysql

import (
	"github.com/mrzalr/queue-api/internal/models"
	"github.com/mrzalr/queue-api/internal/queue"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) queue.Repository {
	return &repository{db}
}

func (r *repository) Find() ([]models.Queue, error) {
	queue := []models.Queue{}
	tx := r.db.Find(&queue)
	return queue, tx.Error
}

// this function will return queue that not done yet
func (r *repository) FindQueue() ([]models.Queue, error) {
	queue := []models.Queue{}
	tx := r.db.Preload("Patient").Where("is_done = ?", false).Find(&queue)
	return queue, tx.Error
}

func (r *repository) Create(queue models.Queue) error {
	tx := r.db.Create(&queue)
	return tx.Error
}

func (r *repository) GetCurrentQueueNumber() (int, error) {
	queue := models.Queue{}
	tx := r.db.Order("number DESC").Last(&queue)
	return queue.Number, tx.Error
}

func (r *repository) FindLast() (models.Queue, error) {
	queue := models.Queue{}
	tx := r.db.Where("is_done = ?", false).Order("created_at").First(&queue)
	return queue, tx.Error
}

func (r *repository) Save(queue models.Queue) error {
	tx := r.db.Where("number = ?", queue.Number).Updates(&queue)
	return tx.Error
}
