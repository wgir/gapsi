## 1. Objetivo
Diseñar y desarrollar una app web con react y vite,siguiendo las mejores prácticas de
desarrollo de software. El objetivo es evaluar habilidades técnicas para implementar una solución escalable,bien documentada y mantenible.

## 2. Context

**Tenpo** is a fintech company that provides personnel services to tenpistas (clients).  
To support this, Tenpo offers a platform that allows to manage their tenpistas allowing them to register their transactions.

The application must allow the registration and management of **tenpistas** and **transactions**.
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
│   │   ├── tenpistas/
│   │   ├── api/
│   │   │   └── tenpistas.api.ts
│   │   ├── components/
│   │   │   ├── TenpistaForm.tsx
│   │   │   ├── TenpistaList.tsx
│   │   │   └── TenpistaItem.tsx
│   │   ├── hooks/
│   │   │   └── useTenpistas.ts
│   │   ├── schemas/
│   │   │   └── tenpista.schema.ts
│   │   ├── types/
│   │   │   └── tenpista.types.ts
│   │   └── index.ts
│   │
│   └── transactions/
│       ├── api/
│       │   └── transactions.api.ts
│       ├── components/
│       │   ├── TransactionForm.tsx
│       │   ├── TransactionList.tsx
│       │   └── TransactionItem.tsx
│       ├── hooks/
│       │   └── useTransactions.ts
│       ├── schemas/
│       │   └── transaction.schema.ts
│       ├── types/
│       │   └── transaction.types.ts
│       └── index.ts
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
│  • Título: Dashboard Tenpista            │
│  • Acciones globales                     │
└──────────────────────────────────────────┘

┌───────────────┬──────────────────────────┐
│ Sidebar       │ Main Content              │
│               │                           │
│ • Transacciones│ • Tabla de Transacciones │
│ • Tenpistas     │ • Tabla de Tenpistas    │
│               │                           │
└───────────────┴───────────────────────────┘

---

## Flujo UX completo (muy importante)

### Flujo realista

Usuario entra → ve transacciones

Hace clic en “Nueva Transacción”

Selecciona tenpista

Si no existen: crea tenpista (modal)

vuelve automáticamente al form

Guarda transacción

Tabla se actualiza (React Query cache)

👉 Flujo fluido, sin recargar ni navegar

7️⃣ Estado visual y feedback
Estados obligatorios:

Loading (spinner o skeleton)

Error (mensaje claro)

Empty state (sin transacciones)

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

### Tenpistas

The application must allow:

- Create new tenpistas with the following fields:
  - `tenpista_name` (varchar)
  - `tenpista_rut` (varchar)
- Retrieve all tenpistas
- Retrieve an tenpista by ID
- Update an tenpista
- Delete an tenpista

---

### Transactions

The application must allow:

- Create new transactions with the following fields:
  - `transaction_amount` (int, in pesos)
  - `merchant_or_business` (varchar)
  - `tenpista_id` (int)
  - `transaction_date` (datetime)
- Retrieve all transactions
- Retrieve a transaction by ID
- Update a transaction
- Delete a transaction

---

### Constraints

- Transaction amounts **cannot be negative**.
- The transaction date **cannot be later than the current date and time**.
