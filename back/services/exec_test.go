package services

import "os/exec"

type execCommandMock struct{}

func (e execCommandMock) Command(name string, args ...string) *exec.Cmd {
	var mock string = name
	switch name {
	case "docker":
		if args[0] == "images" {
			mock = "b97ca34de947;test;latest;6 hours ago;1.44GB\nca26c522ce21;matbabs/docktor;latest;6 months ago;1.39GB"
		}
	case "ls":
		if args[1] == "/" {
			mock = ".\n..\nbin\nboot\n.cache\ndev\netc\nhome\nlib\nlib32\nlib64\nlibx32\nlost+found\nmedia\nmnt\nopt\nproc\nroot\nrun\nsbin\nsnap\nsrv\nsys\ntmp\nusr\nvar"
		}
		if args[1] == "/home/user" {
			mock = ".\n..\nDocuments\nDownloads"
		}
	}
	return exec.Command("echo", mock)
}
