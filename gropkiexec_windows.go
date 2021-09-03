package gropki

import (
	"runtime"

	"golang.org/x/sys/windows"
)

const (
	accessRight_PROCESS_SET_QUOTA = 0x0100
	accessRight_PROCESS_TERMINATE = 0x0001
)

func (gc *gropkiCmd) start() error {
	if err := gc.Cmd.Start(); err != nil {
		return err
	}
	gc.ProcessGroup = &processGroup{parentProc: gc.Process, jobHandle: 0}

	procHandle, err := windows.OpenProcess(accessRight_PROCESS_SET_QUOTA|accessRight_PROCESS_TERMINATE, false, uint32(gc.Process.Pid))
	if err != nil {
		gc.ProcessGroup.err = err
		return nil
	}
	defer windows.CloseHandle(procHandle)

	jobHandle, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		gc.ProcessGroup.err = err
		return nil
	}

	err = windows.AssignProcessToJobObject(jobHandle, procHandle)
	if err != nil {
		gc.ProcessGroup.err = err
		return nil
	}

	gc.ProcessGroup.jobHandle = uintptr(jobHandle)
	runtime.SetFinalizer(gc.ProcessGroup, (*processGroup).Release)
	return nil
}
