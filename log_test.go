package Logger

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yeencloud/lib-shared/log"

	"github.com/yeencloud/lib-logger/domain"
)

type TestLogger struct {
	Logs []LoggerDomain.Message
}

func (l *TestLogger) Log(message StandardMessage) {
	l.Logs = append(l.Logs, message)
}

func TestLog(t *testing.T) {
	middlewares = nil
	AddMiddleware(&TestLogger{})

	ctx := context.TODO()
	logtime := time.Now()
	scope := log.Path{
		Identifier: "http",
	}
	//
	tests := []struct {
		name    string
		logCall func()
		expect  LoggerDomain.Message
	}{
		{
			name: "Standard debug Log",
			logCall: func() {
				Log(LoggerDomain.LogLevelDebug).Msg("Hello, world !")
			},
			expect: StandardMessage{
				Level:   LoggerDomain.LogLevelDebug,
				Message: "Hello, world !",
				Fields:  log.Fields{},
			},
		},
		{
			name: "Error Log",
			logCall: func() {
				Log(LoggerDomain.LogLevelError).WithField(LoggerDomain.LogFieldError, "error message").Msg("An error occurred")
			},
			expect: StandardMessage{
				Level:   LoggerDomain.LogLevelError,
				Message: "An error occurred",
				Fields: log.Fields{
					LoggerDomain.LogFieldError: "error message",
				},
			},
		},
		{
			name: "Info Log",
			logCall: func() {
				Log(LoggerDomain.LogLevelInfo).WithContext(ctx).Msg("Request received")
			},
			expect: StandardMessage{
				Level:   LoggerDomain.LogLevelInfo,
				Message: "Request received",
				Fields:  log.Fields{},
			},
		},
		{
			name: "Timed Log",
			logCall: func() {
				Log(LoggerDomain.LogLevelDebug).At(logtime.Unix()).Msg("Request received")
			},
			expect: StandardMessage{
				Level:   LoggerDomain.LogLevelDebug,
				Message: "Request received",
				Time:    logtime.Unix(),
				Fields:  log.Fields{},
			},
		},
		{
			name: "Multiple Fields Log",
			logCall: func() {
				Log(LoggerDomain.LogLevelDebug).WithFields(log.Fields{
					log.Path{
						Parent:     &scope,
						Identifier: "method",
					}: "GET",
					log.Path{
						Parent:     &scope,
						Identifier: "path",
					}: "/",
				}).Msg("Request received")
			},
			expect: StandardMessage{
				Level:   LoggerDomain.LogLevelDebug,
				Message: "Request received",
				Fields: log.Fields{
					log.Path{
						Parent:     &scope,
						Identifier: "method",
					}: "GET",
					log.Path{
						Parent:     &scope,
						Identifier: "path",
					}: "/",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logCall()
			mw, ok := middlewares[0].(*TestLogger)

			assert.True(t, ok)
			assert.Contains(t, mw.Logs, tt.expect)
		})
	}

	middlewares = nil
}
