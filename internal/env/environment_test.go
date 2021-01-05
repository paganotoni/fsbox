package env

import (
	"os"
	"testing"
)

func TestCurrent(t *testing.T) {
	err := os.Setenv("GO_ENV", "xx")
	if err != nil {
		t.Fail()
	}

	e := Current()
	if e != "xx" {
		t.Errorf("env should be xx, got %v", e)
	}

	err = os.Setenv("GO_ENV", "")
	if err != nil {
		t.Fail()
	}

	e = Current()
	if e != Development {
		t.Errorf("env should be %v, got %v", Development, e)
	}

}
