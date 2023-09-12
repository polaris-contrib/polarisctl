package namespaces

import (
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// NewCmdNamespacesCreate 构建批量创建 namespace 命令
func NewCmdNamespacesCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create (-f FILENAME)",
		Short: "create naemspaces",
		Long:  "create namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			nsBatchOp(fileName, "post", repo.NamespaceURL)
		},
	}
	cmd.Flags().StringVarP(&fileName, "file", "f", "", "json file for create namespace")
	cmd.MarkFlagRequired("file")
	return cmd
}
