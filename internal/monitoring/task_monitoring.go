package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	TaskCreate = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_create_total",
		Help: "The total number of tasks Create",
	})
	TaskUpdate = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_update_total",
		Help: "The total number of tasks Update",
	})
	TaskDelete = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_delete_total",
		Help: "The total number of tasks Delete",
	})
	TaskGetById = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_getById_total",
		Help: "The total number of tasks GetById",
	})
	TaskGetAll = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_getAll_total",
		Help: "The total number of tasks GetAll",
	})
	TasksDistributed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_distributed_total",
		Help: "The total number of tasks distributed",
	})
)

func init() {
	prometheus.MustRegister(TaskCreate)
	prometheus.MustRegister(TaskUpdate)
	prometheus.MustRegister(TaskDelete)
	prometheus.MustRegister(TaskGetById)
	prometheus.MustRegister(TaskGetAll)
	prometheus.MustRegister(TasksDistributed)
}
