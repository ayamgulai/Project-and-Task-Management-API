 
# Project and Task Management API

Mini Jira-like backend API built with Go for managing projects, tasks, and users.

**Repository:** mini-jira-backend

## Overview

This repository implements a RESTful API with JWT authentication, role-based access control, project and task management, and Swagger/OpenAPI documentation.

## Quick Features

- User registration & login (JWT)
- Role-based access (admin, user)
- Create/read/update projects and tasks
- Assign tasks and track status/priority
- Task activity logs
- Auto-generated Swagger docs

## Tech Stack

- Go 1.25.5
- Gin Web Framework
- PostgreSQL
- JWT for authentication
- Swagger (swaggo) for API docs

## Quick Start

1. Clone the repo:

```bash
git clone https://github.com/ayamgulai/Project-and-Task-Management-API.git
cd Project-and-Task-Management-API
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` in the project root (example):

```env
DATABASE_URL=postgres://username:password@localhost:5432/project_management
JWT_SECRET=your_secret_key_here
PORT=8080
```

4. Ensure your PostgreSQL instance and database exist. The app runs migrations from `migrations/` on startup.

5. Run the server:

```bash
go run main.go
```

The API will be available at `http://localhost:8080`.

## Swagger / API Docs

Generate or update docs locally:

```bash
swag init
```

Then open the docs at:

http://localhost:8080/swagger/index.html

You can also see the generated OpenAPI files in the `docs/` folder (`swagger.json`, `swagger.yaml`).

## Database Migrations

- Migration SQL files are in `migrations/` (e.g. `migrations/001_init.sql`).
- The app's migration logic is implemented in `configs/migration.go` and runs at startup.

## Project Structure

```
.
├── main.go
├── go.mod
├── configs/        # DB and migration setup
├── controllers/    # HTTP handlers
├── models/         # Data models
├── repositories/   # DB access
├── services/       # Business logic
├── routes/         # Router definitions
├── middlewares/    # Auth middleware
├── migrations/     # SQL migrations
├── docs/           # Swagger output
└── utils/          # JWT, password helpers
```

## Common Endpoints

- `POST /register` — Register new user
- `POST /login` — Login and receive JWT

- `GET /projects` — List projects (auth)
- `POST /projects` — Create project (auth)

- `GET /tasks` — List tasks (auth)
- `POST /tasks` — Create task (auth)

Refer to Swagger for full list and request/response schemas.

## Development Notes

- To regenerate Swagger docs run `swag init`.
- Environment variables are loaded from `.env` using `github.com/joho/godotenv`.

## Contact

If you need help running the project, open an issue in this repository.

---

This README was updated to reflect the repository layout and usage.
├── migrations/             # Database migrations

