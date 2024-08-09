 Task Management API Documentation

 MongoDB Integration

 Connecting to MongoDB
- Ensure MongoDB is running locally or use a cloud-based MongoDB instance.
- Update the connection URI in `task_service.go` to point to your MongoDB instance.

 Endpoints

The endpoints remain the same as before, but the data is now persisted in MongoDB:

 GET /tasks
Returns a list of all tasks.

 GET /tasks/:id
Returns the details of a specific task.

 POST /tasks
Creates a new task.

 PUT /tasks/:id
Updates a specific task.

 DELETE /tasks/:id
Deletes a specific task.

Make sure to follow the same response formats as before.
