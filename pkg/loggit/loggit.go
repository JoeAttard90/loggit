package loggit

import (
	"fmt"
	"github.com/google/uuid"
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
		correlationID: uuid.NewString(),
		logger:        log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *Loggit) Error(body string, err error, variables ...any) {
	variables = append(variables, err)
	logFmt := fmt.Sprintf("[ERROR] txid:%s: %s: %s", l.correlationID, body, "%v")
	l.logger.Output(4, fmt.Sprintf(logFmt, variables...))
}

func (l *Loggit) Info(body string, variables ...any) {
	logFmt := fmt.Sprintf("[INFO]: txid:%s: %s", l.correlationID, body)
	l.logger.Output(4, fmt.Sprintf(logFmt, variables...))
}
