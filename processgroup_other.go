//go:build plan9 || js || wasm

package gropki

import (
	"errors"
	"os"
)

func (pg *ProcessGroup) release() error {
	return errors.New(EMESSAGE_UNSUPPORTED_PLATFORM)
}

func (pg *ProcessGroup) signal(sig os.Signal) error {
	return errors.New(EMESSAGE_UNSUPPORTED_PLATFORM)
}
