# AGENTS.md

Instrucciones para el agente de IA que abra este repositorio (Claude Code, Cursor, Codex, Copilot, Gemini). Se cargan solas: no hay que pegar nada en ningun chat.

## Que es este repositorio

Es el codigo base de un reto de aprendizaje de Pragma: **Implementación de una API REST para gestión de tareas**.

| | |
|---|---|
| Tema | API REST de tareas con Go y Gin |
| Nivel | senior-l2 |
| Chapter | Backend |
| Especialidad | Go + Gin |
| Stack | Go / Gin |
| Patron arquitectonico | capas estándar |
| Tiempo estimado | 8 horas |

## Receta del stack

Esqueleto obligatorio:

- `manifiesto de dependencias y build de Go en la raiz (segun las convenciones idiomaticas de Go/Gin)`
- `punto de entrada de la aplicacion (segun las convenciones idiomaticas de Go/Gin)`
- `capa de dominio con entidades y puertos (segun las convenciones idiomaticas de Go/Gin)`
- `capa de aplicacion con casos de uso (segun las convenciones idiomaticas de Go/Gin)`
- `capa de infraestructura con los adaptadores y los handlers HTTP de Gin (segun las convenciones idiomaticas de Go/Gin)`

Trampas conocidas:

- Todo el codigo, el manifiesto y los tests son de Go/Gin; no mezclar otro lenguaje ni otro runtime.

## Tu tarea

Dejar este proyecto en estado **verificable**: que el comando de verificacion corra sin errores. Escribi los archivos en disco, en este repositorio. No generes ZIPs ni archivos adjuntos.

En orden:

1. Corre `go mod tidy && go build ./...` y mira que falla.
2. Completa lo que falte de la lista de abajo: manifiesto de dependencias, punto de entrada, capa de interfaz y las capas del patron declarado.
3. Arregla SOLO los errores que impiden compilar o arrancar.
4. Volve a correr `go mod tidy && go build ./...` hasta que pase.
5. Pará ahí.

## Regla dura: las fases son trabajo del humano

**PROHIBIDO implementar los entregables de las fases.** El valor del reto esta en que la persona los resuelva. Tu trabajo es que tenga un proyecto que arranca; el hueco pedagogico se queda como esta.

No resuelvas nada de esto:

- **Fase 1 — Configuración inicial y endpoint de creación de tareas**: Proyecto Go con Gin configurado, middleware de logging implementado y endpoint POST /tasks funcional.
- **Fase 2 — Endpoint para listar tareas**: Endpoint GET /tasks funcional que lista todas las tareas.
- **Fase 3 — Endpoint para actualizar tareas**: Endpoint PUT /tasks/{id} funcional que actualiza el estado de una tarea.
- **Fase 4 — Revisión y optimización**: Código revisado y optimizado.

Distincion operativa:

- **Arreglar** (si): import faltante, tipo que no existe, dependencia sin declarar, error de sintaxis, archivo referenciado que no existe.
- **No tocar** (no): logica de negocio incompleta, validaciones ausentes, secretos hardcodeados, APIs deprecadas que funcionan, concurrencia insegura, patrones mejorables. Eso es lo que la persona tiene que encontrar.

## Superficie de practica (NO completes)

Estos archivos SON el ejercicio de la persona. No los implementes; deja stubs. No toques la logica que el reto pide completar.

- [ ] `internal/domain/interfaces/http/router.go` — La fase "Configuración inicial y endpoint de creación de tareas" (f1) practica router.go (o equivalente) con la definición de rutas y middleware de logging aqui: este archivo es el ejercicio (blueprint).
- [ ] `internal/domain/test/task_service_test.go` — La fase "Revisión y optimización" (f4) practica task_service_test.go (o equivalente) con pruebas unitarias para los casos borde aqui: este archivo es el ejercicio (blueprint).
- [ ] `internal/domain/infrastructure/middleware/logging_middleware.go` — La fase "Configuración inicial y endpoint de creación de tareas" ya venia resuelta por el codigo base: este archivo pasa a ser el ejercicio.

## Lo que falta y tenes que completar

No se detectaron huecos: estan los archivos declarados, el boilerplate del stack y ninguna referencia quedo colgando. Igual corre el comando de verificacion — que los archivos existan no garantiza que compilen.

### Presentes (12)

- `go.mod`
- `main.go`
- `internal/app/routes.go`
- `internal/contracts/types.go`
- `internal/domain/cmd/api/main.go`
- `internal/domain/task.go`
- `internal/domain/application/task_service.go`
- `internal/domain/infrastructure/repository/task_repository.go`
- `internal/domain/infrastructure/middleware/logging_middleware.go`
- `internal/domain/interfaces/http/router.go`
- `internal/domain/interfaces/http/task_handler.go`
- `internal/domain/test/task_service_test.go`

### Capas del patron declarado

Son las capas que la propuesta arquitectonica plantea, marcadas segun existan hoy en el repo. Orientan donde poner el codigo nuevo; no son una condicion para compilar.

- `cmd/api` (ausente)
- `internal/domain` (presente)
- `internal/application` (ausente)
- `internal/infrastructure/repository` (ausente)
- `internal/infrastructure/middleware` (ausente)
- `internal/interfaces/http` (ausente)
- `test` (ausente)

## Verificacion

```bash
go mod tidy && go build ./...
```

El comando tiene que pasar SIN implementar los archivos de la superficie de practica: solo andamiaje.

Ese comando pasando es la definicion de "terminado" para vos.

## Convenciones que tenes que respetar

- Un solo ecosistema: no declares librerias de otro lenguaje ni mezcles gestores de paquetes.
- Toda libreria que uses tiene que estar declarada en el manifiesto de dependencias.
- Todo import declarado tiene que usarse; todo tipo usado tiene que existir o venir de una dependencia declarada.
- El patron es **capas estándar**: los contratos (interfaces, puertos) los define la capa interna y los implementa la externa, nunca al revés.
- Los archivos que crees llevan implementacion real, no stubs: sin `TODO`, sin cuerpos vacios, sin `// getters y setters`.

## Contexto del candidato

Sirve para calibrar el nivel del codigo, no para resolver las fases.

- Perfil: Chapter Backend, Especialidad Desarrollador, Tecnología Go, Senior L2
- Brecha que el reto ataca: Construir una API REST en Go con Gin para gestionar tareas
- Mision: Exponer la gestion de tareas por HTTP

---

*Generado por Challenge Generator — Pragma. `README.md` tiene el enunciado completo del reto para la persona. `PROMPT_MEJORA.md` es la variante para pegar en un chat, si se prefiere ese flujo.*
