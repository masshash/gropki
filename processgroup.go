package gropki

import (
	"os"
)

type ProcessGroup struct {
	parentProcess *os.Process
	err           error

	pgid      int
	jobHandle uintptr
}

func (pg *ProcessGroup) Release() error {
	pg.parentProcess.Release()
	return pg.release()
}

func (pg *ProcessGroup) Signal(sig os.Signal) error {
	return pg.signal(sig)
}

func (pg *ProcessGroup) Kill() error {
	return pg.Signal(os.Kill)
}

func (pg *ProcessGroup) Error() error {
	return pg.err
}
