            Documentation on  Task Management API

Endpoints

GET /tasks
Returns a list of all tasks.

Response:

200 OK: Returns an array of tasks.
GET /tasks/
Returns the details of a specific task.

Response:

200 OK: Returns the task object.
400 Bad Request: Invalid task ID.
404 Not Found: Task not found.
POST /tasks
Creates a new task.

Request:

JSON body with the following fields:
title (string, required)
description (string, optional)
due_date (string, optional, date in ISO 8601 format)
status (string, optional, one of "pending", "in-progress", "completed")
Response:

201 Created: Returns the created task object.
400 Bad Request: Invalid request payload.
PUT /tasks/
Updates a specific task.

Request:

JSON body with the following fields:
title (string, optional)
description (string, optional)
due_date (string, optional, date in ISO 8601 format)
status (string, optional, one of "pending", "in-progress", "completed")
Response:

200 OK: Returns the updated task object.
400 Bad Request: Invalid task ID or request payload.
404 Not Found: Task not found.
DELETE /tasks/
Deletes a specific task.

Response:

204 No Content: Task deleted successfully.
400 Bad Request: Invalid task ID.
404 Not Found: Task not found.

