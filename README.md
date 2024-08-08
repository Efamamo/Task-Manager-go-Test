
# Task-Management MongoDB

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Efamamo/Task-Management-mongoDB
   ```
2. Change to the project directory:
   ```sh
   cd Task-Management-go
   ```

## Configuration

Before running the application, ensure you have a MongoDB instance running. Update the database connection string in `data.go` where it is set with:

```go
clientOptions := options.Client().ApplyURI("your-mongodb-connection-string")
```

Replace `"your-mongodb-connection-string"` with your actual MongoDB connection string.

## Running

To run the application, use the following command:

```sh
go run .
```
