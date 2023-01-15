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
	txID   string
	logger *log.Logger
}

func NewLoggit() *Loggit {
	return &Loggit{
		txID:   uuid.NewString(),
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func NewLoggitWithoutTXID() *Loggit {
	return &Loggit{
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *Loggit) Error(body string, err error, variables ...any) {
	variables = append(variables, err)
	if l.txID != "" {
		logFmt := fmt.Sprintf("[ERROR] txid:%s: %s: %s", l.txID, body, "%v")
		l.logger.Output(4, fmt.Sprintf(logFmt, variables...))
		return
	}
	logFmt := fmt.Sprintf("[ERROR] %s: %s", body, "%v")
	l.logger.Output(4, fmt.Sprintf(logFmt, variables...))
}

func (l *Loggit) Info(body string, variables ...any) {
	if l.txID != "" {
		logFmt := fmt.Sprintf("[INFO]: txid:%s: %s", l.txID, body)
		l.logger.Output(4, fmt.Sprintf(logFmt, variables...))
		return
	}
	logFmt := fmt.Sprintf("[INFO] %s", body)
	l.logger.Output(4, fmt.Sprintf(logFmt, variables...))

}
