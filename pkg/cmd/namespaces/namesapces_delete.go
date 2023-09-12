package namespaces

import (
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// NewCmdNamespacesDelete 构建批量删除 namespace 命令
func NewCmdNamespacesDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete naemspaces",
		Long:  "delete namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			nsBatchOp(fileName, "post", repo.NamespaceDelURL)
		},
	}
	cmd.Flags().StringVarP(&fileName, "file", "f", "", "json file for create namespace")
	cmd.MarkFlagRequired("file")
	return cmd
}
