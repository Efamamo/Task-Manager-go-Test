# Task-Management MongoDB

## Installation

1. Clone the repository:
   
   ```sh
   git clone https://github.com/Efamamo/Task-Management-mongoDB
   
2. Change to the project directory:

  ```
  cd /Task-Management-go

## Configuration

Before running the application, ensure you have a MongoDB instance running. Update the database connection string in data.go where it is set with:

```sh
clientOptions := options.Client().ApplyURI("your-mongodb-connection-string")
```
## Running

```sh
go run .```
