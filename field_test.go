// nolint:testpackage
package configor

import (
	"testing"
)

func TestField(t *testing.T) {
	t.Run("fieldFor", func(t *testing.T) {
		i := "4"
		f := fieldFor(&i)
		if f.v.String() != "4" {
			t.Error("expected 4")
		}

		b := true
		f = fieldFor(&b)
		if f.v.Bool() != true {
			t.Error("expected `true`")
		}
	})

	t.Run("set", func(t *testing.T) {
		i := "4"
		f := fieldFor(&i)
		err := f.set("5")
		if err != nil {
			t.Error(err)
		}

		if i != "5" {
			t.Error("expected 5")
		}

		b := false
		f = fieldFor(&b)
		err = f.set("true")
		if err != nil {
			t.Error(err)
		}

		if !b {
			t.Error("expected b to be `true`")
		}
	})
}
