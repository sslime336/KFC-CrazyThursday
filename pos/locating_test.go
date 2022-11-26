package pos

import (
	"fmt"
	"testing"
)

func TestGetCurrentPubIp(t *testing.T) {
	if getCurrentPubIP() == "" {
		t.FailNow()
	}
	fmt.Println(getCurrentPubIP())
}

func TestGetCurrentPos(t *testing.T) {
	if cp := getCurrentPos(); cp == "" {
		t.FailNow()
	} else {
		fmt.Println(cp)
	}
}
