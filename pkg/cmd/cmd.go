package cmd

import (
	"github.com/0226zy/polarisctl/pkg/cmd/namespaces"
	"github.com/spf13/cobra"
)

// NewDefaultPolarisCommand 构建 root 命令
func NewDefaultPolarisCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "polarisctl",
		Short: "polarisctl is used to quickly initiate related OpenAPI requests",
		Long: `polarisctl polaris command line tool is used to quickly initiate related OpenAPI requests
 Find OpenApi doc at:https://polarismesh.cn/docs/%E5%8F%82%E8%80%83%E6%96%87%E6%A1%A3/%E6%8E%A5%E5%8F%A3%E6%96%87%E6%A1%A3/open_api/`,
		Run: runHelp,
	}

	// register namespaces
	cmds.AddCommand(namespaces.NewCmdNamespaces())
	return cmds

}

// runHelp 输出帮助信息
func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
