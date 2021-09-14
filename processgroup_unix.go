//go:build !windows && !plan9 && !js && !wasm

package gropki

import (
	"errors"
	"os"
	"syscall"
)

func (pg *processGroup) release() error {
	if pg.pgid == -1 {
		return errors.New("gropki: process already released")
	}
	pg.pgid = -1
	return nil
}

func (pg *processGroup) signal(sig os.Signal) error {
	if pg.pgid == -1 {
		return errors.New("gropki: process already released")
	}
	s, ok := sig.(syscall.Signal)
	if !ok {
		return errors.New("gropki: unsupported signal type")
	}
	return syscall.Kill(-pg.pgid, s)
}
