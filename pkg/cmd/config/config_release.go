package config

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// NewCmdRelease build release root cmd
func NewCmdConfigRelease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release [list|history|versions|info|create|delete|rollback]",
		Short: "release [list|history|versions|info|create|delete|rollback]",
		Long:  "release [list|history|versions|info|create|delete|rollback]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/delete/rollback release")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "release print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdReleaseList())
	cmd.AddCommand(NewCmdReleaseHistory())
	cmd.AddCommand(NewCmdReleaseVersions())
	cmd.AddCommand(NewCmdReleaseInfo())

	// write command
	cmd.AddCommand(NewCmdReleaseRelease())
	cmd.AddCommand(NewCmdReleaseDelete())
	cmd.AddCommand(NewCmdReleaseRollback())

	return cmd
}

// list param, eg: limit, offset
var listReleaseParam entity.QueryParam
var listReleaseQueryParam entity.ConfigReleaseQueryParam

// NewCmdReleaseList build release list command
func NewCmdReleaseList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list release",
		Short: "list release",
		Long:  "list release",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RELEASE,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithParam(listReleaseParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listReleaseParam.ResourceParam = &listReleaseQueryParam
	listReleaseParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var historyReleaseParam entity.QueryParam
var historyReleaseQueryParam entity.ConfigReleaseQueryParam

// NewCmdReleaseHistory build release history command
func NewCmdReleaseHistory() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "history release",
		Short: "history release",
		Long:  "history release",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RELEASE_HIST,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigBatchQueryResponse")),

				repo.WithParam(historyReleaseParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	historyReleaseParam.ResourceParam = &historyReleaseQueryParam
	historyReleaseParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var versionsReleaseParam entity.QueryParam
var versionsReleaseQueryParam entity.ConfigReleaseQueryParam

// NewCmdReleaseVersions build release versions command
func NewCmdReleaseVersions() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "versions release",
		Short: "versions release",
		Long:  "versions release",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RELEASE_VER,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigBatchQueryResponse")),

				repo.WithParam(versionsReleaseParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	versionsReleaseParam.ResourceParam = &versionsReleaseQueryParam
	versionsReleaseParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var infoReleaseParam entity.QueryParam
var infoReleaseQueryParam entity.ConfigReleaseQueryParam

// NewCmdReleaseInfo build release info command
func NewCmdReleaseInfo() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "info release",
		Short: "info release",
		Long:  "info release",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RELEASE,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithParam(infoReleaseParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	infoReleaseParam.ResourceParam = &infoReleaseQueryParam
	infoReleaseParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdReleaseCreate build release create command
func NewCmdReleaseRelease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release release",
		Short: "release (-f create_release.json)",
		Long:  "release (-f create_release.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RELEASE,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdReleaseDelete build release delete command
func NewCmdReleaseDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete release",
		Short: "delete (-f delete_release.json)",
		Long:  "delete (-f delete_release.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RELEASE_DEL,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdReleaseRollback build release rollback command
func NewCmdReleaseRollback() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollback release",
		Short: "rollback (-f rollback_release.json)",
		Long:  "rollback (-f rollback_release.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RELEASE_ROLL,
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
