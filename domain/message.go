package LoggerDomain

import (
	"context"

	"github.com/yeencloud/shared/log"
)

type Message interface {
	WithLevel(level Level) Message

	At(time int64) Message

	WithContext(ctx context.Context) Message

	WithField(key log.Path, value interface{}) Message
	WithFields(fields log.Fields) Message

	Msg(message string)
}
