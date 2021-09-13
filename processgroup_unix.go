//go:build !windows && !plan9 && !js && !wasm

package gropki

import (
	"errors"
	"os"
	"syscall"
)

func (pg *processGroup) release() error {
	return nil
}

func (pg *processGroup) signal(sig os.Signal) error {
	if pg.parentProcess.Pid == -1 {
		return errors.New("gropki: process already released")
	}
	s, ok := sig.(syscall.Signal)
	if !ok {
		return errors.New("gropki: unsupported signal type")
	}
	return syscall.Kill(-pg.parentProcess.Pid, s)
}
