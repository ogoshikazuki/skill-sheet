package logger

import (
	"log"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type standardLogger struct {
	logger *log.Logger
}

func (s *standardLogger) Errorf(format string, v ...any) {
	s.logger.Printf(format, v...)
}

func NewStandardLogger() entity.Logger {
	return &standardLogger{
		logger: log.Default(),
	}
}
