package loggit

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Error(body string, err error, variables ...any)
	Info(body string, variables ...any)
}

type Loggit struct {
	correlationID string
	logger        *log.Logger
}

func NewLoggit() *Loggit {
	return &Loggit{
		logger: log.New(os.Stderr, "", log.LstdFlags|log.Llongfile),
	}
}

func (l *Loggit) Error(body string, err error, variables ...any) {
	variables = append(variables, err)
	logFmt := fmt.Sprintf("[ERROR] %s -- %s: %s", l.correlationID, body, "%v")
	l.logger.Output(4, fmt.Sprintf(logFmt, variables))
}

func (l *Loggit) Info(body string, variables ...any) {
	logFmt := fmt.Sprintf("[INFO] %s -- %s: %s", l.correlationID, body, "%v")
	l.logger.Output(4, fmt.Sprintf(logFmt, variables))
}
