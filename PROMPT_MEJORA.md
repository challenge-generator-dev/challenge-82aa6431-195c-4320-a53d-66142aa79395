# Prompt para Mejorar el Codigo Base

Copia y pega el contenido del bloque de abajo en un asistente de IA (Claude, ChatGPT)
para obtener el proyecto completo y arrancable. Guarda los archivos que te devuelva en la
carpeta del repo, respetando exactamente las rutas.

Si preferis trabajar en tu editor con un agente local (Claude Code, Cursor, Copilot), usa `AGENTS.md` en vez de este archivo: dice lo mismo pero para que escriba los archivos en disco.

## Las dos reglas que no se negocian

1. **Completa el boilerplate.** Todo lo que el proyecto necesita para compilar y arrancar: manifiesto de dependencias, punto de entrada, configuracion, capa de interfaz, y las capas del patron arquitectonico declarado. Eso es andamiaje y es tu trabajo.
2. **NO resuelvas el reto.** Los entregables de las fases son el trabajo de la persona. El hueco pedagogico se deja como esta: el proyecto arranca, pero lo que el reto pide implementar NO esta implementado.

Dicho de otra forma: si algo impide compilar, arreglalo. Si algo es logica de negocio incompleta, validaciones ausentes, un secreto hardcodeado o un patron mejorable, dejalo exactamente como esta — es lo que la persona tiene que encontrar.

## Superficie de practica — NO resuelvas

Estos archivos SON el ejercicio de la persona. No los implementes; deja stubs.

- `internal/domain/interfaces/http/router.go` — La fase "Configuración inicial y endpoint de creación de tareas" (f1) practica router.go (o equivalente) con la definición de rutas y middleware de logging aqui: este archivo es el ejercicio (blueprint).
- `internal/domain/test/task_service_test.go` — La fase "Revisión y optimización" (f4) practica task_service_test.go (o equivalente) con pruebas unitarias para los casos borde aqui: este archivo es el ejercicio (blueprint).
- `internal/domain/infrastructure/middleware/logging_middleware.go` — La fase "Configuración inicial y endpoint de creación de tareas" ya venia resuelta por el codigo base: este archivo pasa a ser el ejercicio.

## Como saber que terminaste

```bash
go mod tidy && go build ./...
```

Ese comando corriendo sin errores es la definicion de "listo".

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Perfil
Chapter Backend, Especialidad Desarrollador, Tecnología Go, Senior L2

### Brecha de conocimiento
Construir una API REST en Go con Gin para gestionar tareas

### Misión / candidato
Exponer la gestion de tareas por HTTP

### Reto
- Tema: API REST de tareas con Go y Gin
- Seniority: senior-l2
- Tipo: practical
- Título: Implementación de una API REST para gestión de tareas
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Configuración inicial y endpoint de creación de tareas — objetivo: Configurar el proyecto y crear el endpoint para añadir nuevas tareas. — entregable (NO resolver): Proyecto Go con Gin configurado, middleware de logging implementado y endpoint POST /tasks funcional.
- Fase 2: Endpoint para listar tareas — objetivo: Implementar el endpoint para listar todas las tareas. — entregable (NO resolver): Endpoint GET /tasks funcional que lista todas las tareas.
- Fase 3: Endpoint para actualizar tareas — objetivo: Implementar el endpoint para actualizar el estado de una tarea. — entregable (NO resolver): Endpoint PUT /tasks/{id} funcional que actualiza el estado de una tarea.
- Fase 4: Revisión y optimización — objetivo: Revisar y optimizar el código implementado. — entregable (NO resolver): Código revisado y optimizado.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (entrega de archivos).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Entrega de archivos
Entregá cada archivo con su ruta completa, respetando exactamente la estructura de rutas
indicada por los marcadores: un bloque por archivo, precedido por su ruta. No generes ZIPs
ni archivos adjuntos.
La entrega debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: Entregá los archivos completos al finalizar. No preguntes si el usuario
quiere recibirlos. Simplemente entregalos; No debes desplegar en el chat el resumen de lo que arreglaste, solo entregá los archivos.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del último archivo entregado
No preguntes si el usuario quiere los archivos: simplemente entregalos siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: go.mod ===
module github.com/user/project

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
)

// === ARCHIVO: main.go ===
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/app"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/repository"
	"github.com/user/project/internal/domain/interfaces/http"
)

func main() {
	router := gin.Default()

	// Initialize repository and service
	taskRepo := &repository.TaskRepository{}
	taskService := &application.TaskService{Repo: taskRepo}

	// Initialize routes with the service
	http.InitializeRoutes(router, taskService)

	// Initialize app routes (if any additional routes exist)
	app.InitializeRoutes(router)

	router.Run()
}

// === ARCHIVO: internal/app/routes.go ===
package app

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	// Sample route - will be replaced by domain code
	router.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sample endpoint",
		})
	})
}

// === ARCHIVO: internal/contracts/types.go ===
package contracts

// Sample contract types - interfaces and models go here

// === ARCHIVO: internal/domain/cmd/api/main.go ===
package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/middleware"
	"github.com/user/project/internal/domain/infrastructure/repository"
	"github.com/user/project/internal/domain/interfaces/http"
)

