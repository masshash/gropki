//go:build plan9 || js || wasm

package gropki

import (
	"errors"
)

func (gc *GropkiCmd) start() error {
	return errors.New("gropki: unsupported platform")
}
