# Ratix Backend

A robust, modular backend service for a Cinema Booking Application, built with Go. This project demonstrates a clean architecture approach to building scalable REST APIs.

## 🚀 Tech Stack

- **Language:** [Go (Golang)](https://go.dev/) 1.25+
- **Framework:** [Fiber v2](https://gofiber.io/) - High-performance web framework.
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **ORM:** [GORM](https://gorm.io/) - For database interactions and migrations.
- **Configuration:** [Viper](https://github.com/spf13/viper)
- **Validation:** [Go Playground Validator](https://github.com/go-playground/validator)

## ✨ Features

The application is organized into modular domains:

### 🎬 Movie Module
- Manage movie listings.
- Fetch movie details and categories.

### 📍 Cinema Module
- **Cinema Locations:** Browse cinemas by city (optional filtering).
- **Theaters:** Manage specific theater halls within a cinema.
- **Seat Layout:** Dynamic seat generation (Standard & Premium) with availability status.
  - *Calculation logic includes dynamic pricing based on seat type.*

### 🎫 Ticket Module
- **Booking System:** Reserve seats for specific showtimes.
- **History:** View user booking history.

### 👤 User Module
- **Authentication:** Secure user registration and login.
- **Profile:** Manage user profiles.

## 🏗 Architecture

This project follows **Clean Architecture** principles to separate concerns and ensure maintainability:

1.  **Handler Layer (`handler`)**: Manages HTTP requests and responses. Parses input and calls the Service layer.
2.  **Service Layer (`service`)**: Contains business logic. Orchestrates data flow between Handlers and Repositories.
3.  **Repository Layer (`repository`)**: Handles data access and database interactions using GORM.
4.  **Domain Layer (`domain`)**: Defines core entities and interfaces.

```
internal/
└── modules/
    ├── cinema/
    │   ├── domain/ (Entities & Interfaces)
    │   ├── repository/ (DB Implementation)
    │   ├── service/ (Business Logic)
    │   └── handler/ (HTTP Endpoints)
    ├── movie/
    ├── ticket/
    └── user/
```

## 🛠 Getting Started

### Prerequisites
- Go 1.25 or higher
- PostgreSQL

### Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/geraldiaditya/ratix-backend.git
    cd ratix-backend
    ```

2.  **Install Dependencies**
    ```bash
    go mod download
    ```

3.  **Environment Setup**
    Create a `.env` file in the root directory. You can use the example below:
    ```env
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=ratix_db
    DB_PORT=5432
    PORT=8080
    ```

4.  **Run the Application**
    ```bash
    go run cmd/server/main.go
    ```
    The server will start on `http://localhost:8080`.

## 📚 API Documentation

A Bruno collection is available in the `api-collection` directory to test the endpoints.

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
