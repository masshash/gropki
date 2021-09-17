package gropki

import (
	"errors"
	"os"
	"runtime"

	"golang.org/x/sys/windows"
)

func checkValidJobHandle(jobHandle windows.Handle) error {
	if jobHandle == windows.InvalidHandle {
		return errors.New(EMESSAGE_PROCESSGROUP_RELEASED)
	}
	if jobHandle == NULL {
		return errors.New(EMESSAGE_PROCESSGROUP_NOTINIT)
	}
	return nil
}

func (pg *processGroup) release() error {
	jobHandle := windows.Handle(pg.jobHandle)
	if err := checkValidJobHandle(jobHandle); err != nil {
		return err
	}
	if err := windows.CloseHandle(jobHandle); err != nil {
		return os.NewSyscallError("CloseHandle", err)
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
		return errors.New(EMESSAGE_UNSUPPORTED_SIGNAL)
	}
	if err := windows.TerminateJobObject(jobHandle, 1); err != nil {
		return os.NewSyscallError("TerminateJobObject", err)
	}
	return nil
}
