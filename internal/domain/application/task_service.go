package application

import (
	"github.com/user/project/internal/domain"
	"TODO(reto)"
)

type TaskService struct {
	Repo TaskRepository
}

type TaskRepository interface {
	Save(task *domain.Task) error
	FindByID(id string) (*domain.Task, error)
	FindAll() ([]*domain.Task, error)
	Update(task *domain.Task) error
}

func (s *TaskService) CreateTask(title string) (*domain.Task, error) {
	// TODO(reto): Implementar la lógica para crear una tarea
	// - Generar un ID único para la tarea
	// - Validar el título usando domain.NewTask
	// - Guardar la tarea usando el repositorio
	// - Retornar la tarea creada o un error
	return nil, nil
}

func (s *TaskService) GetAllTasks() ([]*domain.Task, error) {
	// TODO(reto): Implementar la lógica para listar todas las tareas
	// - Usar el repositorio para obtener todas las tareas
	// - Retornar la lista de tareas o un error
	return nil, nil
}

func (s *TaskService) UpdateTaskStatus(id string, newStatus domain.TaskStatus) (*domain.Task, error) {
	// TODO(reto): Implementar la lógica para actualizar el estado de una tarea
	// - Validar el ID de la tarea
	// - Obtener la tarea usando el repositorio
	// - Validar que la tarea existe
	// - Actualizar el estado usando domain.Task.UpdateStatus
	// - Guardar los cambios usando el repositorio
	// - Retornar la tarea actualizada o un error
	return nil, nil
}