package config

import (
	"github.com/spf13/cobra"
)

// fileName resource:config description file(format:josn) for create/delete/update
var fileName string
var configFields string

// NewCmdconfig build config root cmd
func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config [command]",
		Short: "config ",
		Long:  "config ",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for  config")
	cmd.PersistentFlags().StringVar(&configFields, "print", "", "config print field")

	// query command
	cmd.AddCommand(NewCmdConfigfile())
	cmd.AddCommand(NewCmdConfiggroup())
	cmd.AddCommand(NewCmdConfigrelease())

	// write command

	return cmd
}
