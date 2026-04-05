# API Gateway - Gestión de Tareas

Este es el API Gateway desarrollado en **NestJS** que actúa como punto de entrada único para el sistema de gestión de tareas. Orquesta las peticiones hacia el microservicio de **Go**.

## Requerimientos

- Node.js v20+
- Docker (opcional)
- Microservicio de Go ejecutándose

## Configuración

1.  Copia el archivo `.env.example` a `.env`:
    ```bash
    cp .env.example .env
    ```
2.  Asegúrate de configurar la URL del servicio de Go en `GO_SERVICE_URL`.

## Instalación

```bash
$ npm install
```

## Ejecución

```bash
# Desarrollo
$ npm run start:dev

# Producción
$ npm run start:prod
```

## API Endpoints

| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `GET` | `/tasks` | Lista todas las tareas. |
| `GET` | `/tasks?limit=10&last_id=abc` | Lista tareas con paginación. |
| `POST` | `/tasks` | Crea una nueva tarea. |
| `PATCH` | `/tasks/:id/complete` | Marca una tarea como completada. |
| `DELETE` | `/tasks/:id` | Elimina una tarea. |

## Docker

Para construir y ejecutar la imagen:

```bash
docker build -t task-gateway .
docker run -p 8080:8080 --env-file .env task-gateway
```

## Estructura del Proyecto

- `src/tasks`: Módulo principal de tareas (Controladores, Servicios, DTOs).
- `src/common/filters`: Filtros globales de excepciones.
- `docs/requirements.md`: Especificaciones originales del proyecto.

---
Nest is [MIT licensed](https://github.com/nestjs/nest/blob/master/LICENSE).
