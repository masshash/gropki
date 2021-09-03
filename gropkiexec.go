package gropki

import "os/exec"

type gropkiCmd struct {
	*exec.Cmd
	ProcessGroup *processGroup
}

func Command(name string, arg ...string) *gropkiCmd {
	c := exec.Command(name, arg...)
	return &gropkiCmd{Cmd: c}
}

func (gc *gropkiCmd) Start() error {
	return gc.start()
}
