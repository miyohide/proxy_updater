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
		{command: "proxy_updater show", want: "show called: optldap: LDAP ID, optpassword: Password"},
		{command: "proxy_updater show --ldap 10", want: "show called: optldap: 10, optpassword: Password"},
		{command: "proxy_updater show --password test", want: "show called: optldap: LDAP ID, optpassword: test"},
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
