package logze

type LogLevel int8

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	WarningLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func (logLevel LogLevel) String() string {
	switch logLevel {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}
	return "unknown"
}

var AllLogLevels = []LogLevel{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarningLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}
