package config

import (
	"github.com/spf13/cobra"
)

// resourceFile resource:config description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdConfig build config root cmd
func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config []",
		Short: "config []",
		Long:  "config []",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for  config")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "config print field,eg:\"jsontag1,jsontag2\"")

	// query command

	// write command
	cmd.AddCommand(NewCmdConfigGroup())
	cmd.AddCommand(NewCmdConfigFile())
	cmd.AddCommand(NewCmdConfigRelease())

	return cmd
}
