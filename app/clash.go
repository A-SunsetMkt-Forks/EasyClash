package app

import (
	"os"
	"strings"

	gp "github.com/mitchellh/go-ps"

	"easyclash/utils"
)

func NewClash() *clash {
	dir, _ := os.UserHomeDir()
	return &clash{
		dir: dir + "/.easy_clash",
	}
}

type clash struct {
	dir string
}

func (c *clash) Start() error {
	_, err := c.isRunning()
	if err != nil {
		return err
	}
	_, err = utils.Fork(c.dir+"/easy_clash_core", "-f", c.dir+"/config.yml")
	if err != nil {
		return err
	}

	return nil
}

func (c *clash) Stop() error {
	list, err := gp.Processes()
	if err != nil {
		return err
	}
	for _, p1 := range list {
		if strings.Contains(p1.Executable(), "easy_clash_core") {
			utils.KillProcess(p1.Pid())
		}
	}
	return nil
}

func (c *clash) isRunning() (int, error) {
	pid, err := c.getPid()
	if err != nil {
		return -1, err
	}
	if pid < 0 {
		return pid, nil
	}

	return pid, nil
}

func (c *clash) getPid() (int, error) {
	list, err := gp.Processes()
	if err != nil {
		return 0, err
	}
	for _, p1 := range list {
		if strings.Contains(p1.Executable(), "easy_clash_core") {
			return p1.Pid(), nil
		}
	}
	return -1, nil
}
