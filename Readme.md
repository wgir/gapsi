# 🚀 Gapsi Full Stack Todo System

This project is a comprehensive task management system, designed with scalability and performance in mind. It features a modern architectural stack consisting of a React frontend, a Node.js BFF (Backend for Frontend), and a robust Go microservice backend.

## 🏗️ Architecture Layers

The system is divided into three distinct layers:

1.  **Frontend (React)**: High-performance user interface built with React 18, Vite, and Tailwind CSS.
2.  **BFF / API Gateway (Node.js)**: NestJS-based gateway that handles orchestration and connects the frontend to the backend services.
3.  **Backend (Go)**: Go-based microservice following Clean Architecture, using Firestore as its primary data store.

## 🛠️ Technology Stack

### [Frontend](frontend/)
- **React 18** & TypeScript
- **Vite** for optimized building
- **Tailwind CSS** for styling
- **React Query** for server state management
- **Framer Motion** for animations

### [BFF - API Gateway](backend-node/)
- **Node.js** (v20+)
- **NestJS** framework
- **Axios** for microservice communication
- **RESTful API** endpoint orchestration

### [Backend - Task API](backend/)
- **Go 1.24+**
- **Chi Router** for high-performance routing
- **Google Cloud Firestore** (NoSQL Database)
- **Clean Architecture** (Domain, Application, Infrastructure layers)
- **Dockerized** environments

## 📂 Project Structure

```text
gapsi/
├── frontend/       # React Application
├── backend-node/  # Node.js BFF / API Gateway
└── backend/        # Go Microservice API
```

## 🚀 Getting Started

To run the entire system, each component can be started individually or using Docker if available in each directory.

### Running the Frontend
1. Navigate to `frontend/`
2. Run `npm install && npm run dev`

### Running the BFF (API Gateway)
1. Navigate to `backend-node/`
2. Run `npm install && npm run start:dev`

### Running the Backend (Go)
1. Navigate to `backend/`
2. Run `docker-compose up` (Recommended for local Firestore emulator support)

---
*Developed as part of the Gapsi technical evaluation.*