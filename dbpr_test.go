package main

import (
	"testing"
)

func TestNyx(t *testing.T) {
	nyx := "nyx"
	if nyx != nyx {
		t.Errorf("expected nyx to be 'nyx', got '%s'", nyx)
	}
}
