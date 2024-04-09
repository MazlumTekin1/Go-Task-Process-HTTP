package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	UserCreate = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "users_create_total",
		Help: "The total number of users Create",
	})
	UserUpdate = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "users_update_total",
		Help: "The total number of users Update",
	})
	UserDelete = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "users_delete_total",
		Help: "The total number of users Delete",
	})
	UserGetById = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "users_getById_total",
		Help: "The total number of users GetById",
	})
	UserGetAll = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "users_getAll_total",
		Help: "The total number of users GetAll",
	})
)

func init() {
	prometheus.MustRegister(UserCreate)
	prometheus.MustRegister(UserUpdate)
	prometheus.MustRegister(UserDelete)
	prometheus.MustRegister(UserGetById)
	prometheus.MustRegister(UserGetAll)
}
