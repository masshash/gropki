package gropki

import (
	"os"
)

type processGroup struct {
	parentProc *os.Process
	err        error

	jobHandle uintptr
}

func (pg *processGroup) Release() error {
	pg.parentProc.Release()
	return pg.release()
}

func (pg *processGroup) Signal(sig os.Signal) error {
	return pg.signal(sig)
}

func (pg *processGroup) Kill() error {
	return pg.Signal(os.Kill)
}

func (pg *processGroup) Error() error {
	return pg.err
}
