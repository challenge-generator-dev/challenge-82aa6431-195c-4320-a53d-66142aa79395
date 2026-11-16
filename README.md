# Implementación de una API REST para gestión de tareas

El sistema debe exponer una API REST en Go con Gin para gestionar tareas. Cada tarea tiene un título (máximo 100 caracteres) y un estado. El repositorio es en memoria y no se utiliza GORM. Se debe implementar un middleware propio de logging.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | API REST de tareas con Go y Gin |
| **Nivel** | senior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y entregará cada archivo con su ruta.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y entregará cada archivo con su ruta. Guárdalos en la carpeta donde quieras trabajar, respetando las rutas.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no entrega todos los archivos al terminar el análisis, escríbele: "entrégame todos los archivos con su ruta".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Configuración inicial y endpoint de creación de tareas

**Objetivo:** Configurar el proyecto y crear el endpoint para añadir nuevas tareas.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Configura el proyecto en Go con Gin.
- Implementa un middleware de logging personalizado.
- Crea el endpoint POST /tasks para añadir nuevas tareas. Valida que el título no esté vacío y tenga máximo 100 caracteres.

**Entregable:** Proyecto Go con Gin configurado, middleware de logging implementado y endpoint POST /tasks funcional.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda que el repositorio es en memoria.
- El middleware de logging debe registrar cada petición.

</details>

### Fase 2: Endpoint para listar tareas

**Objetivo:** Implementar el endpoint para listar todas las tareas.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Crea el endpoint GET /tasks para listar todas las tareas almacenadas en el repositorio en memoria.

**Entregable:** Endpoint GET /tasks funcional que lista todas las tareas.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda que el repositorio es en memoria.

</details>

### Fase 3: Endpoint para actualizar tareas

**Objetivo:** Implementar el endpoint para actualizar el estado de una tarea.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Crea el endpoint PUT /tasks/{id} para actualizar el estado de una tarea específica.

**Entregable:** Endpoint PUT /tasks/{id} funcional que actualiza el estado de una tarea.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda que el repositorio es en memoria.

</details>

### Fase 4: Revisión y optimización

**Objetivo:** Revisar y optimizar el código implementado.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Revisa el código implementado en las fases anteriores.
- Identifica y corrige posibles errores o ineficiencias.
- Optimiza el código para mejorar la legibilidad y el rendimiento.

**Entregable:** Código revisado y optimizado.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda revisar las validaciones y el manejo de errores.
- Considera la legibilidad y el rendimiento del código.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es una API REST y cómo se implementa en Go con Gin?
- **paraQueSirve**: ¿Para qué sirve el middleware de logging en una API REST?
- **comoSeUsa**: ¿Cómo se usa Gin para crear endpoints en una API REST?
- **erroresComunes**: ¿Cuáles son los errores comunes al implementar una API REST en Go con Gin?
- **queDecisionesImplica**: ¿Qué decisiones implica la implementación de una API REST en Go con Gin?

## Criterios de Evaluacion

- Configuración inicial del proyecto en Go con Gin.
- Implementación de un middleware de logging personalizado.
- Creación del endpoint POST /tasks para añadir nuevas tareas con validaciones de título.
- Creación del endpoint GET /tasks para listar todas las tareas.
- Creación del endpoint PUT /tasks/{id} para actualizar el estado de una tarea.
- Revisión y optimización del código implementado.

## Como trabajar con un asistente de IA

Hay dos caminos, elegi uno:

- **AGENTS.md** (recomendado) — instrucciones nativas del repo. Abri esta carpeta con tu agente local (Claude Code, Cursor, Codex, Copilot, Gemini) y las carga solo. Sabe que archivos faltan y con que comando se verifica, y completa el scaffold escribiendo en disco.
- **PROMPT_MEJORA.md** — para copiar y pegar en un chat (claude.ai, ChatGPT) si no tenes un agente en el IDE. Declara lo mismo que AGENTS.md; los archivos que te devuelva el chat los guardas en esta carpeta respetando las rutas.

Ninguno de los dos resuelve las fases del reto: eso es tu trabajo.

## Verificacion

El proyecto esta listo para trabajar cuando este comando corre sin errores:

```bash
go mod tidy && go build ./...
```

---

*Reto generado automaticamente por Challenge Generator - Pragma*