func main() {
	// Configurar el modo de Gin según el entorno
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializar el router con el middleware de logging
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware())

	// Inicializar el repositorio en memoria con concurrencia
	taskRepo := &repository.TaskRepository{
		Tasks: make(map[string]domain.Task),
		Mutex: sync.Mutex{},
	}

	// Inicializar el servicio de tareas
	taskService := &application.TaskService{
		Repo: taskRepo,
	}

	// Configurar las rutas HTTP
	http.InitializeRoutes(router, taskService)

	// Iniciar el servidor
	serverAddr := ":8080"
	log.Printf("Servidor Gin iniciado en %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

// === ARCHIVO: internal/domain/task.go ===
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

// === ARCHIVO: internal/domain/application/task_service.go ===
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

// === ARCHIVO: internal/domain/infrastructure/repository/task_repository.go ===
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

// === ARCHIVO: internal/domain/infrastructure/middleware/logging_middleware.go ===
// TODO(reto): archivo stubbeado; implementa los cuerpos.
package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
  panic("TODO(reto): import")
}

// === ARCHIVO: internal/domain/interfaces/http/router.go ===
// TODO(reto): archivo stubbeado; implementa los cuerpos.
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
)

func InitializeRoutes(router *gin.Engine, taskService *application.TaskService) {
	taskHandler := NewTaskHandler(taskService)

	// Mount routes according to the contract
	router.POST("/tasks", taskHandler.HandleCreateTask)
	router.GET("/tasks", taskHandler.HandleGetAllTasks)
	router.PUT("/tasks/:id", taskHandler.HandleUpdateTaskStatus)
}

// TaskHandler defines the HTTP handlers for task operations.
type TaskHandler struct {
	Service *application.TaskService
}

func NewTaskHandler(taskService *application.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: taskService,
	}
}

func (h *TaskHandler) HandleCreateTask(c *gin.Context) {
	panic("TODO(reto): HandleCreateTask")
}

func (h *TaskHandler) HandleGetAllTasks(c *gin.Context) {
	panic("TODO(reto): HandleGetAllTasks")
}

func (h *TaskHandler) HandleUpdateTaskStatus(c *gin.Context) {
	panic("TODO(reto): HandleUpdateTaskStatus")
}

// === ARCHIVO: internal/domain/interfaces/http/task_handler.go ===
// TODO(reto): archivo stubbeado; implementa los cuerpos.
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/application"
)

type TaskHandler struct {
	taskService *application.TaskService
}

func NewTaskHandler(taskService *application.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) HandleCreateTask(c *gin.Context) {
  panic("TODO(reto): HandleCreateTask")
}

func (h *TaskHandler) HandleGetAllTasks(c *gin.Context) {
  panic("TODO(reto): HandleGetAllTasks")
}

func (h *TaskHandler) HandleUpdateTaskStatus(c *gin.Context) {
  panic("TODO(reto): HandleUpdateTaskStatus")
}

// === ARCHIVO: internal/domain/test/task_service_test.go ===
// TODO(reto): archivo stubbeado; implementa las pruebas unitarias.
package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/repository"
)

func TestCreateTask_Success(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para creación exitosa de tarea")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	task, err := service.CreateTask("Nueva tarea de prueba")

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, "Nueva tarea de prueba", task.Title)
	assert.Equal(t, domain.TaskStatusPending, task.Status)
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para creación con título vacío")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	task, err := service.CreateTask("")

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.Equal(t, domain.ErrTitleEmpty, err)
}

func TestCreateTask_TitleTooLong(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para título que excede 100 caracteres")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	longTitle := string(make([]byte, 101))
	for i := range longTitle {
		longTitle = string(rune('a' + i%26))
	}

	task, err := service.CreateTask(longTitle)

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.Equal(t, domain.ErrTitleTooLong, err)
}

func TestGetAllTasks_Empty(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para listar tareas cuando no hay tareas")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	tasks, err := service.GetAllTasks()

	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.Empty(t, tasks)
}

func TestGetAllTasks_WithTasks(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para listar tareas existentes")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	_, _ = service.CreateTask("Tarea 1")
	_, _ = service.CreateTask("Tarea 2")

	tasks, err := service.GetAllTasks()

	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.Len(t, tasks, 2)
}

func TestUpdateTaskStatus_Success(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para actualización exitosa de estado")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	task, _ := service.CreateTask("Tarea para actualizar")
	updatedTask, err := service.UpdateTaskStatus(task.ID, domain.TaskStatusCompleted)

	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, domain.TaskStatusCompleted, updatedTask.Status)
}

func TestUpdateTaskStatus_InvalidID(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para actualización con ID inválido")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	updatedTask, err := service.UpdateTaskStatus("id-invalido", domain.TaskStatusCompleted)

	assert.Error(t, err)
	assert.Nil(t, updatedTask)
	assert.Equal(t, domain.ErrTaskNotFound, err)
}

func TestUpdateTaskStatus_InvalidStatus(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para actualización con estado inválido")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	task, _ := service.CreateTask("Tarea con estado inválido")
	invalidStatus := domain.TaskStatus("invalid_status")

	updatedTask, err := service.UpdateTaskStatus(task.ID, invalidStatus)

	assert.Error(t, err)
	assert.Nil(t, updatedTask)
	assert.Equal(t, domain.ErrInvalidTaskStatus, err)
}
```
