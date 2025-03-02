package Logger

import (
	"github.com/yeencloud/shared/log"

	"github.com/yeencloud/logger/domain"
)

func Log(level LoggerDomain.Level) LoggerDomain.Message {
	return (&StandardMessage{
		Fields: make(log.Fields),
	}).WithLevel(level)
}
