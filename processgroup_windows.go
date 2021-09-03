package gropki

import (
	"errors"
	"os"
	"runtime"

	"golang.org/x/sys/windows"
)

func (pg *processGroup) release() error {
	jobHandle := windows.Handle(pg.jobHandle)
	if jobHandle == windows.InvalidHandle {
		return errors.New("gropki: process already released")
	}
	if err := windows.CloseHandle(jobHandle); err != nil {
		return err
	}
	pg.jobHandle = uintptr(windows.InvalidHandle)
	runtime.SetFinalizer(pg, nil)
	return nil
}

func (pg *processGroup) signal(sig os.Signal) error {
	jobHandle := windows.Handle(pg.jobHandle)
	if jobHandle == windows.InvalidHandle {
		return errors.New("gropki: process already released")
	}
	if pg.err != nil {
		return pg.err
	}
	if sig != os.Kill {
		return errors.New("gropki: unsupported signal type")
	}
	return windows.TerminateJobObject(jobHandle, 1)
}
