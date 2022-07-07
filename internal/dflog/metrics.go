package logger

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap/zapcore"
)

var (
	logCounterVec = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "go_zap_log_total",
		Help: "Counter of the total log.",
	}, []string{"name", "level"})
)

var (
	zapLogHook = func(logName string) func(entry zapcore.Entry) error {
		return func(entry zapcore.Entry) error {
			logCounterVec.WithLabelValues(logName, entry.Level.String()).Inc()
			return nil
		}
	}
)
