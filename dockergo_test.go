package dockergo

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {

	expected := "Hello."
	hello := SayHello()
	if hello != expected {
		t.Fail()
	}

}

func TestIntegrationHello(t *testing.T) {
	expected := "Hello."

	response, err := http.Get("http://localhost:3000/")
	if err != nil {
		t.Errorf("Integration test failed %s", err.Error())
	}

	if response.StatusCode != 200 {
		t.Error("Bad status code")
	}

	text, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Cannot read response body %s", err)
	}

	if string(text) != expected {
		t.Error("Invalid response")
	}

}
