package services

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// NewCmdAlias build alias root cmd
func NewCmdAlias() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias [list|create|update|delete] service alias",
		Short: "alias [list|create|update|delete]",
		Long:  "alias [list|create|update|delete]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for create/update/delete service alias")
	cmd.PersistentFlags().StringVar(&serviceFields, "print", "", "services print field")

	// query command
	cmd.AddCommand(NewCmdAliasList())

	// write command
	cmd.AddCommand(NewCmdAliasCreate())
	cmd.AddCommand(NewCmdAliasUpdate())
	cmd.AddCommand(NewCmdAliasDelete())

	return cmd
}

// list param, eg: limit, offset
var listAliasParam entity.QueryParam
var listAliasQueryParam entity.AliasQueryParam

// NewCmdAliasList build alias list command
func NewCmdAliasList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list service alias",
		Short: "list service alias",
		Long:  "list service alias",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("alias", serviceFields)
			rsRepo := repo.NewResourceRepo(repo.RS_ALIAS, repo.API_ALIASLIST)
			rsRepo.Method("GET").Param(listAliasParam.Encode()).Print(print).Build()
		},
	}

	listAliasParam.ResourceParam = &listAliasQueryParam
	listAliasParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdAliasCreate build alias create command
func NewCmdAliasCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create service alias",
		Short: "create (-f create_serivce_alias.json)",
		Long:  "create (-f create_service_alias.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_ALIAS, repo.API_ALIAS)
			rsRepo.Method("POST").File(fileName).Batch(false).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdAliasUpdate build alias update command
func NewCmdAliasUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update alias",
		Short: "update (-f update_service_alias.json)",
		Long:  "update (-f update_service_alias.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_ALIAS, repo.API_ALIAS)
			rsRepo.Method("PUT").File(fileName).Batch(false).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdAliasDelete build alias delete command
func NewCmdAliasDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete alias",
		Short: "delete (-f delete_service_alias.json)",
		Long:  "delete (-f delete_service_alias.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_ALIAS, repo.API_ALIASDEL)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
