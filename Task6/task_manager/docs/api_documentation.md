# Task Management API Documentation

## Overview

The Task Management API allows users to manage tasks with persistent data storage using MongoDB. It includes user authentication and authorization using JSON Web Tokens (JWT).

## Authentication and Authorization

### Authentication

- **Register a New User**  
  `POST /register`  
  Create a new user account with a unique username and password.  
  **Request Body:**
  ```json
  {
    "username": "user1",
    "password": "password123"
  }
Response:

json
Copy code
{
  "token": "your_jwt_token_here"
}
Login
POST /login
Authenticate a user and generate a JWT token upon successful login.
Request Body:
json
Copy code
{
  "username": "user1",
  "password": "password123"
}
Response:
json
Copy code
{
  "token": "your_jwt_token_here"
}
Authorization
JWT Token
Include the JWT token in the Authorization header as Bearer <token> for accessing protected routes.
Endpoints
Public Routes
Register User
POST /register
Create a new user account.
Request Body:

json
Copy code
{
  "username": "string",
  "password": "string"
}
Login
POST /login
Authenticate a user and obtain a JWT token.
Request Body:

json
Copy code
{
  "username": "string",
  "password": "string"
}
Protected Routes
Authentication is required for these routes. Include JWT in the Authorization header.

Get All Tasks
GET /tasks
Retrieve a list of all tasks.
Response:

json
Copy code
[
  {
    "id": "task_id",
    "title": "Task Title",
    "description": "Task Description",
    "status": "open"
  }
]
Get Task by ID
GET /tasks/:id
Retrieve the details of a specific task.
Response:

json
Copy code
{
  "id": "task_id",
  "title": "Task Title",
  "description": "Task Description",
  "status": "open"
}
Create Task
POST /tasks
Create a new task.
Request Body:

json
Copy code
{
  "title": "Task Title",
  "description": "Task Description",
  "status": "open"
}
Response:

json
Copy code
{
  "id": "task_id",
  "title": "Task Title",
  "description": "Task Description",
  "status": "open"
}
Update Task
PUT /tasks/:id
Update a specific task.
Request Body:

json
Copy code
{
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "status": "completed"
}
Response:

json
Copy code
{
  "id": "task_id",
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "status": "completed"
}
Delete Task
DELETE /tasks/:id
Delete a specific task.
Response:
204 No Content

Promote User to Admin
PUT /promote/:id
Promote a user to admin role.
Response:

json
Copy code
{
  "id": "user_id",
  "username": "user1",
  "role": "admin"
}

api documentation : follow the link below:

#### https://documenter.getpostman.com/view/37611141/2sA3s3JC8D