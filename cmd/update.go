package cmd

import (
	"github.com/spf13/cobra"
)

// NewCmdUpdate はupdateサブコマンドを定義したもの
func NewCmdUpdate() *cobra.Command {
	type Options struct {
		optldap     string
		optpassword string
	}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("update called: optldap: %s, optpassword: %s", o.optldap, o.optpassword)
		},
	}
	cmd.Flags().StringVar(&o.optldap, "ldap", "LDAP ID", "ldap option")
	cmd.Flags().StringVar(&o.optpassword, "password", "Password", "password option")

	return cmd
}

func init() {
}
