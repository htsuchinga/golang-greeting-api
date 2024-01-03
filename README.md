# Golang Greeting API with net/http Package

Welcome to the documentation for your Golang API project built with `net/http`. This API provides a simple greeting functionality.

## Getting Started

These instructions will help you set up the project on your local machine.

### Prerequisites

- Go installed on your machine
- Docker and Docker Compose

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/your-api-project.git
    cd your-api-project
    ```

2. Run the API server using docker-compose commands

    ```bash
    docker-compose up -d --build
    docker-compose exec golang go run api/main.go
    ```

Now your API should be running locally.

## Endpoints

### `POST /v1/greeting/hello`

- **Request:**
  - Method: `POST`
  - Path: `/v1/greeting/hello`
  - Body: `{ "name": "your-name" }`

- **Responses:**
  - Success (200 OK):
    ```json
    {
      "resultCd": 1000,
      "contents": {
        "message": "Hello, {name}!",
        "receiveDate": "YYYYMMDD",
        "receiveTime": "hhmmss"
      }
    }
    ```
  - Validation Error (200 OK):
    ```json
    {
      "resultCd": 2000
    }
    ```
  - System Error (200 OK):
    ```json
    {
      "resultCd": 9000
    }
    ```
    