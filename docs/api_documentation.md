## Task Management Using MondoDB

Task MAnagement API is used for Creating, Reading, Updating and Deleting your Daily, Weekly or Monthely or Yearly Tasks which makes Your Life Easy by Managing Tasks.

## Configuration

Before running the application, ensure you have a MongoDB instance running. Update the database connection string in `data.go` where it is set with:

```go
clientOptions := options.Client().ApplyURI("your-mongodb-connection-string")
```

Replace `"your-mongodb-connection-string"` with your actual MongoDB connection string.

## GET - GetAllTasks

localhost:3000/tasks

This endpoint makes an HTTP GET request to localhost:3000/tasks to retrieve a list of tasks. The request does not include a request body. The response will have a status code of 200 and a content type of application/json. The response body will be an array of task objects, each containing an id, title, description, due date, and status. Here's an example of the response body:

```json
[{ "id": "", "title": "", "description": "", "due_date": "", "status": "" }]
```

### Example

#### Request

curl --location 'localhost:3000/tasks'

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

#### Response
```json
- Status: 200
- Content-Type: application/json
- { "id": "", "title": "", "description": "", "due_date": "", "status": ""}
```
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

#### Response
```json
- Status: 202
- Content Type: application/json
- { "id": "", "title": "", "description": "", "due_date": "", "status": ""}
```
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
```json
{
"type": "object",
"properties": {
"id": {
"type": "string"
},
"title": {
"type": "string"
},
"description": {
"type": "string"
},
"due_date": {
"type": "string"
},
"status": {
"type": "string"
}
}
}
```
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
