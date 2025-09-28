package logx

type Level int

const (
	LevelSilent Level = iota + 1
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)
