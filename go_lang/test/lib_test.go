package test

import (
	"testing"

	util "github.com/xdeepakv/go_lang/util"
)

func TestAddPass(t *testing.T) {
	got := util.Add(1, 1)
	if got != 2 {
		t.Errorf("Add not working properly")
	}
}
func TestAddFail(t *testing.T) {
	got := util.Add(1, 2)
	if got != 2 {
		t.Errorf("Add not working properly")
	}
}
