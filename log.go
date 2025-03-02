package Logger

import (
	"github.com/yeencloud/lib-shared/log"

	"github.com/yeencloud/lib-logger/domain"
)

func Log(level LoggerDomain.Level) LoggerDomain.Message {
	return (&StandardMessage{
		Fields: make(log.Fields),
	}).WithLevel(level)
}
