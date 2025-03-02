package Logger

import (
	"context"

	"github.com/yeencloud/shared/log"

	"github.com/yeencloud/logger/domain"
)

type StandardMessage struct {
	Level LoggerDomain.Level

	Message  string
	File     string
	Line     int
	Function string
	Module   string
	Time     int64

	Fields log.Fields
}

func (m StandardMessage) WithLevel(level LoggerDomain.Level) LoggerDomain.Message {
	m.Level = level

	return m
}

func (m StandardMessage) At(time int64) LoggerDomain.Message {
	m.Time = time
	return m
}

func (m StandardMessage) WithContext(ctx context.Context) LoggerDomain.Message {
	return m
}

func (m StandardMessage) WithField(key log.Path, value interface{}) LoggerDomain.Message {
	m.Fields[key] = value

	return m
}

func (m StandardMessage) WithFields(fields log.Fields) LoggerDomain.Message {
	for key, value := range fields {
		m.Fields[key] = value
	}

	return m
}

func (m StandardMessage) Msg(message string) {
	m.Message = message

	for _, middleware := range middlewares {
		middleware.Log(m)
	}
}
