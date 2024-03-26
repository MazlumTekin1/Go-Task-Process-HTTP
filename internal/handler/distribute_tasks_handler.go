package handler

import (
	"encoding/json"
	"net/http"
	"task-process-service/internal/domain"
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
// @Tags distribute-tasks
// @ID distribute-tasks
// @Accept json
// @Produce json
// @Success 200
// @Router /distribute-tasks [get]
func (h DistributeTasksHandler) DistributeTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	usersChan := make(chan []domain.UserGetDataList)
	tasksChan := make(chan []domain.TaskGetDataList)
	errChan := make(chan error)
	go func() {
		tasks, err := h.sTask.GetAllTask()
		if err != nil {
			errChan <- err
			return
		}
		tasksChan <- tasks
	}()

	go func() {
		users, err := h.sUser.GetAllUser()
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
