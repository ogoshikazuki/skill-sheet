package entity

type Logger interface {
	Errorf(format string, v ...any)
}
