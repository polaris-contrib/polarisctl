package namespaces

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName resource:namespaces description file(format:josn) for create/delete/update
var fileName string

// NewCmdNamespaces build namespaces root cmd
func NewCmdNamespaces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "namespaces [list|create|delete|update]",
		Short: "namespaces [list|create|delete|update]",
		Long:  "namespaces [list|create|delete|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for create/delete/update namespaces")

	// query command
	cmd.AddCommand(NewCmdNamespacesList())

	// write command
	cmd.AddCommand(NewCmdNamespacesCreate())
	cmd.AddCommand(NewCmdNamespacesDelete())
	cmd.AddCommand(NewCmdNamespacesUpdate())

	return cmd
}

// list param, eg: limit, offset
var param entity.QueryParam
var namespacesQueryParam entity.NamespacesQueryParam

// NewCmdNamespacesList build namespaces list command
func NewCmdNamespacesList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list namespaces",
		Short: "list --limit=xx --offset=xx",
		Long:  "list --limit==xx --offset=xx",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_NAMESPACES, repo.API_NAMESPACES)
			rsRepo.Method("GET").Param(param.Encode()).Build()
		},
	}

	param.ResourceParam = &namespacesQueryParam
	param.RegisterFlag(cmd)
	return cmd
}

// NewCmdNamespacesCreate build namespaces create command
func NewCmdNamespacesCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create namespaces",
		Short: "create (-f create_namespaces.json)",
		Long:  "create (-f create_namespaces.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_NAMESPACES, repo.API_NAMESPACES)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdNamespacesDelete build namespaces delete command
func NewCmdNamespacesDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete namespaces",
		Short: "delete (-f delete_namespaces.json)",
		Long:  "delete (-f delete_namespaces.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_NAMESPACES, repo.API_NAMESPACES)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdNamespacesUpdate build namespaces update command
func NewCmdNamespacesUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update namespaces",
		Short: "update (-f update_namespaces.json)",
		Long:  "update (-f update_namespaces.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_NAMESPACES, repo.API_NAMESPACES)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
