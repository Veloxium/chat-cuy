package protected

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// logger middleware function

func Logger(db *sql.DB) gin.HandlerFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime).Milliseconds()
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		loggingEntry := logger.WithFields(logrus.Fields{
			"status":    statusCode,
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"latency":   duration,
			"client_ip": clientIP,
		})

		var level, message string

        // logging entry every response statuscode

		switch {
		case statusCode >= 500:
			level = "error"
			message = "internal server error"
			loggingEntry.Error("internal server error")
		case statusCode >= 400:
			level = "warn"
			message = "client bad request"
			loggingEntry.Warn("client bad request")
		case statusCode >= 300:
			level = "info"
			message = "redirect request"
			loggingEntry.Info("redirect request")
		default:
			level = "info"
			message = "request successfully"
			loggingEntry.Info("request successfully")
		}

        //insert log to database
		query := `INSERT INTO logs (level, status_code, method, path, latency, client_ip, message, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, NOW()) returning id , created_at`
		ctx, cancle := context.WithTimeout(c, time.Duration(2)*time.Second)
		defer cancle()

        //execute query to database
		_, err := db.ExecContext(ctx, query, level, statusCode, c.Request.Method, c.Request.URL.Path, duration, clientIP, message)
		if err != nil {
			loggingEntry.Error("failed to save log to database: ", err)
		}

	}
}
