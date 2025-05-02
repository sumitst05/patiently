# Patiently 🏥

**Patiently** is a patient management platform built with a Go backend and a minimal HTML-CSS-JavaScript frontend. It supports role-based access for **Receptionists** and **Doctors**, enabling easy handling of patient records.

🔗 **Live Demo**: [https://patiently-latest.onrender.com/](https://patiently-latest.onrender.com/)

## ✨ Features

- 🔐 **Authentication**: Sign up & login with JWT-based token system with httpOnly cookie
- 🧑‍⚕️ **Role-Based Access**:
  - **Receptionists**: Full CRUD access to patient records
  - **Doctors**: View & update access only
- 🖥️ **Frontend**: Clean, minimal interface built with HTML, CSS, and JS
- 🧱 **Architecture**: Layered (Handler → Service → Repository)
- 🔒 **Security**: Passwords hashed with bcrypt
- 🧪 **Testing**: Unit tests using `testify` and `httptest`

## 🛠️ Tech Stack

| Layer        | Tech                        |
| ------------ | --------------------------- |
| **Backend**  | Go (Gin, GORM, JWT, bcrypt) |
| **Database** | PostgreSQL                  |
| **Frontend** | HTML, CSS, JavaScript       |
| **Testing**  | `testify`, `httptest`       |

## 📁 Project Structure

```
patiently/
├── cmd/
│   └── server/                 # Entry point for starting the server
│       └── main.go
├── internal/
│   ├── config/                 # Configuration files and env setup
│   ├── handler/                # HTTP handlers for each feature
│   ├── middleware/            # Custom Gin middlewares
│   ├── models/                 # Data models
│   ├── repository/             # DB access layer (CRUD operations)
│   ├── router/                 # Route definitions
│   ├── service/                # Business logic layer
│   └── utils/                  # Utility functions (hashing, JWT, etc.)
├── static/                     # Minimal HTML,CSS,JS frontend
│   ├── index.html
│   ├── receptionist.html
│   └── doctor.html
│   └── history.html
├── test/
│   └── handler/                # Unit tests for handlers
│       └── auth_test.go
│       └── patient_test.go
├── go.mod
└── README.md
```

## API Documentation

For complete API documentation, visit the Postman collection:

👉 [View API Docs on Postman](https://documenter.getpostman.com/view/28183247/2sB2j4eWMV)

This includes:

- Authentication endpoints (`/signup`, `/signin`, `/logout`)
- Patient management endpoints (CRUD operations with role-based access)
- Expected request-response formats

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL
- Git

### Clone the repository

```bash
git clone https://github.com/sumitst05/patiently.git
cd patiently
```

### Setup environment variables

Create a `.env` file in the root directory with the following variables:

```env
PORT=3000
MODE=dev
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_username
DB_PASSWORD=your_postgres_password
DB_NAME=patiently
DB_SSLMODE=disable
JWT_SECRET=your_jwt_secret
```

### Create the database

Login to PostgreSQL and create the database:

```bash
psql -U your_postgres_user
CREATE DATABASE patiently;
```

### Run the server

```bash
go run cmd/server/main.go
```

### Development with Live Reload (Air)

This project supports live reloading with [Air](https://github.com/cosmtrek/air). If you don't have it installed:

```bash
go install github.com/cosmtrek/air@latest
```

To start the development server with live reload:

```bash
air
```

### Access the platform

Open your browser and go to:

```
http://localhost:3000
```

### Running Tests

Make sure your database is properly set up and the environment variables are configured.

##### Additionally put another .env file in ./test/ directory and replace the Database Name with test database's name

```bash
go test ./... -v
```
