                      # Task Management API Documentation
    ## Overview
This documentation provides an overview of the Task Management API, which is designed using Clean Architecture principles to enhance maintainability, scalability, and testability. The API allows for managing tasks and user accounts, with endpoints for creating, updating, retrieving, and deleting tasks, as well as user registration and authentication.

##Folder Structure
The API codebase is organized into distinct layers following Clean Architecture principles:

go
Copy code
task-manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/
│   ├── task_repository.go
│   └── user_repository.go
└── Usecases/
    ├── task_usecases.go
    └── user_usecases.go
  ###Delivery Layer

main.go: Initializes the HTTP server, sets up middleware, and defines the routing configuration.
controllers/controller.go: Contains HTTP request handlers for tasks and users, interacting with use cases.
routers/router.go: Configures routes and applies middleware.

  ###Domain Layer
domain.go: Defines core business entities and logic, including Task and User models.

  ###Infrastructure Layer
auth_middleWare.go: Provides JWT authentication and authorization middleware.
jwt_service.go: Implements JWT token generation and validation.
password_service.go: Handles password hashing and comparison.

  ###Repositories Layer
task_repository.go: Defines data access operations for tasks.
user_repository.go: Defines data access operations for users.

 ###Usecases Layer
task_usecases.go: Contains business logic for task operations.
user_usecases.go: Contains business logic for user operations.

api documentation : follow the link below:

#### https://documenter.getpostman.com/view/37611141/2sA3s3JC8D