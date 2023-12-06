package logs

import "go.uber.org/zap"

type ManagerLogs struct {
	Logger *zap.Logger
}

func NewManagerLogs(logger *zap.Logger) *ManagerLogs {
	return &ManagerLogs{
		Logger: logger,
	}
}

func(l ManagerLogs) NewInfoLog(msg, action, service string) {
	l.Logger.Info(msg,
		zap.String("action", action),
		zap.String("service", service),
	)
}
