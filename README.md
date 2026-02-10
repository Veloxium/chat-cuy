# Chat Cuy

A real-time chat web application built with Go and React.

## Tech Stack

**Backend** — Go, Fiber, PostgreSQL, WebSocket, JWT

**Frontend** — React, TypeScript, Vite, Tailwind CSS, Zustand

## Features

- User authentication (register & login)
- Real-time private messaging via WebSocket
- Contact management (search & add users)
- Image stories with auto-slideshow
- Avatar upload with fallback
- Responsive layout (desktop & mobile)

## Getting Started

### Prerequisites

- Go 1.25+
- Node.js & npm
- Docker

### Backend

```bash
cd be
cp .env.example .env    # adjust values if needed
make db-up              # start PostgreSQL
make migrate-up         # run migrations
make run                # start the server (default: :3001)
```

### Frontend

```bash
cd fe
npm install
```

Create a `.env` file:

```
VITE_API_URL=http://localhost:3001
```

Then start the dev server:

```bash
npm run dev
```

## Project Structure

```
chat-cuy/
├── be/                     # Backend (Go)
│   ├── cmd/api/            # Entry point
│   ├── internal/
│   │   ├── config/         # Config & DB connection
│   │   ├── handler/        # HTTP handlers
│   │   ├── middleware/      # JWT auth middleware
│   │   ├── model/          # Data models
│   │   ├── repository/     # Database queries
│   │   ├── routes/         # Route definitions
│   │   ├── service/        # Business logic
│   │   └── websocket/      # WebSocket hub
│   ├── pkg/utils/          # Helper utilities
│   ├── db/migrations/      # SQL migrations
│   ├── docker-compose.yml
│   └── Makefile
├── fe/                     # Frontend (React)
│   ├── src/
│   │   ├── pages/          # Page components
│   │   ├── components/     # UI components
│   │   ├── store/          # Zustand stores
│   │   ├── hooks/          # Custom hooks
│   │   ├── utils/          # Axios, queries
│   │   └── schema/         # Zod validation
│   └── vite.config.ts
└── README.md
```

## Environment Variables

### Backend (`be/.env`)

| Variable | Default | Description |
|----------|---------|-------------|
| APP_PORT | 3001 | Server port |
| DB_HOST | localhost | PostgreSQL host |
| DB_PORT | 5432 | PostgreSQL port |
| DB_USER | chatcuy | DB username |
| DB_PASSWORD | chatcuy123 | DB password |
| DB_NAME | chat_cuy | Database name |
| JWT_SECRET | — | JWT signing key |

### Frontend (`fe/.env`)

| Variable | Description |
|----------|-------------|
| VITE_API_URL | Backend API URL (e.g. `http://localhost:3001`) |
