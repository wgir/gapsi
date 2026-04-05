## 1. Objetivo
Diseñar y desarrollar una app web con react y vite,siguiendo las mejores prácticas de
desarrollo de software. El objetivo es evaluar habilidades técnicas para implementar una solución escalable,bien documentada y mantenible.

La aplicacion es una todo-list permitiendo gestionar tareas. Creacion, listado, edicion, eliminacion y cambio de estado de las tareas.

---

## 🛠️ Tech Stack
- **React 18** 
- **Vite** 
- **Tailwind CSS**
- **Axios** (fetching obligatorio)
- **@tanstack/react-query** (plus)
- **React Hook Form** (validaciones)
- **Zod** (schema validation)
- **Docker**
- **docker-compose**

---

## Estructura de carpetas

src/
├── api/
│   ├── axiosInstance.ts
│   └── index.ts
│
├── features/
│   │   ├── tasks/
│   │   ├── api/
│   │   │   └── tasks.api.ts
│   │   ├── components/
│   │   │   ├── TaskForm.tsx
│   │   │   ├── TaskList.tsx
│   │   │   └── TaskItem.tsx
│   │   ├── hooks/
│   │   │   └── useTasks.ts
│   │   ├── schemas/
│   │   │   └── task.schema.ts
│   │   ├── types/
│   │   │   └── tenpista.types.ts
│   │   └── index.ts
│   │
│
├── components/
│   └── ui/
│       ├── Button.tsx
│       ├── Input.tsx
│       ├── Select.tsx
│       └── Modal.tsx
│
├── pages/
│   └── Dashboard.tsx
│
├── layouts/
│   └── MainLayout.tsx
│
├── utils/
│   └── date.ts
│
├── App.tsx
├── main.tsx
└── index.css


### Layout
Estructura general del Dashboard
┌──────────────────────────────────────────┐
│ Header                                   │
│  • Título: Todo List            │
│  • Acciones globales                     │
└──────────────────────────────────────────┘

┌───────────────┬──────────────────────────┐
│ Sidebar       │ Main Content              │
│               │                           │
│ • Tasks       │ • Tabla de Tasks          │
│               │                           │
└───────────────┴───────────────────────────┘

---

## Flujo UX completo (muy importante)

### Flujo realista

Usuario entra → ve tareas (titulo, descripcion, estado) en scroll infinito con paginacion de 10 tareas por pagina ordenadas por fecha de creacion descendente -> ve boton de nueva tarea

Hace clic en “Nueva Tarea”, se abre un modal con un formulario para crear una tarea

Digita el titulo de la tarea
Digita la descripcion de la tarea

Guarda la tarea

Tabla se actualiza (React Query cache)

👉 Flujo fluido, sin recargar ni navegar

7️⃣ Estado visual y feedback
Estados obligatorios:

Loading (spinner o skeleton)

Error (mensaje claro)

Empty state (sin tareas)

---

## 📦 Deliverables
- Folder structure
- Layout files
- Reusable UI components
- Example pages
- Minimal but clean UI

---

## ✨ Optional Enhancements
- SaaS-style UI (spacing, shadows, typography)
- Dark mode support
- Route protection via middleware
- State management for authentication


---

### Tasks

The application must allow:

- Create new tasks with the following fields:
  - `title` (varchar)
  - `description` (varchar)
- Retrieve all tasks
- Retrieve a task by ID
- Update a task
- Delete an task

---

### Constraints

- Task title **cannot be empty**.
- Task description **cannot be empty**.
