//go:build !windows && !plan9 && !js && !wasm

package gropki

import (
	"errors"
	"os"
	"syscall"
)

func checkValidPgid(pgid int) error {
	if pgid == -1 {
		return errors.New("gropki: process already released")
	}
	return nil
}

func (pg *processGroup) release() error {
	if err := checkValidPgid(pg.pgid); err != nil {
		return err
	}
	pg.pgid = -1
	return nil
}

func (pg *processGroup) signal(sig os.Signal) error {
	if err := checkValidPgid(pg.pgid); err != nil {
		return err
	}
	s, ok := sig.(syscall.Signal)
	if !ok {
		return errors.New("gropki: unsupported signal type")
	}
	return syscall.Kill(-pg.pgid, s)
}
