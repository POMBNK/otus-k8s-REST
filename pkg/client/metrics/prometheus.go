package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
	"time"
)

var mu = &sync.Mutex{}

type Prometheus struct {
	gatherer      prometheus.Gatherer
	latency       *prometheus.SummaryVec
	rps           *prometheus.CounterVec
	errorRate     *prometheus.CounterVec
	scriptRuntime prometheus.Histogram
}

func New() *Prometheus {
	registry := prometheus.NewRegistry()

	p := initScriptRuntimeMetrics(registry)

	timer := prometheus.NewTimer(p.scriptRuntime)
	defer timer.ObserveDuration()

	return &p
}

func (p *Prometheus) Handle(app fiber.Router, url string, handlers ...fiber.Handler) {
	h := append(handlers, adaptor.HTTPHandler(promhttp.HandlerFor(p.gatherer, promhttp.HandlerOpts{})))
	app.Get(url, h...)
}

func initScriptRuntimeMetrics(registry prometheus.Registerer) Prometheus {

	latency := promauto.With(registry).NewSummaryVec(
		prometheus.SummaryOpts{
			Name:        "api_latency_seconds",
			Help:        "API response latency in seconds.",
			ConstLabels: nil,
			Objectives: map[float64]float64{
				0.5:  0.05,
				0.95: 0.01,
				0.99: 0.005,
			},
		},
		[]string{"latency"})

	rps := promauto.With(registry).NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total API requests.",
		},
		[]string{"rps"})

	errorRate := promauto.With(registry).NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_errors_total",
			Help: "Total number of errors (500 responses).",
		},
		[]string{"errorRate"},
	)

	scriptRuntimeHistogram := promauto.With(registry).NewHistogram(prometheus.HistogramOpts{
		Name:    "kuber_rest_script_runtime",
		Help:    "The duration of script runtime",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})

	gatherer, ok := registry.(prometheus.Gatherer)
	if !ok {
		gatherer = prometheus.DefaultGatherer
	}

	return Prometheus{
		gatherer, latency, rps, errorRate, scriptRuntimeHistogram}
}

func (p *Prometheus) MetricMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		method := c.Method()

		mu.Lock()
		p.rps.WithLabelValues(method).Inc()
		mu.Unlock()
		//routeName := string(c.Request().RequestURI())
		//p.rps.WithLabelValues(routeName).Inc()

		err := c.Next()
		mu.Lock()
		p.latency.WithLabelValues(method).Observe(time.Since(start).Seconds())
		//p.latency.WithLabelValues(routeName).Observe(time.Since(start).Seconds())
		mu.Unlock()
		if c.Response().StatusCode() >= http.StatusInternalServerError {
			mu.Lock()
			p.errorRate.WithLabelValues(method).Inc()
			mu.Unlock()
			//p.errorRate.WithLabelValues(routeName).Inc()
		}

		return err
	}
}

//todo: добавить в кубере еще прометеус, указать путь до приложения+/metrics

/*
Шаг 2: Настройка Grafana
После настройки вашего сервера и сбора метрик, вам нужно будет наполнять ваш дашборд в Grafana.

Создайте новый дашборд.

Добавьте панель для каждой метрики:

Latency: используйте запрос sum(rate(api_latency_seconds_count[1m])) by (method) для графика с квантилями.
RPS: используйте запрос sum(rate(api_requests_total[1m])) by (method) для графика.
Error Rate: используйте запрос sum(rate(api_errors_total[1m])) by (method) для графика.
Общие метрики с nginx-ingress-controller:

Для общего Latency используйте sum(rate(nginx_ingress_controller_latency_seconds_sum[1m])) by (ingress) с квантилями.
Для общего RPS используйте sum(rate(nginx_ingress_controller_requests[1m])).
Для общего Error Rate используйте sum(rate(nginx_ingress_controller_errors_total[1m])).
*/
