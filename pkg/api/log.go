package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Logger is the logrus logger handler
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path

		start := time.Now()
		c.Next()
		dur := time.Since(start)

		size := c.Writer.Size()
		if size < 0 {
			size = 0
		}

		status := c.Writer.Status()
		entry := log.WithFields(log.Fields{
			"status":   status,
			"latency":  dur.String(),
			"clientIP": c.ClientIP(),
			"method":   c.Request.Method,
			"path":     path,
			"size":     size,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%d %s %s", status, c.Request.Method, path)
			if status >= 500 {
				entry.Error(msg)
			} else if status >= 400 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
