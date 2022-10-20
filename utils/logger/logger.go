package logger

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// NewServiceLogger
// svc logger constructor
func NewServiceLogger(svcName string) log.Logger {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = level.NewFilter(logger, level.AllowDebug())
	logger = log.With(logger,
		"svc", svcName,
		"ts", log.DefaultTimestampUTC,
		"clr", log.DefaultCaller,
	)

	return logger
}
