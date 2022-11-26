package kfc_test

import (
	"fmt"
	"testing"

	kfc "github.com/sslime336/kfct"
)

func TestGetKFCPos(t *testing.T) {
	if p := kfc.PosRaw(); p == nil {
		t.FailNow()
	} else {
		fmt.Printf("%+v", p)
	}
}

func TestGetPos(t *testing.T) {
	if p := kfc.Position(); p == nil {
		t.FailNow()
	} else {
		fmt.Printf("%+v", p)
	}
}
