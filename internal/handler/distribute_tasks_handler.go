package handler

import (
	"encoding/json"
	"net/http"
	"task-process-service/internal/domain"
	m "task-process-service/internal/monitoring"
	"task-process-service/internal/service"
)

type DistributeTasksHandler struct {
	sTask service.TaskService
	sUser service.UserService
}

func NewDistributeTasksHandler(sTask service.TaskService, sUser service.UserService) *DistributeTasksHandler {
	return &DistributeTasksHandler{
		sTask: sTask,
		sUser: sUser,
	}
}

// @Summary Distribute tasks
// @Description Distribute tasks to users
// @Tags Distribute Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true " you must start Bearer and then space and then token"
// @Success 200 {object} map[string]interface{}
// @Router /distributeTasks [get]
func (h DistributeTasksHandler) DistributeTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	usersChan := make(chan []domain.UserGetDataList)
	tasksChan := make(chan []domain.TaskGetDataList)
	errChan := make(chan error)
	go func() {
		tasks, err := h.sTask.GetAll()
		if err != nil {
			errChan <- err
			return
		}
		tasksChan <- tasks
	}()

	go func() {
		users, err := h.sUser.GetAll()
		if err != nil {
			errChan <- err
			return
		}
		usersChan <- users
	}()

	var users []domain.UserGetDataList
	var tasks []domain.TaskGetDataList
	for i := 0; i < 2; i++ {
		select {
		case users = <-usersChan:
		case tasks = <-tasksChan:
		}
	}

	data, maxWeeks := service.DistributeTasks(users, tasks)

	result := map[string]interface{}{
		"data":     data,
		"maxWeeks": maxWeeks,
	}

	m.TasksDistributed.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
