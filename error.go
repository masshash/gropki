package gropki

import "fmt"

type GropkiError struct {
	Msg string
	Err error
}

func (gerr *GropkiError) Error() string {
	return gerr.Msg
}

func (gerr *GropkiError) Unwrap() error {
	return gerr.Err
}

func newerr(message string, err error) error {
	return &GropkiError{Msg: "gropki: " + message, Err: err}
}

func NewError(message string) error {
	return newerr(message, nil)
}

func NewErrorWithWinSyscall(apiName string, syserr error) error {
	msg := fmt.Sprintf("Windows system call [%s]: %s", apiName, syserr.Error())
	return newerr(msg, syserr)
}
