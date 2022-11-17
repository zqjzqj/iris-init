package sErr

import "fmt"

type SErr struct {
	err string
}

func New(msg string) *SErr {
	return &SErr{err: msg}
}

func NewFmt(format string, a ...any) *SErr {
	return New(fmt.Sprintf(format, a...))
}

func NewByError(err error) *SErr {
	return &SErr{err: err.Error()}
}

func (errS *SErr) Error() string {
	return errS.err
}
