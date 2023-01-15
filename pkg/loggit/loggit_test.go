package loggit

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"
)

func TestLoggit_Error(t *testing.T) {
	type args struct {
		body      string
		variables []any
	}
	tests := []struct {
		name     string
		logger   *Loggit
		args     args
		err      error
		expected string
		wantErr  bool
	}{
		{
			name: "loggit error matches expected",
			logger: &Loggit{
				txID:   "abc",
				logger: log.New(os.Stderr, "", 0),
			},
			args: args{
				body:      "it's easy as %d, %d, %d",
				variables: []any{1, 2, 3},
			},
			err:      errors.New("test error"),
			expected: "[ERROR] txid:abc: it's easy as 1, 2, 3: test error\n",
			wantErr:  false,
		},
		{
			name: "loggit error does not meet expected",
			logger: &Loggit{
				txID:   "abc",
				logger: log.New(os.Stderr, "", 0),
			},
			args: args{
				body:      "it's easy as %d, %d, %d",
				variables: []any{1, 2, 3},
			},
			err:      errors.New("another test error"),
			expected: "fail",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loggit{
				txID:   tt.logger.txID,
				logger: tt.logger.logger,
			}
			var buf bytes.Buffer
			l.logger.SetOutput(&buf)
			l.Error(tt.args.body, tt.err, tt.args.variables...)
			if buf.String() != tt.expected && !tt.wantErr {
				t.Fatalf("want %s, got %s", tt.expected, buf.String())
			}
		})
	}
}

func TestLoggit_Info(t *testing.T) {
	type args struct {
		body      string
		variables []any
	}
	tests := []struct {
		name     string
		logger   *Loggit
		args     args
		expected string
		wantErr  bool
	}{
		{
			name: "loggit info matches expected",
			logger: &Loggit{
				txID:   "abc",
				logger: log.New(os.Stderr, "", 0),
			},
			args: args{
				body:      "it's easy as %d, %d, %d",
				variables: []any{1, 2, 3},
			},
			expected: "[INFO]: txid:abc: it's easy as 1, 2, 3\n",
		},
		{
			name: "loggit info does not meet expected",
			logger: &Loggit{
				txID:   "abc",
				logger: log.New(os.Stderr, "", 0),
			},
			args: args{
				body:      "it's easy as %d, %d, %d",
				variables: []any{1, 2, 3},
			},
			expected: "fail",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loggit{
				txID:   tt.logger.txID,
				logger: tt.logger.logger,
			}
			var buf bytes.Buffer
			l.logger.SetOutput(&buf)
			l.Info(tt.args.body, tt.args.variables...)
			if buf.String() != tt.expected && !tt.wantErr {
				t.Fatalf("want %s, got %s", tt.expected, buf.String())
			}
		})
	}
}
