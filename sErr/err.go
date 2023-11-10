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

type FieldErr struct {
	Field string
	Err   string
}

func (fErr *FieldErr) Error() string {
	return fErr.Err
}

func New_Field(field, err string) error {
	return &FieldErr{
		Field: field,
		Err:   err,
	}
}

func New_Field_Errors(field, err string) error {
	return New_Errors(&FieldErr{
		Field: field,
		Err:   err,
	})
}

type Errors []error

func (errs *Errors) Error() string {
	_err := ""
	for _, v := range *errs {
		if _err == "" {
			_err = v.Error()
		} else {
			_err = fmt.Sprintf("%s,%s", _err, v.Error())
		}
	}
	return _err
}

func New_Errors(errs ...error) error {
	r := &Errors{}
	for _, v := range errs {
		*r = append(*r, v)
	}
	return r
}
