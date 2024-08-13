package controllers

import (
	"net/http"
	"task_manager/domain"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUseCase usecases.TaskUseCase
}

func NewTaskController(taskUseCase usecases.TaskUseCase) *TaskController {
	return &TaskController{TaskUseCase: taskUseCase}
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := t.TaskUseCase.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdTask)
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask, err := t.TaskUseCase.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := t.TaskUseCase.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func (t *TaskController) GetTasks(c *gin.Context) {
	tasks, err := t.TaskUseCase.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (t *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := t.TaskUseCase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

type UserController struct {
	UserUseCase usecases.UserUseCase
}

func NewUserController(userUseCase usecases.UserUseCase) *UserController {
	return &UserController{UserUseCase: userUseCase}
}

func (u *UserController) RegisterUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := u.UserUseCase.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

func (u *UserController) LoginUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loggedInUser, err := u.UserUseCase.LoginUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loggedInUser)
}

func (u *UserController) PromoteUser(c *gin.Context) {
	id := c.Param("id")
	promotedUser, err := u.UserUseCase.PromoteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promotedUser)
}
