package gropki

import "os/exec"

type GropkiCmd struct {
	*exec.Cmd
	ProcessGroup *ProcessGroup
}

func Command(name string, arg ...string) *GropkiCmd {
	c := exec.Command(name, arg...)
	return &GropkiCmd{Cmd: c}
}

func (gc *GropkiCmd) Start() error {
	return gc.start()
}
