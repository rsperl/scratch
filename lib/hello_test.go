package lib

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	result := Hello()
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
