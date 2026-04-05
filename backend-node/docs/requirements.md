# Prompt: Desarrollo de API Gateway en Node.js para Microservicios Go

**Rol:** Actúa como un Desarrollador Backend Senior experto en Arquitectura de Microservicios, Node.js y comunicación entre servicios.

**Contexto:** Necesito crear un **API Gateway** en Node.js que actúe como punto de entrada único (Entry Point) para una aplicación de gestión de tareas. Este Gateway debe consumir un set de microservicios desarrollados en **Go**.

**Requerimientos Funcionales del Gateway:**
Debes exponer los siguientes endpoints públicos y orquestar la llamada al servicio de Go correspondiente:

1.  **Listar Tareas:** `GET /tasks` -> Llama a Go para obtener el array de tareas.
2.  **Listar Tareas paginadas:** `GET /tasks?limit=10` -> Llama a Go para obtener el array de tareas.
3. **Listar Tareas paginadas a partir de un id:** `GET /tasks?limit=10&last_id=abc123` -> Llama a Go para obtener el array de tareas a partir del id proporcionado.
4.  **Crear Tarea:** `POST /tasks` -> Recibe `titulo` y `descripcion`. Envía los datos a Go (quien genera el ID y timestamps).
5.  **Completar Tarea:** `PATCH /tasks/:id/complete` -> Notifica a Go para cambiar el estado.
6.  **Eliminar Tarea:** `DELETE /tasks/:id` -> Solicita la eliminación en el servicio de Go.

**Especificaciones del Modelo de Datos:**
Cada tarea manejada por el sistema tiene esta estructura (asegúrate de que el Gateway mapee correctamente estos campos):
* `id`: Definido por el sistema (backend).
* `titulo`: String.
* `descripcion`: String.
* `fecha_creacion`: String.
* `completada`: Booleano.

**Requerimientos Técnicos:**
* **Framework:** Node.js con Express.
* **Cliente HTTP:** Utiliza `axios` para la comunicación entre servicios.
* **Configuración:** Usa variables de entorno (dotenv) para la URL base del servicio de Go (`GO_SERVICE_URL`).
* **Buenas Prácticas:** * Implementa manejo de errores global.
    * Usa códigos de estado HTTP semánticos (200, 201, 404, 500).
    * Estructura el código de forma modular (rutas y controladores separados).
    * Incluye un middleware para el parseo de JSON y logs básicos de peticiones.

**Entregables:**
1.  Comando de instalación de dependencias.
2.  Estructura de archivos sugerida.
3.  Código completo y documentado del Gateway.
4.  Ejemplo de configuración del archivo `.env`.
5.  Archivo Readme.md con instrucciones de uso.
6.  Archivo Dockerfile para el despliegue del Gateway.
