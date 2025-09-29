package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := log.With(
			slog.String("component", "middleware/logger"),
		)
		start := time.Now()
		c.Next()
		end := time.Now()
		entry := log.With(
			slog.Time("start", start),
			slog.Time("end", end),
			slog.Duration("duration", end.Sub(start)),
			slog.String("path", c.Request.URL.Path),
			slog.String("raw", c.Request.URL.RawQuery),
			slog.String("method", c.Request.Method),
			slog.Int("status_code", c.Writer.Status()),
			slog.String("user_agent", c.GetHeader("User-Agent")),
			slog.String("ip", c.ClientIP()),
			slog.String("request_id", c.GetHeader("X-Request-ID")),
		)
		entry.Info("request completed")
	}
}
