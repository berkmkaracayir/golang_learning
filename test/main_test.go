//go mod init main.go
//go mod tidy
//go test

package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sum(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d.", expected, result)
	}
}

