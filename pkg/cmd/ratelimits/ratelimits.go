package ratelimits

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName resource:ratelimits description file(format:josn) for create/delete/update
var fileName string
var ratelimitsFields string

// NewCmdRatelimits build ratelimits root cmd
func NewCmdRatelimits() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ratelimits [list|create|update|enable|delete]",
		Short: "ratelimits [list|create|update|enable|delete]",
		Long:  "ratelimits [list|create|update|enable|delete]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for create/update/enable/delete ratelimits")
	cmd.PersistentFlags().StringVar(&ratelimitsFields, "print", "", "ratelimits print field")

	// query command
	cmd.AddCommand(NewCmdRatelimitsList())

	// write command
	cmd.AddCommand(NewCmdRatelimitsCreate())
	cmd.AddCommand(NewCmdRatelimitsUpdate())
	cmd.AddCommand(NewCmdRatelimitsEnable())
	cmd.AddCommand(NewCmdRatelimitsDelete())

	return cmd
}

// list param, eg: limit, offset
var listParam entity.QueryParam
var listRatelimitsQueryParam entity.RatelimitsQueryParam

// NewCmdRatelimitsList build ratelimits list command
func NewCmdRatelimitsList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list ratelimits",
		Short: "list ratelimits",
		Long:  "list ratelimits",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("rateLimits", ratelimitsFields)
			rsRepo := repo.NewResourceRepo(repo.RS_RATELIMITS, repo.API_RATELIMITS)
			rsRepo.Method("GET").Param(listParam.Encode()).Print(print).Build()
		},
	}

	listParam.ResourceParam = &listRatelimitsQueryParam
	listParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdRatelimitsCreate build ratelimits create command
func NewCmdRatelimitsCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create ratelimits",
		Short: "create (-f create_ratelimits.json)",
		Long:  "create (-f create_ratelimits.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_RATELIMITS, repo.API_RATELIMITS)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdRatelimitsUpdate build ratelimits update command
func NewCmdRatelimitsUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update ratelimits",
		Short: "update (-f update_ratelimits.json)",
		Long:  "update (-f update_ratelimits.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_RATELIMITS, repo.API_RATELIMITS)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdRatelimitsEnable build ratelimits enable command
func NewCmdRatelimitsEnable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable ratelimits",
		Short: "enable (-f enable_ratelimits.json)",
		Long:  "enable (-f enable_ratelimits.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_RATELIMITS, repo.API_RATELIMITS_ENABLE)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdRatelimitsDelete build ratelimits delete command
func NewCmdRatelimitsDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete ratelimits",
		Short: "delete (-f delete_ratelimits.json)",
		Long:  "delete (-f delete_ratelimits.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_RATELIMITS, repo.API_RATELIMITS_DEL)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
