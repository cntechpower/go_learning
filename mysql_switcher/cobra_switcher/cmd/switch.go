package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	//"strings"
	//"os"
	//"flag"
)

var SwitchCMD = &cobra.Command{
	Use:   "switch",
	Short: "switch mysql M-S",
	Long:  "switch mysql M-S",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Exiting!\n")
		return
	},
}

func init() {
	RootCmd.AddCommand(SwitchCMD)
}

