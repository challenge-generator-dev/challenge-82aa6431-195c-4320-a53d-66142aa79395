package domain

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrTitleEmpty      = errors.New("el título de la tarea no puede estar vacío")
	ErrTitleTooLong    = errors.New("el título de la tarea no puede exceder 100 caracteres")
	ErrInvalidTaskID   = errors.New("el ID de la tarea es inválido")
	ErrTaskNotFound    = errors.New("tarea no encontrada")
	ErrInvalidTaskStatus = errors.New("estado de tarea inválido")
)

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusCompleted TaskStatus = "completed"
)

type Task struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewTask(id, title string) (*Task, error) {
	if strings.TrimSpace(title) == "" {
		return nil, ErrTitleEmpty
	}
	if len(title) > 100 {
		return nil, ErrTitleTooLong
	}

	return &Task{
		ID:        id,
		Title:     title,
		Status:    TaskStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (t *Task) UpdateStatus(newStatus TaskStatus) error {
	if newStatus != TaskStatusPending && newStatus != TaskStatusCompleted {
		return ErrInvalidTaskStatus
	}
	t.Status = newStatus
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return ErrTitleEmpty
	}
	if len(t.Title) > 100 {
		return ErrTitleTooLong
	}
	if t.Status != TaskStatusPending && t.Status != TaskStatusCompleted {
		return ErrInvalidTaskStatus
	}
	return nil
}