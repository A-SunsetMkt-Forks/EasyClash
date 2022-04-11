package utils

import (
	"testing"
)

func TestSetSystemProxy(t *testing.T) {
	err := SetSystemProxy(7890, 7891)
}

func TestUnSetSystemProxy(t *testing.T) {
	err := UnSetSystemProxy()
}
