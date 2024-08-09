## Task Management Using MondoDB

Task Management API is used for Creating, Reading, Updating and Deleting your Daily, Weekly or Monthely or Yearly Tasks which makes Your Life Easy by Managing Tasks.

## Configuration

Before running the application, ensure you have a MongoDB instance running. Update the database connection string in `data.go` where it is set with:

```go
clientOptions := options.Client().ApplyURI("your-mongodb-connection-string")
```

Replace `"your-mongodb-connection-string"` with your actual MongoDB connection string.

## Authentication and Authorization

User authentication and Authorization is also Included. Some of the endpoints are protected for logged in users only so they want authorization header such as:

```json
GET localhost:3000/tasks
GET localhost:3000/tasks/:id
```

Some Other Endpoints also want Authorization header and Are also Only Permitted to Loggedin Admins such as:

```json
POST localhost:3000/tasks
DELETE localhost:3000/tasks/:id
PUT localhost:3000/tasks/:id
PATCH localhost:3000/promote
```

## POST - Signup

localhost:3000/register

This endpoint makes an HTTP POST request to localhost:3000/register to signup a user. The request include a request body that contains username and password. The response will have a status code of 201 if the user successfully signs up (When signing up a user if there is no user inside db it makes the signed up user admin else regular user) or 409 if there is already another user with the same username. When saving a user inside a database the entered password is hashed. The response body will be user object containing an id, username, hashedPassword,and isAdmin boolean. Here's an example of the response body:

### Example

#### Request

The request body should be in raw format and include the following parameters:

- username (string): The username.
- password (string): Plain Password

#### Response

```json
{
  "_id": "66b5c3300582f8449ccf7449",
  "username": "bini",
  "password": "$2a$10$GbSmtaskMhQjPsO1VcDsWuskUi4LA7V6m8.gcyBXklPBSXtNXia82",
  "isAdmin": false
}
```

## POST - Login

localhost:3000/login

This endpoint makes an HTTP POST request to localhost:3000/login to login a user. The request include a request body that contains username and password. The response will have a status code of 201 if the user successfully logs in and responds with a token created using username and role. or 404 if there is no user with provided username or 401 if the user exists but password is wrong. The response body will be jwt token string.

### Example

#### Request

The request body should be in raw format and include the following parameters:

- username (string): The username.
- password (string): Plain Password

#### Response
```json
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxODgxOTUsImlzQWRtaW4iOnRydWUsInVzZXJuYW1lIjoiYmVrYSJ9.KISrov7DvDsjSli_jUvgfvtiz07xFYQd3Wjpy9VTWL8
```


## PATCH - PromoteUser

localhost:3000/patch

This endpoint makes an HTTP PATCH request to localhost:3000/promote to promote a user from regular user to admin. This end point is protected so regular users cant access it. The request does not include a request body. The response will have a status code of 203 if the action is successful. The response body will be an object telling if the action is successful or not. Here's an example of the response body:

- Method: PATCH
- Endpoint: localhost:3000/promote?username=abeni

#### Response
```json
"message": "User promoted"
```

## GET - GetAllTasks

localhost:3000/tasks

This endpoint makes an HTTP GET request to localhost:3000/tasks to retrieve a list of tasks. The request does not include a request body. The response will have a status code of 200 and a content type of application/json. The response body will be an array of task objects, each containing an id, title, description, due date, and status. Here's an example of the response body:

### Example

#### Request

- Method: GET
- Endpoint: localhost:3000/tasks

#### Response

```json
[
  {
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2024-08-07T09:41:00.564238382+03:00",
    "status": "Pending"
  },
  {
    "id": "2",
    "title": "Task 2",
    "description": "Second task",
    "due_date": "2024-08-08T09:41:00.564241718+03:00",
    "status": "In Progress"
  },
  {
    "id": "3",
    "title": "Task 3",
    "description": "Third task",
    "due_date": "2024-08-09T09:41:00.564361323+03:00",
    "status": "Completed"
  }
]
```

## GET - GetTaskByID

localhost:3000/tasks/:id

This endpoint retrieves a specific task by its ID.

#### Request

- Method: GET
- Endpoint: localhost:3000/tasks/:id


#### Example Response

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "First task",
  "due_date": "2024-08-07T09:41:00.564238382+03:00",
  "status": "Pending"
}
```

## PUT UpdateTaskByID

localhost:3000/tasks/:id

This endpoint allows you to update a specific task identified by its ID. The request should be sent to localhost:3000/tasks/:id using the HTTP PUT method.

#### Request Body

The request body should be in raw format and include the following parameters:

- Title (string): The updated title of the task.
- description (string): The updated description of the task.
- status (string): The updated status of the task.
- due_date (string): The updated due date of the task.

#### Response

Upon successful execution, the endpoint returns a status code of 201 and a JSON object with the updated details of the task, including the following properties:

- id (string): The ID of the task.
- title (string): The title of the task.
- description (string): The description of the task.
- due_date (string): The due date of the task.
- status (string): The status of the task.

#### Example Response

```json
{
  "id": "66b47d33320e17e93d99daab",
  "title": "Do Authentication",
  "description": "Do Authentication and Authorization by the morning",
  "due_date": "2024-08-07T09:41:00.564238382+03:00",
  "status": "Pending"
}
```

## DELETE DeleteTask

localhost:3000/tasks/:id

This endpoint is used to delete a specific task identified by its ID.

#### Request

- Method: DELETE
- URL: localhost:3000/tasks/:id


#### Example Response

```json
{
  "_id": "66b47d33320e17e93d99da22",
  "title": "Task 2",
  "description": "Second task",
  "due_date": "2024-08-08T09:41:00.564241718+03:00",
  "status": "In Progress"
}
```

## POST AddTask

localhost:3000/tasks

This endpoint is used to create a new task.

#### Request Body

- title (string, required): The title of the task.
- description (string, required): The description of the task.
- due_date (string, required): The due date of the task.
- status (string, required): The status of the task.

#### Response

The response is in JSON format with the following schema:

#### Example Response

```json
{
  "id": "66b47d33320e17e93d99da22",
  "title": "Task 4",
  "description": "Fourth task",
  "due_date": "2024-08-08T09:41:00.564241718+03:00",
  "status": "Pending"
}
```
