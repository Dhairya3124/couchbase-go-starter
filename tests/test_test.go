package tests

import (
	"testing"
)

func TestExampleFunction(t *testing.T) {
	t.Run("Test",func (t* testing.T)  {
		got:="Jack"
		want:="Jade"
		if got!=want{
			t.Errorf("got %v,want %v",got,want)
		}

	})
}
