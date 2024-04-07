package handler

import (
	"encoding/json"
	"net/http"
	"task-process-service/internal/domain"
	"task-process-service/internal/service"

	"github.com/prometheus/client_golang/prometheus"
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

var (
	tasksDistributed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_distributed_total",
		Help: "The total number of tasks distributed",
	})
)

func init() {
	prometheus.MustRegister(tasksDistributed)
}

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

	tasksDistributed.Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
