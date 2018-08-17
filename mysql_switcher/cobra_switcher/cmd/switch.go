package cmd

import (
	"github.com/spf13/cobra"
)

var SwitchCMD = &cobra.Command{
	Use:   "switch",
	Short: "switch mysql M-S",
	Long:  "switch mysql M-S",
	Run: func(cmd *cobra.Command, args []string) {
		db_pass, db_user, master_db_host, master_db_port, slave_db_host, slave_db_port = get_flags(cmd)
		print_flags(db_pass, db_user, master_db_host, master_db_port, slave_db_host, slave_db_port)
		return
	},
}

func init() {
	RootCmd.AddCommand(SwitchCMD)
}
