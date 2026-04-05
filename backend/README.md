# Gapsi Todo Backend

Una API REST de lista de tareas estable y escalable construida con **Go (Golang)** y **Google Cloud Firestore**, siguiendo los principios de **Clean Architecture**.

## 🚀 Características
- **Arquitectura Limpia**: Separación clara de responsabilidades (Domain, Application, Infrastructure).
- **Base de Datos NoSQL**: Integración con Google Firestore.
- **Contenerización**: Configuración completa con Docker y Docker Compose.
- **Resiliencia**: Manejo de apagado gradual (Graceful Shutdown) y logging estructurado con Zap.
- **Configuración Dinámica**: Gestión de entornos mediante Viper.

## 🛠️ Tecnologías
- **Lenguaje**: Go 1.24+
- **Enrutador**: [Chi](https://github.com/go-chi/chi)
- **Base de Datos**: [Google Cloud Firestore](https://cloud.google.com/firestore)
- **Logger**: [Zap](https://github.com/uber-go/zap)
- **Configuración**: [Viper](https://github.com/spf13/viper)

## 🏗️ Estructura del Proyecto
```text
.
├── cmd/api/          # Punto de entrada (main.go)
├── internal/
│   ├── domain/       # Modelos e interfaces de negocio
│   ├── application/  # Servicios y lógica de aplicación
│   └── infrastructure/
│       ├── db/       # Implementación de Firestore
│       ├── web/      # Handlers HTTP y Routas
│       ├── config/   # Carga de configuración
│       └── logger/   # Configuración de logs
├── docs/             # Documentación, bitácora y guías
├── Dockerfile        # Definición de la imagen de la API
└── docker-compose.yaml # Orquestación de API + Emulador Firestore
```

## 🏃 Cómo Ejecutar

### Requisitos Previos
- Docker y Docker Compose instalados.

### Opción 1: Con Docker Compose (Recomendado)
Este método levanta automáticamente la API y un emulador de Firestore para desarrollo local.

```bash
docker-compose up --build
```
- La API estará disponible en: `http://localhost:8080`
- El emulador de Firestore en el puerto: `8081`

### Opción 2: Ejecución Local (Sin Docker)
1. Instala las dependencias:
   ```bash
   go mod tidy
   ```
2. Arrancar el emulador de Firestore:
   ```bash
   firebase emulators:start --only firestore
   ```

3. Configura las variables de entorno (puedes crear un archivo `.env`):
   ```env
   APP_PORT=8090
   PROJECT_ID=todo-project
   FIRESTORE_EMULATOR_HOST=localhost:8080
   ```
3. Ejecuta la aplicación:
   ```bash
   go run cmd/api/main.go
   ```

## 📡 API Endpoints

| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `GET` | `/health` | Estado de la aplicación |
| `POST` | `/tasks` | Crear una nueva tarea |
| `GET` | `/tasks` | Listar todas las tareas (soporta query `?status=`, `?limit=10`, `?last_id=`) |
| `GET` | `/tasks/{id}` | Obtener una tarea por ID |
| `PUT` | `/tasks/{id}` | Actualizar una tarea existente |
| `DELETE` | `/tasks/{id}` | Eliminar una tarea |

### Ejemplo de Uso (cURL)

**Crear Tarea:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Mi Primera Tarea", "description": "Implementar Clean Architecture"}'
```

**Listar Tareas:**
```bash
curl http://localhost:8080/tasks
```

## 🧪 Pruebas
Ejecuta las pruebas unitarias con:
```bash
go test ./...
```

## 📝 Documentación Adicional
- [Bitácora de Desarrollo](docs/bitacora.md)
- [Guía del Emulador](docs/firestore-readme.md)
- [Requerimientos](docs/requirements.md)
