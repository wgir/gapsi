# 🚀 Gapsi Todo App - Frontend

A premium, modern Todo List application built with **React 18**, **Vite**, and **Tailwind CSS**. Designed for performance, scalability, and a superior user experience.

## ✨ Features

- **✅ Full Task CRUD**: Create, read, update, and delete tasks with ease.
- **🔄 Infinite Scroll**: Cursor-based pagination using `last_id` for seamless browsing of large task lists.
- **🌓 Dark Mode**: Full support for light and dark themes with system preference detection and persistence.
- **⚡ Real-time Updates**: Powered by **@tanstack/react-query** for efficient server state management and automatic cache invalidation.
- **📋 Type-Safe Forms**: Built with **React Hook Form** and **Zod** for robust client-side validation.
- **🎨 SaaS Aesthetics**: Modern UI with smooth animations powered by **Framer Motion** and a responsive layout.

## 🛠️ Tech Stack

- **Framework**: React 18 (TypeScript)
- **Build Tool**: Vite 8
- **Styling**: Tailwind CSS 3
- **API Client**: Axios
- **Server State**: React Query v5
- **Form Management**: React Hook Form
- **Validation**: Zod
- **Icons**: Lucide React
- **Animations**: Framer Motion

## 📂 Project Structure

```text
src/
├── api/                # Global Axios instance and interceptors
├── components/ui/      # Reusable UI components (Button, Input, Modal)
├── features/tasks/     # Task-specific feature logic
│   ├── api/            # Task API services
│   ├── components/     # Task-related React components
│   ├── hooks/          # Custom hooks (Infinite scroll, Mutations)
│   └── schemas/        # Zod schemas and TypeScript types
├── layouts/            # Page layouts (MainLayout)
├── pages/              # Application pages (Dashboard)
├── utils/              # Utility functions
└── App.tsx             # Main application entry
```

## 🚀 Getting Started

### Prerequisites

- **Node.js**: v18+ (v20 recommended)
- **npm**: v9+
- **Backend**: A running Gapsi Task API at `http://localhost:8080`

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```

### Development

Run the development server with HMR:
```bash
npm run dev
```

The application will be available at `http://localhost:5173`.

### Production Build

Build the optimized production bundle:
```bash
npm run build
```

Preview the production build locally:
```bash
npm run preview
```

## 🐳 Docker Deployment

You can run the application in a containerized environment using the provided configurations:

### Using Docker Compose
```bash
docker-compose up --build
```
The application will be served via Nginx at `http://localhost:3001`.

## ⚙️ Configuration

Environment variables can be set in a `.env` file at the root:
- `VITE_API_URL`: The URL of the backend API (Default: `http://localhost:8080`)

---
*Developed as part of the Gapsi technical evaluation.*
