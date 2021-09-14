package gropki

import (
	"errors"
	"os"
	"runtime"

	"golang.org/x/sys/windows"
)

func checkValidJobHandle(jobHandle windows.Handle) error {
	if jobHandle == windows.InvalidHandle {
		return errors.New("gropki: process already released")
	}
	if jobHandle == NULL {
		return errors.New("gropki: process not initialized")
	}
	return nil
}

func (pg *processGroup) release() error {
	jobHandle := windows.Handle(pg.jobHandle)
	if err := checkValidJobHandle(jobHandle); err != nil {
		return err
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
	if err := checkValidJobHandle(jobHandle); err != nil {
		return err
	}
	if sig != os.Kill {
		return errors.New("gropki: unsupported signal type")
	}
	return windows.TerminateJobObject(jobHandle, 1)
}
