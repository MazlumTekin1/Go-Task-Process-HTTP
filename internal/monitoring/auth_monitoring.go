package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	Login = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "login_total",
		Help: "The total number of login",
	})
)

func init() {
	prometheus.MustRegister(Login)
}
