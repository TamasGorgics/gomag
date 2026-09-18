package logx

type Level int

const (
	LevelSilent Level = iota + 1
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)
