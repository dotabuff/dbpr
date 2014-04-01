package main

import (
	"testing"
)

func TestNyx(t *testing.T) {
	nyx := "nyx"
	if nyx != nyx {
		t.Errorf("nyx nyx nyx 'nyx', nyx '%s'", nyx)
	}
}
