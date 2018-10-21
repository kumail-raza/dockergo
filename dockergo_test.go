package dockergo

import (
	"testing"
)

func TestHello(t *testing.T) {

	expected := "Hello."
	hello := SayHello()
	if hello != expected {
		t.Fail()
	}

}
