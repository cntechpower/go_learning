package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func get_flags(cmd *cobra.Command) (string, string, string, string, string, string) {
	db_pass, _ := cmd.Flags().GetString("db_pass")
	db_user, _ := cmd.Flags().GetString("db_user")
	master_db_host, _ := cmd.Flags().GetString("master_db_host")
	master_db_port, _ := cmd.Flags().GetString("master_db_port")
	slave_db_host, _ := cmd.Flags().GetString("slave_db_host")
	slave_db_port, _ := cmd.Flags().GetString("slave_db_port")
	return db_pass, db_user, master_db_host, master_db_port, slave_db_host, slave_db_port
}
func print_flags(v1 string, v2 string, v3 string, v4 string, v5 string, v6 string) {
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("db_user: ", v2)
	fmt.Println("db_pass: ", v1)
	fmt.Println("master_db_host: ", v3)
	fmt.Println("master_db_port: ", v4)
	fmt.Println("slave_db_host: ", v5)
	fmt.Println("slave_db_port: ", v6)
	fmt.Println("-------------------------------------------------------------")
}
