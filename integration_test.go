// +build integration

package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestBadArgs(t *testing.T) {
	var err error
	cmd := exec.Command("forecast", "-wrong")
	out, err := cmd.CombinedOutput()
	sout := string(out) // because out is []byte
	if err != nil && !strings.Contains(sout, "Usage: forecast [flags]") {
		fmt.Println(sout) // so we can see the full Output
		t.Errorf("%v", err)
	}
}


