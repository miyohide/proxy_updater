package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestShow(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "proxy_updater show", want: "show called: optint: 0, optstr: default"},
		{command: "proxy_updater show --int 10", want: "show called: optint: 10, optstr: default"},
		{command: "proxy_updater show --str test", want: "show called: optint: 0, optstr: test"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%v, get:%+v\n", c.want, get)
		}
	}
}
