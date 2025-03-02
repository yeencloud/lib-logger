package LoggerDomain

import "testing"

func TestLevel(t *testing.T) {
	tests := []struct {
		name     string
		level    Level
		expected string
	}{
		{
			name:     "Debug Level",
			level:    LogLevelDebug,
			expected: "DEBUG",
		},
		{
			name:     "Info Level",
			level:    LogLevelInfo,
			expected: "INFO",
		},
		{
			name:     "SQL Level",
			level:    LogLevelSQL,
			expected: "SQL",
		},
		{
			name:     "Warn Level",
			level:    LogLevelWarn,
			expected: "WARN",
		},
		{
			name:     "Error Level",
			level:    LogLevelError,
			expected: "ERROR",
		},
		{
			name:     "Panic Level",
			level:    LogLevelPanic,
			expected: "PANIC",
		},
		{
			name:     "Fatal Level",
			level:    LogLevelFatal,
			expected: "FATAL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.String(); got != tt.expected {
				t.Errorf("Level.String() = '%v', want '%v'", got, tt.expected)
			}
		})
	}
}
