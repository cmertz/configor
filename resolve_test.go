package configor_test

import (
	"os"
	"testing"

	"github.com/cmertz/configor"
)

func TestResolve(t *testing.T) {
	c := struct{ A string }{}

	err := os.Setenv("A", "C")
	if err != nil {
		t.Error(err)
	}

	err = configor.Resolve(&c, configor.Env())
	if err != nil {
		t.Error(err)
	}

	err = os.Unsetenv("A")
	if err != nil {
		t.Error(err)
	}

	if c.A != "C" {
		t.Error("expected C")
	}
}
