package api

import (
	"fmt"
	"net/http"
	"strconv"
	"task/internal/models"
	"task/internal/repos"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskRepo *repos.TaskRepo
}

// NewTaskHandler creates a new instance of TaskHandler.
func NewTaskHandler(tr *repos.TaskRepo) *TaskHandler {
	return &TaskHandler{TaskRepo: tr}
}

// StandartError represents standard error response structure.
type StandartError struct {
	Error string `json:"error"`
}

// @Summary Create a new task
// @Description Create a new task with provided JSON payload
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token"
// @Param body body models.Task true "Task object to be created"
// @Success 201 {string} message "Task created successfully"
// @Failure 400 {object} StandartError "Invalid data"
// @Failure 500 {object} StandartError "Error while creating tasks"
// @Router /tasks [post]
func (t *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task

	token := c.GetHeader("Authorization")
	fmt.Println("Token : ", token)

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, StandartError{
			Error: "Invalid data",
		})
		return
	}

	res, err := t.TaskRepo.CreateTaskInDatabase(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, StandartError{
			Error: "Error while creating tasks",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": res})
}

// @Summary Retrieve all tasks
// @Description Retrieves all tasks from the database
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query int true "ID"
// @Success 200 {array} models.Task "List of tasks"
// @Failure 500 {object} StandartError "Error while getting tasks"
// @Router /tasks [get]
func (t *TaskHandler) GetTask(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("id >>>>>>>>>>>>>>", id)

	res, err := t.TaskRepo.GetTasksFromDatabase()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, StandartError{
			Error: "Error while getting tasks",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Retrieve a task by ID
// @Description Retrieves a task from the database by its ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task "Task details"
// @Failure 400 {string} string "Task does not exist"
// @Router /tasks/{id} [get]
func (t *TaskHandler) GetTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := t.TaskRepo.GetTaskByIdFromDatabase(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Task does not exist")
		return
	}

	c.JSON(http.StatusOK, task)
}

// @Summary Update a task
// @Description Updates a task in the database
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Param body body models.Task true "Updated task object"
// @Success 200 {object} models.Task "Updated task details"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Request denied"
// @Router /tasks/{id} [put]
func (t *TaskHandler) UpdateTask(c *gin.Context) {
	var task models.Task
	idStr := c.Param("id")
	task.Id, _ = strconv.Atoi(idStr)

	if err := c.BindJSON(&task); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	newTask, err := t.TaskRepo.UpdateTaskInDatabase(task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Request denied")
		return
	}

	c.JSON(http.StatusOK, newTask)
}

// @Summary Delete a task
// @Description Deletes a task from the database by its ID
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {string} string "Task deleted successfully"
// @Failure 500 {string} string "Request denied"
// @Router /tasks/{id} [delete]
func (t *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	message, err := t.TaskRepo.DeleteTaskFromDatabase(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Request denied")
		return
	}

	c.String(http.StatusOK, message)
}
