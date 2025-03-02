package LoggerDomain

type Logger interface {
	Log(Level Level) Message
}
