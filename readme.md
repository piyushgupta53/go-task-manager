# Task Manager

This is a task manager application built using Go. It leverages the Fiber web framework for handling HTTP requests and uses UUIDs for unique identification of tasks.

## Prerequisites

- Go 1.23.1 or later

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/piyushgupta53/task-manager.git
   ```

2. Navigate to the project directory:

   ```bash
   cd task-manager
   ```

3. Install the dependencies:

   ```bash
   go mod tidy
   ```

## Usage

To start the application, run:

```bash
  go run main.go
```

The application will start on the default port. You can access it by navigating to `http://localhost:3000` in your web browser.

## Dependencies

The project uses the following Go modules:

- [Fiber](https://github.com/gofiber/fiber) v2.52.5: An Express-inspired web framework written in Go.
- [UUID](https://github.com/google/uuid) v1.6.0: A package for generating and working with UUIDs.
