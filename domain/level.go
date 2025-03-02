package LoggerDomain

// MARK: - Objects
type Level string

// MARK: - Levels
const (
	LogLevelDebug = Level("DEBUG")
	LogLevelInfo  = Level("INFO")
	LogLevelSQL   = Level("SQL")
	LogLevelWarn  = Level("WARN")
	LogLevelError = Level("ERROR")
	LogLevelPanic = Level("PANIC")
	LogLevelFatal = Level("FATAL")
)

// MARK: - Functions
func (l Level) String() string {
	return string(l)
}
