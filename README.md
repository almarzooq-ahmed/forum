# Forum

## Overview

**Forum** is a Go application demonstrating Clean Architecture principles. This project includes features such as user registration, user retrieval by ID, and integration with an SQLite database. It aims to showcase a modular and maintainable codebase following best practices in Go development.

## Features

- User registration
- Retrieve user information by ID
- SQLite database integration
- Clean Architecture with clear separation of concerns

## Technology Stack

- **Go**: Programming language used for the application.
- **SQLite**: Database used for storing user data.
- **GORM**: ORM used for database interactions.
- **Gorilla Mux**: Router used for handling HTTP requests.
- **Godotenv**: Library for loading environment variables from a `.env` file.

## Project Structure

Here's an overview of the project structure:

project/
│
├── cmd/
│ └── main.go # Entry point of the application
│
├── pkg/
│ ├── config/
│ │ └── config.go # Configuration management
│ │
│ ├── database/
│ │ └── database.go # Database initialization and setup
│ │
│ ├── entities/
│ │ └── user.go # Domain models and business entities
│ │
│ ├── models/
│ │ └── user.go # Database models
│ │
│ ├── repository/
│ │ └── user_repository.go # Data access and repository layer
│ │
│ ├── services/
│ │ └── user_service.go # Business logic and service layer
│ │
│ └── handlers/
│ └── user_handler.go # HTTP handlers and API logic
│
├── scripts/
│ └── migrate.go # Database migration scripts (if needed)
│
├── .env # Environment variables (optional)
├── go.mod # Go module definition
├── go.sum # Go module checksum
└── README.md # Project documentation

bash


## Setup and Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/aalmarzo/forum.git
   cd project

    Install Dependencies

    Ensure you have Go installed. Then, run:

    sh

go mod tidy

Set Up Environment Variables

Create a .env file in the root directory with the following content:

env

DATABASE_DSN=./db.sqlite
SERVER_PORT=8080

Run Migrations (if applicable)

If you have migration scripts, you can run them using:

sh

go run scripts/migrate.go

Start the Application

sh

    go run cmd/main.go

    The server will start on port 8080 (or the port specified in .env).

Configuration

Configuration values are loaded from environment variables. You can set the following in your .env file:

    DATABASE_DSN: Data Source Name for the SQLite database file.
    SERVER_PORT: Port on which the server will listen (default is 8080).

Usage

Once the application is running, you can interact with it via HTTP requests.
API Endpoints
User Registration

    Endpoint: /users/register
    Method: POST
    Request Body:

    json

{
  "username": "exampleuser",
  "email": "user@example.com",
  "password": "securepassword"
}

Response:

json

    {
      "id": 1,
      "username": "exampleuser",
      "email": "user@example.com"
    }

Retrieve User By ID

    Endpoint: /users/{id}
    Method: GET
    URL Parameters:
        id: User ID
    Response:

    json

    {
      "id": 1,
      "username": "exampleuser",
      "email": "user@example.com"
    }

Contributing

Contributions are welcome! Please follow these steps to contribute:

    Fork the Repository: Create your own fork of the repository.
    Create a Branch: Make a new branch for your feature or fix.
    Make Changes: Implement your changes and write tests if applicable.
    Submit a Pull Request: Open a pull request with a clear description of your changes.

License

This project is licensed under the MIT License. See the LICENSE file for details.