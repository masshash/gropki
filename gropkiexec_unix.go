//go:build !windows && !plan9 && !js && !wasm

package gropki

import "syscall"

func (gc *GropkiCmd) start() error {
	if gc.Cmd.SysProcAttr != nil {
		gc.Cmd.SysProcAttr.Setpgid = true
	} else {
		gc.Cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}

	if err := gc.Cmd.Start(); err != nil {
		return err
	}
	gc.ProcessGroup = &ProcessGroup{parentProcess: gc.Process, pgid: gc.Process.Pid}

	return nil
}
