package namespaces

import "github.com/spf13/cobra"

// NewCmdNamespaces 构建 namespaces 的跟命令
func NewCmdNamespaces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "namespaces subcommand",
		Short: "namespaces ",
		Long:  "namespaces cmd",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	return cmd
}
