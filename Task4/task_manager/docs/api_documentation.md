# Task Management API Documentation

## Endpoints

### GET /tasks
Returns a list of all tasks.

**Response:**
- 200 OK: Returns an array of tasks.

### GET /tasks/:id
Returns the details of a specific task.

**Response:**
- 200 OK: Returns the task object.
- 400 Bad Request: Invalid task ID.
- 404 Not Found: Task not found.

### POST /tasks
Creates a new task.

**Request:**
- JSON body with title, description, due_date, and status.

**Response:**
- 201 Created: Returns the created task object.
- 400 Bad Request: Invalid request payload.

### PUT /tasks/:id
Updates a specific task.

**Request:**
- JSON body with title, description, due_date, and status.

**Response:**
- 200 OK: Returns the updated task object.
- 400 Bad Request: Invalid task ID or request payload.
- 404 Not Found: Task not found.

### DELETE /tasks/:id
Deletes a specific task.

**Response:**
- 204 No Content: Task deleted successfully.
- 400 Bad Request: Invalid task ID.
- 404 Not Found: Task not found.

api documentation : follow the link below:

#### https://documenter.getpostman.com/view/37611141/2sA3s3JBWo