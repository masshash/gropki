package gropki

import (
	"os"
	"runtime"

	"golang.org/x/sys/windows"
)

const NULL = 0

const (
	accessRight_PROCESS_SET_QUOTA = 0x0100
	accessRight_PROCESS_TERMINATE = 0x0001
)

func (gc *GropkiCmd) start() error {
	if err := gc.Cmd.Start(); err != nil {
		return err
	}
	gc.ProcessGroup = &ProcessGroup{parentProcess: gc.Process, jobHandle: NULL}

	procHandle, err := windows.OpenProcess(accessRight_PROCESS_SET_QUOTA|accessRight_PROCESS_TERMINATE, false, uint32(gc.Process.Pid))
	if err != nil {
		gc.ProcessGroup.err = os.NewSyscallError("OpenProcess", err)
		return nil
	}
	defer windows.CloseHandle(procHandle)

	jobHandle, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		gc.ProcessGroup.err = os.NewSyscallError("CreateJobObject", err)
		return nil
	}

	err = windows.AssignProcessToJobObject(jobHandle, procHandle)
	if err != nil {
		gc.ProcessGroup.err = os.NewSyscallError("AssignProcessToJobObject", err)
		return nil
	}

	gc.ProcessGroup.jobHandle = uintptr(jobHandle)
	runtime.SetFinalizer(gc.ProcessGroup, (*ProcessGroup).Release)
	return nil
}
