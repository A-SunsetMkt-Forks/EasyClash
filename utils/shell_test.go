package utils

import (
	"fmt"
	"strings"
	"testing"

	gp "github.com/mitchellh/go-ps"
)

func TestExec(t *testing.T) {
	output, err := Exec("/Users/sandao/.easy_clash/clash", "-t", "-f", "/Users/sandao/.easy_clash/config.yml")
}

func TestFork(t *testing.T) {
	output, err := Fork("/Users/sandao/.easy_clash/clash", "-f", "/Users/sandao/.easy_clash/config.yml")
	select {}
}

func TestKillProcess(t *testing.T) {
	ok := KillProcess(43175)
	fmt.Println(ok)
}

func TestCheckProRunning(t *testing.T) {
	p, err := gp.Processes()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(p) <= 0 {
		t.Fatal("should have processes")
	}

	for _, p1 := range p {
		if strings.Contains(p1.Executable(), "clash") {
		}
	}

}
