package tests

import (
	"testing"

	"github.com/rsperl/scratch/lib"
)

func TestHelloIntegration(t *testing.T) {
	expected := "hello integration"
	result := lib.HelloIntegration()
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
