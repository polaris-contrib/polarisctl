package config

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// NewCmdRelease build release root cmd
func NewCmdConfigrelease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "releases [list|history|versions|info|create|delete|rollback]",
		Short: "releases [list|history|versions|info|create|delete|rollback]",
		Long:  "releases [list|history|versions|info|create|delete|rollback]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	// query command
	cmd.AddCommand(NewCmdReleaseList())
	cmd.AddCommand(NewCmdReleaseHistory())
	cmd.AddCommand(NewCmdReleaseVersions())
	cmd.AddCommand(NewCmdReleaseInfo())

	// write command
	cmd.AddCommand(NewCmdReleaseCreate())
	cmd.AddCommand(NewCmdReleaseDelete())
	cmd.AddCommand(NewCmdReleaseRollback())

	return cmd
}

// list param, eg: limit, offset
var listreleaseParam entity.QueryParam
var listreleaseQueryParam entity.ReleaseQueryParam

// NewCmdReleaseList build release list command
func NewCmdReleaseList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list releases",
		Short: "list releases",
		Long:  "list releases",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("release", configFields)
			rsRepo := repo.NewResourceRepo(repo.RS_RELEASE, repo.API_RELEASES)
			rsRepo.Method("GET").Param(listreleaseParam.Encode()).Print(print).Build()
		},
	}

	listreleaseParam.ResourceParam = &listreleaseQueryParam
	listreleaseParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var historyParam entity.QueryParam
var historyreleaseQueryParam entity.ReleaseHistoryQueryParam

// NewCmdReleaseHistory build release history command
func NewCmdReleaseHistory() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "history releases",
		Short: "history releases",
		Long:  "history releases",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("release", configFields)
			rsRepo := repo.NewResourceRepo(repo.RS_RELEASE, repo.API_RELEASE_HIST)
			rsRepo.Method("GET").Param(historyParam.Encode()).Print(print).Build()
		},
	}

	historyParam.ResourceParam = &historyreleaseQueryParam
	historyParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var versionsParam entity.QueryParam
var versionsreleaseQueryParam entity.ReleaseVersionQueryParam

// NewCmdReleaseVersions build release versions command
func NewCmdReleaseVersions() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "versions releases",
		Short: "versions releases",
		Long:  "versions releases",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("release", configFields)
			rsRepo := repo.NewResourceRepo(repo.RS_RELEASE, repo.API_RELEASE_VER)
			rsRepo.Method("GET").Param(versionsParam.Encode()).Print(print).Build()
		},
	}

	versionsParam.ResourceParam = &versionsreleaseQueryParam
	versionsParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var infoParam entity.QueryParam
var inforeleaseQueryParam entity.ReleaseInfoQueryParam

// NewCmdReleaseInfo build release info command
func NewCmdReleaseInfo() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "info releases",
		Short: "info releases",
		Long:  "info releases",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("release", configFields)
			rsRepo := repo.NewResourceRepo(repo.RS_RELEASE, repo.API_RELEASE)
			rsRepo.Method("GET").Param(infoParam.Encode()).Print(print).Build()
		},
	}

	infoParam.ResourceParam = &inforeleaseQueryParam
	infoParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdReleaseCreate build release create command
func NewCmdReleaseCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release releases",
		Short: "release (-f create_release.json)",
		Long:  "release (-f create_release.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_RELEASE, repo.API_RELEASE)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdReleaseDelete build release delete command
func NewCmdReleaseDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete releases",
		Short: "delete (-f delete_release.json)",
		Long:  "delete (-f delete_release.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_RELEASE, repo.API_RELEASE_DEL)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdReleaseRollback build release rollback command
func NewCmdReleaseRollback() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollback releases",
		Short: "rollback (-f rollback_release.json)",
		Long:  "rollback (-f rollback_release.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_RELEASE, repo.API_RELEASE_ROLL)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
