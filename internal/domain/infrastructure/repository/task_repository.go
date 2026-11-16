package repository

import (
	"errors"
	"sync"

	"github.com/user/project/internal/domain"
	"TODO(reto)"
)

type TaskRepository struct {
	Tasks map[string]domain.Task
	Mutex sync.Mutex
}

func (r *TaskRepository) Save(task *domain.Task) error {
	// TODO(reto): Implementar la lógica para guardar una tarea
	// - Validar que la tarea no sea nil
	// - Validar la tarea usando task.Validate()
	// - Usar Mutex para manejar concurrencia
	// - Guardar la tarea en el mapa
	// - Retornar nil si todo es correcto o un error
	return nil
}

func (r *TaskRepository) FindByID(id string) (*domain.Task, error) {
	// TODO(reto): Implementar la lógica para buscar una tarea por ID
	// - Validar que el ID no esté vacío
	// - Usar Mutex para manejar concurrencia
	// - Buscar la tarea en el mapa
	// - Retornar la tarea si existe o domain.ErrTaskNotFound si no existe
	return nil, nil
}

func (r *TaskRepository) FindAll() ([]*domain.Task, error) {
	// TODO(reto): Implementar la lógica para listar todas las tareas
	// - Usar Mutex para manejar concurrencia
	// - Crear una slice para almacenar las tareas
	// - Iterar sobre el mapa y agregar cada tarea a la slice
	// - Retornar la slice de tareas o un error
	return nil, nil
}

func (r *TaskRepository) Update(task *domain.Task) error {
	// TODO(reto): Implementar la lógica para actualizar una tarea
	// - Validar que la tarea no sea nil
	// - Validar la tarea usando task.Validate()
	// - Usar Mutex para manejar concurrencia
	// - Verificar que la tarea existe en el mapa
	// - Actualizar la tarea en el mapa
	// - Retornar nil si todo es correcto o un error
	return nil
}