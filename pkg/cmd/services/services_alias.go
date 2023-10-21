package services

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// NewCmdAlias build alias root cmd
func NewCmdServicesAlias() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias [list|create|update|delete]",
		Short: "alias [list|create|update|delete]",
		Long:  "alias [list|create|update|delete]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/update/delete alias")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "alias print field,eg:\"jsontag1,jsontag2\"")

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
var listAliasQueryParam entity.ServicesAliasQueryParam

// NewCmdAliasList build alias list command
func NewCmdAliasList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list alias",
		Short: "list alias",
		Long:  "list alias",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ALIAS_LIST,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(listAliasParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listAliasParam.ResourceParam = &listAliasQueryParam
	listAliasParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdAliasCreate build alias create command
func NewCmdAliasCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create alias",
		Short: "create (-f create_alias.json)",
		Long:  "create (-f create_alias.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ALIAS,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.Response")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdAliasUpdate build alias update command
func NewCmdAliasUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update alias",
		Short: "update (-f update_alias.json)",
		Long:  "update (-f update_alias.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ALIAS,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchWriteResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("PUT"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdAliasDelete build alias delete command
func NewCmdAliasDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete alias",
		Short: "delete (-f delete_alias.json)",
		Long:  "delete (-f delete_alias.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ALIAS_DEL,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchWriteResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
