package log

import (
	"log/slog"
	"net/http"
	"os"
	// "strconv"

	// "github.com/fnxr21/item-list/pkg/middleware"
	"github.com/labstack/echo/v4"
)

// struct
type Logger struct {
	logger *slog.Logger
}

var defaultLogger *Logger

// init log
func Init() {
	var handler slog.Handler
	// Choose handler based on format
	handler = slog.NewJSONHandler(os.Stdout, nil)
	// Create the global logger
	defaultLogger = &Logger{
		logger: slog.New(handler),
	}
}

// information
func Info(r *http.Request, c echo.Context, msg string) {
	if defaultLogger == nil {
		panic("Logger is not initialized. Call log.Init() first.")
	}

	// Generate user attributes
	attrs := userAttrs(c, r)

	// Log the message with attributes
	defaultLogger.logger.Info(msg, attrs)
}

// error log
func Error(r *http.Request, c echo.Context, msg string, errCode string) {
	if defaultLogger == nil {
		panic("Logger is not initialized. Call log.Init() first.")
	}

	defaultLogger.logger.Error(
		msg,
		slog.String("code", errCode),
		userAttrs(c, r),
	)
}

// detail log
func userAttrs(c echo.Context, r *http.Request) slog.Attr {
	var attrs []any

	if r != nil {
		// var userID string
		// id := middleware.GetUserIdFromContext(c)
		// if id == -1 {
		// 	userID = "-"
		// } else {
		// 	userID = strconv.Itoa(id)
		// }

		attrs = append(attrs,
			// slog.String("UserID", userID),
			slog.String("Agent", r.UserAgent()),
			slog.String("IP", r.RemoteAddr),
			slog.String("Method", r.Method),
			slog.String("URL", r.URL.String()),
		)
	}

	return slog.Group("request", attrs...)
}
