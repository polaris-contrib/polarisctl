package namespaces

import (
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// NewCmdNamespacesUpdate 构建批量创建 namespace 命令
func NewCmdNamespacesUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update naemspaces",
		Long:  "update namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			nsBatchOp(fileName, "put", repo.NamespaceURL)
		},
	}
	cmd.Flags().StringVarP(&fileName, "file", "f", "", "json file for create namespace")
	cmd.MarkFlagRequired("file")
	return cmd
}
