package util

import (
	"os"
	"testing"
)

func TestCreateTempScript(t *testing.T) {
	path, err := CreateTempScript(`echo "hello, guyan"`, "guyan-")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
	if err = os.Remove(path); err != nil {
		t.Fatal(err)
	}
}
