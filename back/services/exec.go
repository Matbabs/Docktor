package services

import "os/exec"

type execCommand interface {
	Command(name string, args ...string) *exec.Cmd
}

type execCommandImpl struct{}

func (e execCommandImpl) Command(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

var executor execCommand = execCommandImpl{}
